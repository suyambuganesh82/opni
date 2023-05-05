// Code generated by cli_gen.go DO NOT EDIT.
// source: github.com/rancher/opni/plugins/metrics/pkg/apis/cortexops/cortexops.proto

package cortexops

import (
	context "context"
	cli "github.com/rancher/opni/internal/cli"
	v1 "github.com/rancher/opni/pkg/apis/storage/v1"
	flagutil "github.com/rancher/opni/pkg/util/flagutil"
	cobra "github.com/spf13/cobra"
	pflag "github.com/spf13/pflag"
	v2 "github.com/thediveo/enumflag/v2"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	strings "strings"
	time "time"
)

type contextKey_CortexOps_type struct{}

var contextKey_CortexOps contextKey_CortexOps_type

func ContextWithCortexOpsClient(ctx context.Context, client CortexOpsClient) context.Context {
	return context.WithValue(ctx, contextKey_CortexOps, client)
}

func CortexOpsClientFromContext(ctx context.Context) (CortexOpsClient, bool) {
	client, ok := ctx.Value(contextKey_CortexOps).(CortexOpsClient)
	return client, ok
}

func BuildCortexOpsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "ops",
		Short:             `The CortexOps service contains setup and configuration lifecycle actions for the managed Cortex cluster.`,
		Args:              cobra.NoArgs,
		ValidArgsFunction: cobra.NoFileCompletions,
	}

	cmd.AddCommand(BuildGetClusterConfigurationCmd())
	cmd.AddCommand(BuildConfigureClusterCmd())
	cmd.AddCommand(BuildGetClusterStatusCmd())
	cmd.AddCommand(BuildUninstallClusterCmd())
	cli.AddOutputFlag(cmd)
	return cmd
}

func BuildGetClusterConfigurationCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-configuration",
		Short: "Gets the current configuration of the managed Cortex cluster.",
		Long: `
HTTP handlers for this method:
- get:"/configuration"
`[1:],
		Args:              cobra.NoArgs,
		ValidArgsFunction: cobra.NoFileCompletions,
		RunE: func(cmd *cobra.Command, args []string) error {
			client, ok := CortexOpsClientFromContext(cmd.Context())
			if !ok {
				cmd.PrintErrln("failed to get client from context")
				return nil
			}
			response, err := client.GetClusterConfiguration(cmd.Context(), &emptypb.Empty{})
			if err != nil {
				return err
			}
			cli.RenderOutput(cmd, response)
			return nil
		},
	}
	return cmd
}

func BuildConfigureClusterCmd() *cobra.Command {
	input := &ClusterConfiguration{}
	cmd := &cobra.Command{
		Use:   "configure",
		Short: "Updates the configuration of the managed Cortex cluster to match the provided configuration.",
		Long: `
If the cluster is not installed, it will be configured and installed.
Otherwise, the already-installed cluster will be reconfigured.

Note: some fields may contain secrets. The placeholder value "***" can be used to
keep an existing secret when updating the cluster configuration.

HTTP handlers for this method:
- post:"/configure"  body:"*"
`[1:],
		Args:              cobra.NoArgs,
		ValidArgsFunction: cobra.NoFileCompletions,
		RunE: func(cmd *cobra.Command, args []string) error {
			client, ok := CortexOpsClientFromContext(cmd.Context())
			if !ok {
				cmd.PrintErrln("failed to get client from context")
				return nil
			}
			_, err := client.ConfigureCluster(cmd.Context(), input)
			if err != nil {
				return err
			}
			return nil
		},
	}
	cmd.Flags().AddFlagSet(input.FlagSet())
	cmd.RegisterFlagCompletionFunc("mode", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{"AllInOne", "HighlyAvailable"}, cobra.ShellCompDirectiveDefault
	})
	cmd.RegisterFlagCompletionFunc("storage.backend", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{"filesystem", "s3", "gcs", "azure", "swift"}, cobra.ShellCompDirectiveDefault
	})
	return cmd
}

func BuildGetClusterStatusCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "status",
		Short: "Gets the current status of the managed Cortex cluster.",
		Long: `
The status includes the current install state, version, and metadata. If
the cluster is in the process of being reconfigured or uninstalled, it will
be reflected in the install state.
No guarantees are made about the contents of the metadata field; its
contents are strictly informational.

HTTP handlers for this method:
- get:"/status"
`[1:],
		Args:              cobra.NoArgs,
		ValidArgsFunction: cobra.NoFileCompletions,
		RunE: func(cmd *cobra.Command, args []string) error {
			client, ok := CortexOpsClientFromContext(cmd.Context())
			if !ok {
				cmd.PrintErrln("failed to get client from context")
				return nil
			}
			response, err := client.GetClusterStatus(cmd.Context(), &emptypb.Empty{})
			if err != nil {
				return err
			}
			cli.RenderOutput(cmd, response)
			return nil
		},
	}
	return cmd
}

func BuildUninstallClusterCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "uninstall",
		Short: "Uninstalls the managed Cortex cluster.",
		Long: `
Implementation details including error handling and system state requirements
are left to the cluster driver, and this API makes no guarantees about
the state of the cluster after the call completes (regardless of success).

HTTP handlers for this method:
- post:"/uninstall"
`[1:],
		Args:              cobra.NoArgs,
		ValidArgsFunction: cobra.NoFileCompletions,
		RunE: func(cmd *cobra.Command, args []string) error {
			client, ok := CortexOpsClientFromContext(cmd.Context())
			if !ok {
				cmd.PrintErrln("failed to get client from context")
				return nil
			}
			_, err := client.UninstallCluster(cmd.Context(), &emptypb.Empty{})
			if err != nil {
				return err
			}
			return nil
		},
	}
	return cmd
}

func (input *ClusterConfiguration) FlagSet(prefix ...string) *pflag.FlagSet {
	fs := pflag.NewFlagSet("ClusterConfiguration", pflag.ExitOnError)
	fs.SortFlags = true
	fs.Var(v2.New(&input.Mode, "DeploymentMode", map[DeploymentMode][]string{
		DeploymentMode_AllInOne:        {"AllInOne"},
		DeploymentMode_HighlyAvailable: {"HighlyAvailable"},
	}, v2.EnumCaseSensitive), strings.Join(append(prefix, "mode"), "."), "The deployment mode to use for Cortex.")
	if input.Storage == nil {
		input.Storage = &v1.StorageSpec{}
	}
	fs.AddFlagSet(input.Storage.FlagSet(append(prefix, "storage")...))
	if input.Grafana == nil {
		input.Grafana = &GrafanaConfig{}
	}
	fs.AddFlagSet(input.Grafana.FlagSet(append(prefix, "grafana")...))
	if input.Workloads == nil {
		input.Workloads = &Workloads{}
	}
	fs.AddFlagSet(input.Workloads.FlagSet(append(prefix, "workloads")...))
	if input.Cortex == nil {
		input.Cortex = &CortexConfig{}
	}
	fs.AddFlagSet(input.Cortex.FlagSet(append(prefix, "cortex")...))
	return fs
}

func (input *ClusterConfiguration) RedactSecrets() {
	if input == nil {
		return
	}
	input.Storage.RedactSecrets()
}

func (input *ClusterConfiguration) UnredactSecrets(unredacted *ClusterConfiguration) error {
	if input == nil {
		return nil
	}
	if err := input.Storage.UnredactSecrets(unredacted.GetStorage()); err != nil {
		return err
	}
	return nil
}

func (input *GrafanaConfig) FlagSet(prefix ...string) *pflag.FlagSet {
	fs := pflag.NewFlagSet("GrafanaConfig", pflag.ExitOnError)
	fs.SortFlags = true
	fs.BoolVar(&input.Enabled, strings.Join(append(prefix, "enabled"), "."), false, "Whether to deploy a managed Grafana instance.")
	fs.StringVar(&input.Hostname, strings.Join(append(prefix, "hostname"), "."), "", "DNS name at which Grafana will be available in the browser.")
	return fs
}

func (input *Workloads) FlagSet(prefix ...string) *pflag.FlagSet {
	fs := pflag.NewFlagSet("Workloads", pflag.ExitOnError)
	fs.SortFlags = true
	if input.Distributor == nil {
		input.Distributor = &CortexWorkloadSpec{}
	}
	fs.AddFlagSet(input.Distributor.FlagSet(append(prefix, "distributor")...))
	fs.Lookup(strings.Join(append(prefix, "distributor", "replicas"), ".")).DefValue = "1"
	if input.Ingester == nil {
		input.Ingester = &CortexWorkloadSpec{}
	}
	fs.AddFlagSet(input.Ingester.FlagSet(append(prefix, "ingester")...))
	if input.Compactor == nil {
		input.Compactor = &CortexWorkloadSpec{}
	}
	fs.AddFlagSet(input.Compactor.FlagSet(append(prefix, "compactor")...))
	if input.StoreGateway == nil {
		input.StoreGateway = &CortexWorkloadSpec{}
	}
	fs.AddFlagSet(input.StoreGateway.FlagSet(append(prefix, "store-gateway")...))
	if input.Ruler == nil {
		input.Ruler = &CortexWorkloadSpec{}
	}
	fs.AddFlagSet(input.Ruler.FlagSet(append(prefix, "ruler")...))
	if input.QueryFrontend == nil {
		input.QueryFrontend = &CortexWorkloadSpec{}
	}
	fs.AddFlagSet(input.QueryFrontend.FlagSet(append(prefix, "query-frontend")...))
	fs.Lookup(strings.Join(append(prefix, "query-frontend", "replicas"), ".")).DefValue = "1"
	if input.Querier == nil {
		input.Querier = &CortexWorkloadSpec{}
	}
	fs.AddFlagSet(input.Querier.FlagSet(append(prefix, "querier")...))
	if input.Purger == nil {
		input.Purger = &CortexWorkloadSpec{}
	}
	fs.AddFlagSet(input.Purger.FlagSet(append(prefix, "purger")...))
	fs.Lookup(strings.Join(append(prefix, "purger", "replicas"), ".")).DefValue = "1"
	return fs
}

func (input *CortexWorkloadSpec) FlagSet(prefix ...string) *pflag.FlagSet {
	fs := pflag.NewFlagSet("CortexWorkloadSpec", pflag.ExitOnError)
	fs.SortFlags = true
	fs.Int32Var(&input.Replicas, strings.Join(append(prefix, "replicas"), "."), 0, "Number of replicas to run for this workload. Should be an odd number.")
	fs.StringSliceVar(&input.ExtraArgs, strings.Join(append(prefix, "extra-args"), "."), nil, "Any additional arguments to pass to Cortex.")
	return fs
}

func (input *CortexConfig) FlagSet(prefix ...string) *pflag.FlagSet {
	fs := pflag.NewFlagSet("CortexConfig", pflag.ExitOnError)
	fs.SortFlags = true
	if input.Compactor == nil {
		input.Compactor = &CompactorConfig{}
	}
	fs.AddFlagSet(input.Compactor.FlagSet(append(prefix, "compactor")...))
	if input.Querier == nil {
		input.Querier = &QuerierConfig{}
	}
	fs.AddFlagSet(input.Querier.FlagSet(append(prefix, "querier")...))
	if input.Distributor == nil {
		input.Distributor = &DistributorConfig{}
	}
	fs.AddFlagSet(input.Distributor.FlagSet(append(prefix, "distributor")...))
	if input.Ingester == nil {
		input.Ingester = &IngesterConfig{}
	}
	fs.AddFlagSet(input.Ingester.FlagSet(append(prefix, "ingester")...))
	return fs
}

func (input *CompactorConfig) FlagSet(prefix ...string) *pflag.FlagSet {
	fs := pflag.NewFlagSet("CompactorConfig", pflag.ExitOnError)
	fs.SortFlags = true
	fs.Var(flagutil.DurationpbSliceValue([]time.Duration{2 * time.Hour, 12 * time.Hour, 24 * time.Hour}, &input.BlockRanges), strings.Join(append(prefix, "block-ranges"), "."), "List of compaction time ranges")
	fs.Var(flagutil.DurationpbValue(1*time.Hour, &input.CompactionInterval), strings.Join(append(prefix, "compaction-interval"), "."), "The frequency at which the compaction runs")
	fs.Var(flagutil.DurationpbValue(15*time.Minute, &input.CleanupInterval), strings.Join(append(prefix, "cleanup-interval"), "."), "How frequently compactor should run blocks cleanup and maintenance, as well as update the bucket index")
	fs.Var(flagutil.DurationpbValue(12*time.Hour, &input.DeletionDelay), strings.Join(append(prefix, "deletion-delay"), "."), "Time before a block marked for deletion is deleted from the bucket")
	fs.Var(flagutil.DurationpbValue(6*time.Hour, &input.TenantCleanupDelay), strings.Join(append(prefix, "tenant-cleanup-delay"), "."), "For tenants marked for deletion, this is time between deleting of last block, and doing final cleanup (marker files, debug files) of the tenant")
	return fs
}

func (input *QuerierConfig) FlagSet(prefix ...string) *pflag.FlagSet {
	fs := pflag.NewFlagSet("QuerierConfig", pflag.ExitOnError)
	fs.SortFlags = true
	fs.Var(flagutil.DurationpbValue(2*time.Minute, &input.QueryTimeout), strings.Join(append(prefix, "query-timeout"), "."), "The timeout for a query")
	fs.Int32Var(&input.MaxSamples, strings.Join(append(prefix, "max-samples"), "."), 50e6, "Maximum number of samples a single query can load into memory")
	fs.Var(flagutil.DurationpbValue(0, &input.QueryIngestersWithin), strings.Join(append(prefix, "query-ingesters-within"), "."), "Maximum lookback beyond which queries are not sent to ingester. 0 means all queries are sent to ingester.")
	fs.Var(flagutil.DurationpbValue(10*time.Minute, &input.MaxQueryIntoFuture), strings.Join(append(prefix, "max-query-into-future"), "."), "Maximum duration into the future you can query. 0 to disable")
	fs.Var(flagutil.DurationpbValue(1*time.Minute, &input.DefaultEvaluationInterval), strings.Join(append(prefix, "default-evaluation-interval"), "."), "The default evaluation interval or step size for subqueries")
	fs.Var(flagutil.DurationpbValue(0, &input.QueryStoreAfter), strings.Join(append(prefix, "query-store-after"), "."), "The time after which a metric should be queried from storage and not just ingesters. 0 means all queries are sent to store.  When running the blocks storage, if this option is enabled, the time range of the query sent to the store will be manipulated  to ensure the query end is not more recent than 'now - query-store-after'.")
	fs.Var(flagutil.DurationpbValue(5*time.Minute, &input.LookbackDelta), strings.Join(append(prefix, "lookback-delta"), "."), "Time since the last sample after which a time series is considered stale and ignored by expression evaluations")
	fs.Var(flagutil.DurationpbValue(0, &input.ShuffleShardingIngestersLookbackPeriod), strings.Join(append(prefix, "shuffle-sharding-ingesters-lookback-period"), "."), "When distributor's sharding strategy is shuffle-sharding and this setting is > 0, queriers fetch in-memory series from  the minimum set of required ingesters, selecting only ingesters which may have received series since 'now - lookback period'.  The lookback period should be greater or equal than the configured 'query store after' and 'query ingesters within'.  If this setting is 0, queriers always query all ingesters (ingesters shuffle sharding on read path is disabled).")
	fs.Int32Var(&input.MaxFetchedSeriesPerQuery, strings.Join(append(prefix, "max-fetched-series-per-query"), "."), 0, "The maximum number of unique series for which a query can fetch samples from each ingesters and blocks storage. This limit is enforced in the querier, ruler and store-gateway. 0 to disable")
	return fs
}

func (input *DistributorConfig) FlagSet(prefix ...string) *pflag.FlagSet {
	fs := pflag.NewFlagSet("DistributorConfig", pflag.ExitOnError)
	fs.SortFlags = true
	fs.Float64Var(&input.IngestionRate, strings.Join(append(prefix, "ingestion-rate"), "."), 600000, "Per-user ingestion rate limit in samples per second.")
	fs.StringVar(&input.IngestionRateStrategy, strings.Join(append(prefix, "ingestion-rate-strategy"), "."), "local", "Whether the ingestion rate limit should be applied individually to each distributor instance (local), or evenly shared across the cluster (global).")
	fs.Int32Var(&input.IngestionBurstSize, strings.Join(append(prefix, "ingestion-burst-size"), "."), 1000000, "Per-user allowed ingestion burst size (in number of samples).")
	return fs
}

func (input *IngesterConfig) FlagSet(prefix ...string) *pflag.FlagSet {
	fs := pflag.NewFlagSet("IngesterConfig", pflag.ExitOnError)
	fs.SortFlags = true
	fs.Int32Var(&input.MaxLocalSeriesPerUser, strings.Join(append(prefix, "max-local-series-per-user"), "."), 0, "The maximum number of active series per user, per ingester. 0 to disable.")
	fs.Int32Var(&input.MaxLocalSeriesPerMetric, strings.Join(append(prefix, "max-local-series-per-metric"), "."), 0, "The maximum number of active series per metric name, per ingester. 0 to disable.")
	fs.Int32Var(&input.MaxGlobalSeriesPerUser, strings.Join(append(prefix, "max-global-series-per-user"), "."), 0, "The maximum number of active series per user, across the cluster before replication. 0 to disable.")
	fs.Int32Var(&input.MaxGlobalSeriesPerMetric, strings.Join(append(prefix, "max-global-series-per-metric"), "."), 0, "The maximum number of active series per metric name, across the cluster before replication. 0 to disable.")
	fs.Int32Var(&input.MaxLocalMetricsWithMetadataPerUser, strings.Join(append(prefix, "max-local-metrics-with-metadata-per-user"), "."), 0, "The maximum number of active metrics with metadata per user, per ingester. 0 to disable.")
	fs.Int32Var(&input.MaxLocalMetadataPerMetric, strings.Join(append(prefix, "max-local-metadata-per-metric"), "."), 0, "The maximum number of metadata per metric, per ingester. 0 to disable.")
	fs.Int32Var(&input.MaxGlobalMetricsWithMetadataPerUser, strings.Join(append(prefix, "max-global-metrics-with-metadata-per-user"), "."), 0, "The maximum number of active metrics with metadata per user, across the cluster. 0 to disable.")
	fs.Int32Var(&input.MaxGlobalMetadataPerMetric, strings.Join(append(prefix, "max-global-metadata-per-metric"), "."), 0, "The maximum number of metadata per metric, across the cluster. 0 to disable.")
	return fs
}
