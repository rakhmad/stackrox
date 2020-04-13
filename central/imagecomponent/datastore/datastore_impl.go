package datastore

import (
	"context"

	"github.com/pkg/errors"
	"github.com/stackrox/rox/central/imagecomponent/index"
	"github.com/stackrox/rox/central/imagecomponent/search"
	"github.com/stackrox/rox/central/imagecomponent/store"
	"github.com/stackrox/rox/central/ranking"
	riskDataStore "github.com/stackrox/rox/central/risk/datastore"
	"github.com/stackrox/rox/central/role/resources"
	v1 "github.com/stackrox/rox/generated/api/v1"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/logging"
	"github.com/stackrox/rox/pkg/sac"
	pkgSearch "github.com/stackrox/rox/pkg/search"
)

var (
	log = logging.LoggerForModule()
	// TODO: Need to setup sac for Image Components correctly instead of relying on global access.
	imagesSAC = sac.ForResource(resources.Image)
)

type datastoreImpl struct {
	storage  store.Store
	indexer  index.Indexer
	searcher search.Searcher

	risks                riskDataStore.DataStore
	imageComponentRanker *ranking.Ranker
}

func (ds *datastoreImpl) Search(ctx context.Context, q *v1.Query) ([]pkgSearch.Result, error) {
	return ds.searcher.Search(ctx, q)
}

func (ds *datastoreImpl) SearchImageComponents(ctx context.Context, q *v1.Query) ([]*v1.SearchResult, error) {
	return ds.searcher.SearchImageComponents(ctx, q)
}

func (ds *datastoreImpl) SearchRawImageComponents(ctx context.Context, q *v1.Query) ([]*storage.ImageComponent, error) {
	components, err := ds.searcher.SearchRawImageComponents(ctx, q)
	if err != nil {
		return nil, err
	}
	ds.updateImageComponentPriority(components...)
	return components, nil
}

func (ds *datastoreImpl) Count(ctx context.Context) (int, error) {
	results, err := ds.searcher.Search(ctx, pkgSearch.EmptyQuery())
	if err != nil {
		return 0, err
	}
	return len(results), nil
}

func (ds *datastoreImpl) Get(ctx context.Context, id string) (*storage.ImageComponent, bool, error) {
	if ok, err := imagesSAC.ReadAllowed(ctx); err != nil || !ok {
		return nil, false, err
	}
	component, found, err := ds.storage.Get(id)
	if err != nil || !found {
		return nil, false, err
	}

	ds.updateImageComponentPriority(component)
	return component, true, nil
}

func (ds *datastoreImpl) Exists(ctx context.Context, id string) (bool, error) {
	if ok, err := imagesSAC.ReadAllowed(ctx); err != nil || !ok {
		return false, err
	}
	found, err := ds.storage.Exists(id)
	if err != nil || !found {
		return false, err
	}
	return true, nil
}

func (ds *datastoreImpl) GetBatch(ctx context.Context, ids []string) ([]*storage.ImageComponent, error) {
	if ok, err := imagesSAC.ReadAllowed(ctx); err != nil || !ok {
		return nil, err
	}
	components, _, err := ds.storage.GetBatch(ids)
	if err != nil {
		return nil, err
	}

	ds.updateImageComponentPriority(components...)
	return components, nil
}

// UpsertImage dedupes the image with the underlying storage and adds the image to the index.
func (ds *datastoreImpl) Upsert(ctx context.Context, imagecomponents ...*storage.ImageComponent) error {
	if len(imagecomponents) == 0 {
		return nil
	}
	if ok, err := imagesSAC.WriteAllowed(ctx); err != nil {
		return err
	} else if !ok {
		return errors.New("permission denied")
	}

	// Update image components with latest risk score
	for _, component := range imagecomponents {
		component.RiskScore = ds.imageComponentRanker.GetScoreForID(component.GetId())
	}

	return ds.storage.Upsert(imagecomponents...)
}

func (ds *datastoreImpl) Delete(ctx context.Context, ids ...string) error {
	if ok, err := imagesSAC.WriteAllowed(ctx); err != nil {
		return err
	} else if !ok {
		return errors.New("permission denied")
	}

	if err := ds.storage.Delete(ids...); err != nil {
		return err
	}

	deleteRiskCtx := sac.WithGlobalAccessScopeChecker(ctx,
		sac.AllowFixedScopes(
			sac.AccessModeScopeKeys(storage.Access_READ_WRITE_ACCESS),
			sac.ResourceScopeKeys(resources.Risk),
		))

	for _, id := range ids {
		if err := ds.risks.RemoveRisk(deleteRiskCtx, id, storage.RiskSubjectType_IMAGE_COMPONENT); err != nil {
			return err
		}
	}
	return nil
}

func (ds *datastoreImpl) initializeRankers() {
	readCtx := sac.WithGlobalAccessScopeChecker(context.Background(),
		sac.AllowFixedScopes(
			sac.AccessModeScopeKeys(storage.Access_READ_ACCESS), sac.ResourceScopeKeys(resources.Image)))

	results, err := ds.Search(readCtx, pkgSearch.EmptyQuery())
	if err != nil {
		log.Error(err)
		return
	}

	for _, id := range pkgSearch.ResultsToIDs(results) {
		component, found, err := ds.storage.Get(id)
		if err != nil {
			log.Error(err)
			continue
		} else if !found {
			continue
		}
		ds.imageComponentRanker.Add(id, component.GetRiskScore())
	}
}

func (ds *datastoreImpl) updateImageComponentPriority(ics ...*storage.ImageComponent) {
	for _, ic := range ics {
		ic.Priority = ds.imageComponentRanker.GetRankForID(ic.GetId())
	}
}
