// Code generated by internal/codegen. DO NOT EDIT.

// Code generated by cli_gen.go DO NOT EDIT.
// source: github.com/rancher/opni/internal/cortex/config/runtimeconfig/runtimeconfig.proto

package runtimeconfig

import (
	flagutil "github.com/rancher/opni/pkg/util/flagutil"
	pflag "github.com/spf13/pflag"
	proto "google.golang.org/protobuf/proto"
	strings "strings"
)

func (in *IngesterInstanceLimits) DeepCopyInto(out *IngesterInstanceLimits) {
	out.Reset()
	proto.Merge(out, in)
}

func (in *IngesterInstanceLimits) DeepCopy() *IngesterInstanceLimits {
	return proto.Clone(in).(*IngesterInstanceLimits)
}

func (in *KvMultiRuntimeConfig) DeepCopyInto(out *KvMultiRuntimeConfig) {
	out.Reset()
	proto.Merge(out, in)
}

func (in *KvMultiRuntimeConfig) DeepCopy() *KvMultiRuntimeConfig {
	return proto.Clone(in).(*KvMultiRuntimeConfig)
}

func (in *RuntimeConfigValues) DeepCopyInto(out *RuntimeConfigValues) {
	out.Reset()
	proto.Merge(out, in)
}

func (in *RuntimeConfigValues) DeepCopy() *RuntimeConfigValues {
	return proto.Clone(in).(*RuntimeConfigValues)
}

func (in *RuntimeConfigValues) FlagSet(prefix ...string) *pflag.FlagSet {
	fs := pflag.NewFlagSet("RuntimeConfigValues", pflag.ExitOnError)
	fs.SortFlags = true
	if in.MultiKvConfig == nil {
		in.MultiKvConfig = &KvMultiRuntimeConfig{}
	}
	fs.AddFlagSet(in.MultiKvConfig.FlagSet(append(prefix, "multi-kv-config")...))
	fs.Var(flagutil.BoolPtrValue(&in.IngesterStreamChunksWhenUsingBlocks), strings.Join(append(prefix, "ingester-stream-chunks-when-using-blocks"), "."), "")
	if in.IngesterLimits == nil {
		in.IngesterLimits = &IngesterInstanceLimits{}
	}
	fs.AddFlagSet(in.IngesterLimits.FlagSet(append(prefix, "ingester-limits")...))
	return fs
}

func (in *KvMultiRuntimeConfig) FlagSet(prefix ...string) *pflag.FlagSet {
	fs := pflag.NewFlagSet("KvMultiRuntimeConfig", pflag.ExitOnError)
	fs.SortFlags = true
	fs.Var(flagutil.StringPtrValue(&in.Primary), strings.Join(append(prefix, "primary"), "."), "")
	fs.Var(flagutil.BoolPtrValue(&in.MirrorEnabled), strings.Join(append(prefix, "mirror-enabled"), "."), "")
	return fs
}

func (in *IngesterInstanceLimits) FlagSet(prefix ...string) *pflag.FlagSet {
	fs := pflag.NewFlagSet("IngesterInstanceLimits", pflag.ExitOnError)
	fs.SortFlags = true
	fs.Var(flagutil.FloatPtrValue(&in.MaxIngestionRate), strings.Join(append(prefix, "max-ingestion-rate"), "."), "")
	fs.Var(flagutil.IntPtrValue(&in.MaxTenants), strings.Join(append(prefix, "max-tenants"), "."), "")
	fs.Var(flagutil.IntPtrValue(&in.MaxSeries), strings.Join(append(prefix, "max-series"), "."), "")
	fs.Var(flagutil.IntPtrValue(&in.MaxInflightPushRequests), strings.Join(append(prefix, "max-inflight-push-requests"), "."), "")
	return fs
}
