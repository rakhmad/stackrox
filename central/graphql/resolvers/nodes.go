package resolvers

import (
	"context"
	"strings"
	"time"

	"github.com/graph-gophers/graphql-go"
	"github.com/pkg/errors"
	complianceStandards "github.com/stackrox/rox/central/compliance/standards"
	"github.com/stackrox/rox/central/metrics"
	v1 "github.com/stackrox/rox/generated/api/v1"
	"github.com/stackrox/rox/generated/storage"
	pkgMetrics "github.com/stackrox/rox/pkg/metrics"
	"github.com/stackrox/rox/pkg/search"
	"github.com/stackrox/rox/pkg/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func init() {
	schema := getBuilder()
	utils.Must(
		schema.AddQuery("node(id:ID!): Node"),
		schema.AddQuery("nodes(query: String, pagination: Pagination): [Node!]!"),
		schema.AddQuery("nodeCount(query: String): Int!"),
		schema.AddExtraResolver("Node", "complianceResults(query: String): [ControlResult!]!"),
		schema.AddType("ComplianceControlCount", []string{"failingCount: Int!", "passingCount: Int!", "unknownCount: Int!"}),
		schema.AddExtraResolver("Node", "nodeComplianceControlCount(query: String) : ComplianceControlCount!"),
		schema.AddExtraResolver("Node", "controlStatus(query: String): String!"),
		schema.AddExtraResolver("Node", "failingControls(query: String): [ComplianceControl!]!"),
		schema.AddExtraResolver("Node", "passingControls(query: String): [ComplianceControl!]!"),
		schema.AddExtraResolver("Node", "controls(query: String): [ComplianceControl!]!"),
		schema.AddExtraResolver("Node", "cluster: Cluster!"),
	)
}

// Node returns a resolver for a matching node, or nil if no node is found in any cluster
func (resolver *Resolver) Node(ctx context.Context, args struct{ graphql.ID }) (*nodeResolver, error) {
	defer metrics.SetGraphQLOperationDurationTime(time.Now(), pkgMetrics.Root, "Node")
	if err := readNodes(ctx); err != nil {
		return nil, err
	}
	clusters, err := resolver.ClusterDataStore.GetClusters(ctx)
	if err != nil {
		return nil, err
	}
	var output *nodeResolver
	for _, cluster := range clusters {
		store, err := resolver.NodeGlobalDataStore.GetClusterNodeStore(ctx, cluster.GetId(), false)
		if err != nil {
			return nil, err
		}
		node, err := store.GetNode(string(args.ID))
		if err != nil {
			return nil, err
		}
		if node != nil {
			if output == nil {
				output = &nodeResolver{root: resolver, data: node}
			} else {
				return nil, status.Error(codes.Internal, "multiple matching node ids found")
			}
		}
	}
	return output, nil
}

// Nodes returns resolvers for a matching nodes, or nil if no node is found in any cluster
func (resolver *Resolver) Nodes(ctx context.Context, args PaginatedQuery) ([]*nodeResolver, error) {
	defer metrics.SetGraphQLOperationDurationTime(time.Now(), pkgMetrics.Root, "Nodes")
	if err := readNodes(ctx); err != nil {
		return nil, err
	}
	query, err := args.AsV1QueryOrEmpty()
	if err != nil {
		return nil, err
	}

	var nodeResolvers []*nodeResolver
	nodes, err := resolver.NodeGlobalDataStore.SearchRawNodes(ctx, query)
	if err != nil {
		return nil, err
	}

	for _, node := range nodes {
		nodeResolvers = append(nodeResolvers, &nodeResolver{root: resolver, data: node})
	}

	resolvers, err := paginationWrapper{
		pv: query.Pagination,
	}.paginate(nodeResolvers, nil)

	return resolvers.([]*nodeResolver), err
}

// NodeCount returns count of nodes across clusters
func (resolver *Resolver) NodeCount(ctx context.Context, args RawQuery) (int32, error) {
	defer metrics.SetGraphQLOperationDurationTime(time.Now(), pkgMetrics.Root, "NodeCount")
	if err := readNodes(ctx); err != nil {
		return 0, err
	}
	query, err := args.AsV1QueryOrEmpty()
	if err != nil {
		return 0, err
	}
	results, err := resolver.NodeGlobalDataStore.Search(ctx, query)
	if err != nil {
		return 0, err
	}
	return int32(len(results)), nil
}

func (resolver *nodeResolver) Cluster(ctx context.Context) (*clusterResolver, error) {
	defer metrics.SetGraphQLOperationDurationTime(time.Now(), pkgMetrics.Nodes, "Cluster")
	if err := readClusters(ctx); err != nil {
		return nil, err
	}
	return resolver.root.wrapCluster(resolver.root.ClusterDataStore.GetCluster(ctx, resolver.data.GetClusterId()))
}

func (resolver *nodeResolver) ComplianceResults(ctx context.Context, args RawQuery) ([]*controlResultResolver, error) {
	defer metrics.SetGraphQLOperationDurationTime(time.Now(), pkgMetrics.Nodes, "ComplianceResults")
	if err := readCompliance(ctx); err != nil {
		return nil, err
	}

	runResults, err := resolver.root.ComplianceAggregator.GetResultsWithEvidence(ctx, args.String())
	if err != nil {
		return nil, err
	}
	output := newBulkControlResults()
	nodeID := resolver.data.GetId()
	output.addNodeData(resolver.root, runResults, func(node *storage.Node, _ *v1.ComplianceControl) bool {
		return node.GetId() == nodeID
	})
	return *output, nil
}

func (resolver *nodeResolver) NodeComplianceControlCount(ctx context.Context, args RawQuery) (*complianceControlCountResolver, error) {
	defer metrics.SetGraphQLOperationDurationTime(time.Now(), pkgMetrics.Nodes, "NodeComplianceControlCount")
	if err := readCompliance(ctx); err != nil {
		return nil, err
	}
	scope := []storage.ComplianceAggregation_Scope{storage.ComplianceAggregation_NODE, storage.ComplianceAggregation_CONTROL}
	results, err := resolver.getNodeLastSuccessfulComplianceRunAggregatedResult(ctx, scope, args)
	if err != nil {
		return nil, err
	}
	if results == nil {
		return &complianceControlCountResolver{}, nil
	}
	return getComplianceControlCountFromAggregationResults(results), nil
}

func (resolver *nodeResolver) ControlStatus(ctx context.Context, args RawQuery) (string, error) {
	defer metrics.SetGraphQLOperationDurationTime(time.Now(), pkgMetrics.Nodes, "ControlStatus")
	if err := readCompliance(ctx); err != nil {
		return "Fail", err
	}
	r, err := resolver.getNodeLastSuccessfulComplianceRunAggregatedResult(ctx, []storage.ComplianceAggregation_Scope{storage.ComplianceAggregation_NODE}, args)
	if err != nil || r == nil {
		return "Fail", err
	}
	if len(r) != 1 {
		return "Fail", errors.Errorf("unexpected node aggregation results length: expected: 1, actual: %d", len(r))
	}
	return getControlStatusFromAggregationResult(r[0]), nil
}

func (resolver *nodeResolver) getNodeLastSuccessfulComplianceRunAggregatedResult(ctx context.Context, scope []storage.ComplianceAggregation_Scope, args RawQuery) ([]*storage.ComplianceAggregation_Result, error) {
	if err := readCompliance(ctx); err != nil {
		return nil, err
	}
	standardIDs, err := getStandardIDs(ctx, resolver.root.ComplianceStandardStore)
	if err != nil {
		return nil, err
	}
	hasComplianceSuccessfullyRun, err := resolver.root.ComplianceDataStore.IsComplianceRunSuccessfulOnCluster(ctx, resolver.data.GetClusterId(), standardIDs)
	if err != nil || !hasComplianceSuccessfullyRun {
		return nil, err
	}
	query, err := search.NewQueryBuilder().AddExactMatches(search.ClusterID, resolver.data.GetClusterId()).
		AddExactMatches(search.NodeID, resolver.data.GetId()).RawQuery()
	if err != nil {
		return nil, err
	}
	if args.Query != nil {
		query = strings.Join([]string{query, args.String()}, "+")
	}
	r, _, _, err := resolver.root.ComplianceAggregator.Aggregate(ctx, query, scope, storage.ComplianceAggregation_CONTROL)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (resolver *nodeResolver) FailingControls(ctx context.Context, args RawQuery) ([]*complianceControlResolver, error) {
	defer metrics.SetGraphQLOperationDurationTime(time.Now(), pkgMetrics.Nodes, "FailingControls")
	if err := readCompliance(ctx); err != nil {
		return nil, err
	}
	scope := []storage.ComplianceAggregation_Scope{storage.ComplianceAggregation_NODE, storage.ComplianceAggregation_CONTROL}
	results, err := resolver.getNodeLastSuccessfulComplianceRunAggregatedResult(ctx, scope, args)
	if err != nil {
		return nil, err
	}
	resolvers, err := resolver.root.wrapComplianceControls(getComplianceControlsFromAggregationResults(results, failing, resolver.root.ComplianceStandardStore))
	if err != nil {
		return nil, err
	}
	return resolvers, nil
}

func (resolver *nodeResolver) PassingControls(ctx context.Context, args RawQuery) ([]*complianceControlResolver, error) {
	defer metrics.SetGraphQLOperationDurationTime(time.Now(), pkgMetrics.Nodes, "PassingControls")
	if err := readCompliance(ctx); err != nil {
		return nil, err
	}
	scope := []storage.ComplianceAggregation_Scope{storage.ComplianceAggregation_NODE, storage.ComplianceAggregation_CONTROL}
	results, err := resolver.getNodeLastSuccessfulComplianceRunAggregatedResult(ctx, scope, args)
	if err != nil {
		return nil, err
	}
	resolvers, err := resolver.root.wrapComplianceControls(getComplianceControlsFromAggregationResults(results, passing, resolver.root.ComplianceStandardStore))
	if err != nil {
		return nil, err
	}
	return resolvers, nil
}

func (resolver *nodeResolver) Controls(ctx context.Context, args RawQuery) ([]*complianceControlResolver, error) {
	defer metrics.SetGraphQLOperationDurationTime(time.Now(), pkgMetrics.Nodes, "Controls")
	if err := readCompliance(ctx); err != nil {
		return nil, err
	}
	scope := []storage.ComplianceAggregation_Scope{storage.ComplianceAggregation_NODE, storage.ComplianceAggregation_CONTROL}
	results, err := resolver.getNodeLastSuccessfulComplianceRunAggregatedResult(ctx, scope, args)
	if err != nil {
		return nil, err
	}
	resolvers, err := resolver.root.wrapComplianceControls(getComplianceControlsFromAggregationResults(results, any, resolver.root.ComplianceStandardStore))
	if err != nil {
		return nil, err
	}
	return resolvers, nil
}

func getComplianceControlsFromAggregationResults(results []*storage.ComplianceAggregation_Result, controlType resultType, cs complianceStandards.Repository) ([]*v1.ComplianceControl, error) {
	if cs == nil {
		return nil, errors.New("empty compliance standards store encountered: argument cs is nil")
	}
	var controls []*v1.ComplianceControl
	for _, r := range results {
		if (controlType == passing && r.GetNumPassing() == 0) || (controlType == failing && r.GetNumFailing() == 0) {
			continue
		}
		controlID, err := getScopeIDFromAggregationResult(r, storage.ComplianceAggregation_CONTROL)
		if err != nil {
			continue
		}
		control := cs.Control(controlID)
		if control == nil {
			continue
		}
		controls = append(controls, control)
	}
	return controls, nil
}

func getComplianceControlCountFromAggregationResults(results []*storage.ComplianceAggregation_Result) *complianceControlCountResolver {
	ret := &complianceControlCountResolver{}
	for _, r := range results {
		if r.GetNumFailing() != 0 {
			ret.failingCount++
		} else if r.GetNumPassing() != 0 {
			ret.passingCount++
		} else {
			ret.unknownCount++
		}
	}
	return ret
}

type complianceControlCountResolver struct {
	failingCount int32
	passingCount int32
	unknownCount int32
}

func (resolver *complianceControlCountResolver) FailingCount() int32 {
	return resolver.failingCount
}

func (resolver *complianceControlCountResolver) PassingCount() int32 {
	return resolver.passingCount
}

func (resolver *complianceControlCountResolver) UnknownCount() int32 {
	return resolver.unknownCount
}
