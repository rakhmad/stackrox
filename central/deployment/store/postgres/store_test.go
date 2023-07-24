// Code generated by pg-bindings generator. DO NOT EDIT.

//go:build sql_integration

package postgres

import (
	"context"
	"fmt"
	"testing"

	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/postgres/pgtest"
	"github.com/stackrox/rox/pkg/sac"
	"github.com/stackrox/rox/pkg/testutils"
	"github.com/stackrox/rox/pkg/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type DeploymentsStoreSuite struct {
	suite.Suite
	store  Store
	testDB *pgtest.TestPostgres
}

func TestDeploymentsStore(t *testing.T) {
	suite.Run(t, new(DeploymentsStoreSuite))
}

func (s *DeploymentsStoreSuite) SetupSuite() {

	s.testDB = pgtest.ForT(s.T())
	s.store = New(s.testDB.DB)
}

func (s *DeploymentsStoreSuite) SetupTest() {
	ctx := sac.WithAllAccess(context.Background())
	tag, err := s.testDB.Exec(ctx, "TRUNCATE deployments CASCADE")
	s.T().Log("deployments", tag)
	s.NoError(err)
}

func (s *DeploymentsStoreSuite) TearDownSuite() {
	s.testDB.Teardown(s.T())
}

func (s *DeploymentsStoreSuite) TestStore() {
	ctx := sac.WithAllAccess(context.Background())

	store := s.store

	deployment := &storage.Deployment{}
	s.NoError(testutils.FullInit(deployment, testutils.SimpleInitializer(), testutils.JSONFieldsFilter))

	foundDeployment, exists, err := store.Get(ctx, deployment.GetId())
	s.NoError(err)
	s.False(exists)
	s.Nil(foundDeployment)

	withNoAccessCtx := sac.WithNoAccess(ctx)

	s.NoError(store.Upsert(ctx, deployment))
	foundDeployment, exists, err = store.Get(ctx, deployment.GetId())
	s.NoError(err)
	s.True(exists)
	s.Equal(deployment, foundDeployment)

	deploymentCount, err := store.Count(ctx)
	s.NoError(err)
	s.Equal(1, deploymentCount)
	deploymentCount, err = store.Count(withNoAccessCtx)
	s.NoError(err)
	s.Zero(deploymentCount)

	deploymentExists, err := store.Exists(ctx, deployment.GetId())
	s.NoError(err)
	s.True(deploymentExists)
	s.NoError(store.Upsert(ctx, deployment))
	s.ErrorIs(store.Upsert(withNoAccessCtx, deployment), sac.ErrResourceAccessDenied)

	foundDeployment, exists, err = store.Get(ctx, deployment.GetId())
	s.NoError(err)
	s.True(exists)
	s.Equal(deployment, foundDeployment)

	s.NoError(store.Delete(ctx, deployment.GetId()))
	foundDeployment, exists, err = store.Get(ctx, deployment.GetId())
	s.NoError(err)
	s.False(exists)
	s.Nil(foundDeployment)
	s.NoError(store.Delete(withNoAccessCtx, deployment.GetId()))

	var deployments []*storage.Deployment
	var deploymentIDs []string
	for i := 0; i < 200; i++ {
		deployment := &storage.Deployment{}
		s.NoError(testutils.FullInit(deployment, testutils.UniqueInitializer(), testutils.JSONFieldsFilter))
		deployments = append(deployments, deployment)
		deploymentIDs = append(deploymentIDs, deployment.GetId())
	}

	s.NoError(store.UpsertMany(ctx, deployments))

	deploymentCount, err = store.Count(ctx)
	s.NoError(err)
	s.Equal(200, deploymentCount)

	s.NoError(store.DeleteMany(ctx, deploymentIDs))

	deploymentCount, err = store.Count(ctx)
	s.NoError(err)
	s.Equal(0, deploymentCount)
}

const (
	withAllAccess           = "AllAccess"
	withNoAccess            = "NoAccess"
	withAccessToDifferentNs = "AccessToDifferentNs"
	withAccess              = "Access"
	withAccessToCluster     = "AccessToCluster"
	withNoAccessToCluster   = "NoAccessToCluster"
)

var (
	withAllAccessCtx = sac.WithAllAccess(context.Background())
)

type testCase struct {
	context                context.Context
	expectedIDs            []string
	expectedIdentifiers    []string
	expectedMissingIndices []int
	expectedObjects        []*storage.Deployment
	expectedWriteError     error
}

func (s *DeploymentsStoreSuite) getTestData(access storage.Access) (*storage.Deployment, *storage.Deployment, map[string]testCase) {
	objA := &storage.Deployment{}
	s.NoError(testutils.FullInit(objA, testutils.SimpleInitializer(), testutils.JSONFieldsFilter))

	objB := &storage.Deployment{}
	s.NoError(testutils.FullInit(objB, testutils.SimpleInitializer(), testutils.JSONFieldsFilter))

	testCases := map[string]testCase{
		withAllAccess: {
			context:                sac.WithAllAccess(context.Background()),
			expectedMissingIndices: nil,
			expectedWriteError:     nil,
		},
		withNoAccess: {
			context:                sac.WithNoAccess(context.Background()),
			expectedMissingIndices: nil,
			expectedWriteError:     sac.ErrResourceAccessDenied,
		},
		withNoAccessToCluster: {
			context: sac.WithGlobalAccessScopeChecker(context.Background(),
				sac.AllowFixedScopes(
					sac.AccessModeScopeKeys(access),
					sac.ResourceScopeKeys(targetResource),
					sac.ClusterScopeKeys(uuid.Nil.String()),
				),
			),
			expectedMissingIndices: nil,
			expectedWriteError:     sac.ErrResourceAccessDenied,
		},
		withAccessToDifferentNs: {
			context: sac.WithGlobalAccessScopeChecker(context.Background(),
				sac.AllowFixedScopes(
					sac.AccessModeScopeKeys(access),
					sac.ResourceScopeKeys(targetResource),
					sac.ClusterScopeKeys(objA.GetClusterId()),
					sac.NamespaceScopeKeys("unknown ns"),
				),
			),
			expectedMissingIndices: nil,
			expectedWriteError:     sac.ErrResourceAccessDenied,
		},
		withAccess: {
			context: sac.WithGlobalAccessScopeChecker(context.Background(),
				sac.AllowFixedScopes(
					sac.AccessModeScopeKeys(access),
					sac.ResourceScopeKeys(targetResource),
					sac.ClusterScopeKeys(objA.GetClusterId()),
					sac.NamespaceScopeKeys(objA.GetNamespace()),
				),
			),
			expectedMissingIndices: nil,
			expectedWriteError:     nil,
		},
		withAccessToCluster: {
			context: sac.WithGlobalAccessScopeChecker(context.Background(),
				sac.AllowFixedScopes(
					sac.AccessModeScopeKeys(access),
					sac.ResourceScopeKeys(targetResource),
					sac.ClusterScopeKeys(objA.GetClusterId()),
				),
			),
			expectedMissingIndices: nil,
			expectedWriteError:     nil,
		},
	}

	return objA, objB, testCases
}

func (s *DeploymentsStoreSuite) TestSACUpsert() {
	obj, _, testCases := s.getTestData(storage.Access_READ_WRITE_ACCESS)
	for name, testCase := range testCases {
		s.T().Run(fmt.Sprintf("with %s", name), func(t *testing.T) {
			assert.ErrorIs(t, s.store.Upsert(testCase.context, obj), testCase.expectedWriteError)
		})
	}
}

func (s *DeploymentsStoreSuite) TestSACUpsertMany() {
	obj, _, testCases := s.getTestData(storage.Access_READ_WRITE_ACCESS)
	for name, testCase := range testCases {
		s.T().Run(fmt.Sprintf("with %s", name), func(t *testing.T) {
			assert.ErrorIs(t, s.store.UpsertMany(testCase.context, []*storage.Deployment{obj}), testCase.expectedWriteError)
		})
	}
}

func (s *DeploymentsStoreSuite) TestSACCount() {
	objA := &storage.Deployment{}
	s.NoError(testutils.FullInit(objA, testutils.UniqueInitializer(), testutils.JSONFieldsFilter))

	objB := &storage.Deployment{}
	s.NoError(testutils.FullInit(objB, testutils.UniqueInitializer(), testutils.JSONFieldsFilter))

	withAllAccessCtx := sac.WithAllAccess(context.Background())
	s.Require().NoError(s.store.Upsert(withAllAccessCtx, objA))
	s.Require().NoError(s.store.Upsert(withAllAccessCtx, objB))

	ctxs := getSACContexts(objA, storage.Access_READ_ACCESS)
	for name, expectedCount := range map[string]int{
		withAllAccess:           2,
		withNoAccess:            0,
		withNoAccessToCluster:   0,
		withAccessToDifferentNs: 0,
		withAccess:              1,
		withAccessToCluster:     1,
	} {
		s.T().Run(fmt.Sprintf("with %s", name), func(t *testing.T) {
			count, err := s.store.Count(ctxs[name])
			assert.NoError(t, err)
			assert.Equal(t, expectedCount, count)
		})
	}
}

func (s *DeploymentsStoreSuite) TestSACWalk() {
	objA := &storage.Deployment{}
	s.NoError(testutils.FullInit(objA, testutils.UniqueInitializer(), testutils.JSONFieldsFilter))

	objB := &storage.Deployment{}
	s.NoError(testutils.FullInit(objB, testutils.UniqueInitializer(), testutils.JSONFieldsFilter))

	withAllAccessCtx := sac.WithAllAccess(context.Background())
	s.Require().NoError(s.store.Upsert(withAllAccessCtx, objA))
	s.Require().NoError(s.store.Upsert(withAllAccessCtx, objB))

	ctxs := getSACContexts(objA, storage.Access_READ_ACCESS)
	for name, expectedIDs := range map[string][]string{
		withAllAccess:           []string{objA.GetId(), objB.GetId()},
		withNoAccess:            []string{},
		withNoAccessToCluster:   []string{},
		withAccessToDifferentNs: []string{},
		withAccess:              []string{objA.GetId()},
		withAccessToCluster:     []string{objA.GetId()},
	} {
		s.T().Run(fmt.Sprintf("with %s", name), func(t *testing.T) {
			identifiers := []string{}
			getIDs := func(obj *storage.Deployment) error {
				identifiers = append(identifiers, obj.GetId())
				return nil
			}
			err := s.store.Walk(ctxs[name], getIDs)
			assert.NoError(t, err)
			assert.ElementsMatch(t, expectedIDs, identifiers)
		})
	}
}

func (s *DeploymentsStoreSuite) TestSACGetIDs() {
	objA := &storage.Deployment{}
	s.NoError(testutils.FullInit(objA, testutils.UniqueInitializer(), testutils.JSONFieldsFilter))

	objB := &storage.Deployment{}
	s.NoError(testutils.FullInit(objB, testutils.UniqueInitializer(), testutils.JSONFieldsFilter))

	withAllAccessCtx := sac.WithAllAccess(context.Background())
	s.Require().NoError(s.store.Upsert(withAllAccessCtx, objA))
	s.Require().NoError(s.store.Upsert(withAllAccessCtx, objB))

	ctxs := getSACContexts(objA, storage.Access_READ_ACCESS)
	for name, expectedIDs := range map[string][]string{
		withAllAccess:           []string{objA.GetId(), objB.GetId()},
		withNoAccess:            []string{},
		withNoAccessToCluster:   []string{},
		withAccessToDifferentNs: []string{},
		withAccess:              []string{objA.GetId()},
		withAccessToCluster:     []string{objA.GetId()},
	} {
		s.T().Run(fmt.Sprintf("with %s", name), func(t *testing.T) {
			identifiers, err := s.store.GetIDs(ctxs[name])
			assert.NoError(t, err)
			assert.EqualValues(t, expectedIDs, identifiers)
		})
	}
}

func (s *DeploymentsStoreSuite) TestSACExists() {
	objA := &storage.Deployment{}
	s.NoError(testutils.FullInit(objA, testutils.UniqueInitializer(), testutils.JSONFieldsFilter))

	withAllAccessCtx := sac.WithAllAccess(context.Background())
	s.Require().NoError(s.store.Upsert(withAllAccessCtx, objA))

	ctxs := getSACContexts(objA, storage.Access_READ_ACCESS)
	for name, expected := range map[string]bool{
		withAllAccess:           true,
		withNoAccess:            false,
		withNoAccessToCluster:   false,
		withAccessToDifferentNs: false,
		withAccess:              true,
		withAccessToCluster:     true,
	} {
		s.T().Run(fmt.Sprintf("with %s", name), func(t *testing.T) {
			exists, err := s.store.Exists(ctxs[name], objA.GetId())
			assert.NoError(t, err)
			assert.Equal(t, expected, exists)
		})
	}
}

func (s *DeploymentsStoreSuite) TestSACGet() {
	objA := &storage.Deployment{}
	s.NoError(testutils.FullInit(objA, testutils.UniqueInitializer(), testutils.JSONFieldsFilter))

	withAllAccessCtx := sac.WithAllAccess(context.Background())
	s.Require().NoError(s.store.Upsert(withAllAccessCtx, objA))

	ctxs := getSACContexts(objA, storage.Access_READ_ACCESS)
	for name, expected := range map[string]bool{
		withAllAccess:           true,
		withNoAccess:            false,
		withNoAccessToCluster:   false,
		withAccessToDifferentNs: false,
		withAccess:              true,
		withAccessToCluster:     true,
	} {
		s.T().Run(fmt.Sprintf("with %s", name), func(t *testing.T) {
			actual, exists, err := s.store.Get(ctxs[name], objA.GetId())
			assert.NoError(t, err)
			assert.Equal(t, expected, exists)
			if expected == true {
				assert.Equal(t, objA, actual)
			} else {
				assert.Nil(t, actual)
			}
		})
	}
}

func (s *DeploymentsStoreSuite) TestSACDelete() {
	objA := &storage.Deployment{}
	s.NoError(testutils.FullInit(objA, testutils.UniqueInitializer(), testutils.JSONFieldsFilter))

	objB := &storage.Deployment{}
	s.NoError(testutils.FullInit(objB, testutils.UniqueInitializer(), testutils.JSONFieldsFilter))
	withAllAccessCtx := sac.WithAllAccess(context.Background())

	ctxs := getSACContexts(objA, storage.Access_READ_WRITE_ACCESS)
	for name, expectedCount := range map[string]int{
		withAllAccess:           0,
		withNoAccess:            2,
		withNoAccessToCluster:   2,
		withAccessToDifferentNs: 2,
		withAccess:              1,
		withAccessToCluster:     1,
	} {
		s.T().Run(fmt.Sprintf("with %s", name), func(t *testing.T) {
			s.SetupTest()

			s.NoError(s.store.Upsert(withAllAccessCtx, objA))
			s.NoError(s.store.Upsert(withAllAccessCtx, objB))

			assert.NoError(t, s.store.Delete(ctxs[name], objA.GetId()))
			assert.NoError(t, s.store.Delete(ctxs[name], objB.GetId()))

			count, err := s.store.Count(withAllAccessCtx)
			assert.NoError(t, err)
			assert.Equal(t, expectedCount, count)
		})
	}
}

func (s *DeploymentsStoreSuite) TestSACDeleteMany() {
	objA := &storage.Deployment{}
	s.NoError(testutils.FullInit(objA, testutils.UniqueInitializer(), testutils.JSONFieldsFilter))

	objB := &storage.Deployment{}
	s.NoError(testutils.FullInit(objB, testutils.UniqueInitializer(), testutils.JSONFieldsFilter))
	withAllAccessCtx := sac.WithAllAccess(context.Background())

	ctxs := getSACContexts(objA, storage.Access_READ_WRITE_ACCESS)
	for name, expectedCount := range map[string]int{
		withAllAccess:           0,
		withNoAccess:            2,
		withNoAccessToCluster:   2,
		withAccessToDifferentNs: 2,
		withAccess:              1,
		withAccessToCluster:     1,
	} {
		s.T().Run(fmt.Sprintf("with %s", name), func(t *testing.T) {
			s.SetupTest()

			s.NoError(s.store.Upsert(withAllAccessCtx, objA))
			s.NoError(s.store.Upsert(withAllAccessCtx, objB))

			assert.NoError(t, s.store.DeleteMany(ctxs[name], []string{
				objA.GetId(),
				objB.GetId(),
			}))

			count, err := s.store.Count(withAllAccessCtx)
			assert.NoError(t, err)
			assert.Equal(t, expectedCount, count)
		})
	}
}

func (s *DeploymentsStoreSuite) TestSACGetMany() {
	objA := &storage.Deployment{}
	s.NoError(testutils.FullInit(objA, testutils.UniqueInitializer(), testutils.JSONFieldsFilter))

	objB := &storage.Deployment{}
	s.NoError(testutils.FullInit(objB, testutils.UniqueInitializer(), testutils.JSONFieldsFilter))

	withAllAccessCtx := sac.WithAllAccess(context.Background())
	s.Require().NoError(s.store.Upsert(withAllAccessCtx, objA))
	s.Require().NoError(s.store.Upsert(withAllAccessCtx, objB))

	ctxs := getSACContexts(objA, storage.Access_READ_ACCESS)
	for name, expected := range map[string]struct {
		elems          []*storage.Deployment
		missingIndices []int
	}{
		withAllAccess:           {elems: []*storage.Deployment{objA, objB}, missingIndices: []int{}},
		withNoAccess:            {elems: []*storage.Deployment{}, missingIndices: []int{0, 1}},
		withNoAccessToCluster:   {elems: []*storage.Deployment{}, missingIndices: []int{0, 1}},
		withAccessToDifferentNs: {elems: []*storage.Deployment{}, missingIndices: []int{0, 1}},
		withAccess:              {elems: []*storage.Deployment{objA}, missingIndices: []int{1}},
		withAccessToCluster:     {elems: []*storage.Deployment{objA}, missingIndices: []int{1}},
	} {
		s.T().Run(fmt.Sprintf("with %s", name), func(t *testing.T) {
			actual, missingIndices, err := s.store.GetMany(ctxs[name], []string{objA.GetId(), objB.GetId()})
			assert.NoError(t, err)
			assert.Equal(t, expected.elems, actual)
			assert.Equal(t, expected.missingIndices, missingIndices)
		})
	}

	s.T().Run("with no identifiers", func(t *testing.T) {
		actual, missingIndices, err := s.store.GetMany(withAllAccessCtx, []string{})
		assert.Nil(t, err)
		assert.Nil(t, actual)
		assert.Nil(t, missingIndices)
	})
}

func getSACContexts(obj *storage.Deployment, access storage.Access) map[string]context.Context {
	return map[string]context.Context{
		withAllAccess: sac.WithAllAccess(context.Background()),
		withNoAccess:  sac.WithNoAccess(context.Background()),
		withAccessToDifferentNs: sac.WithGlobalAccessScopeChecker(context.Background(),
			sac.AllowFixedScopes(
				sac.AccessModeScopeKeys(access),
				sac.ResourceScopeKeys(targetResource),
				sac.ClusterScopeKeys(obj.GetClusterId()),
				sac.NamespaceScopeKeys("unknown ns"),
			)),
		withAccess: sac.WithGlobalAccessScopeChecker(context.Background(),
			sac.AllowFixedScopes(
				sac.AccessModeScopeKeys(access),
				sac.ResourceScopeKeys(targetResource),
				sac.ClusterScopeKeys(obj.GetClusterId()),
				sac.NamespaceScopeKeys(obj.GetNamespace()),
			)),
		withAccessToCluster: sac.WithGlobalAccessScopeChecker(context.Background(),
			sac.AllowFixedScopes(
				sac.AccessModeScopeKeys(access),
				sac.ResourceScopeKeys(targetResource),
				sac.ClusterScopeKeys(obj.GetClusterId()),
			)),
		withNoAccessToCluster: sac.WithGlobalAccessScopeChecker(context.Background(),
			sac.AllowFixedScopes(
				sac.AccessModeScopeKeys(access),
				sac.ResourceScopeKeys(targetResource),
				sac.ClusterScopeKeys(uuid.Nil.String()),
			)),
	}
}
