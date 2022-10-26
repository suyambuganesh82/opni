package main

import (
	"context"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rancher/opni/pkg/plugins"
	"github.com/rancher/opni/pkg/plugins/meta"
	"github.com/rancher/opni/pkg/tracing"
	"github.com/rancher/opni/pkg/util/waitctx"
	"github.com/rancher/opni/plugins/topology/pkg/topology/agent"
	"github.com/rancher/opni/plugins/topology/pkg/topology/gateway"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	ctx, ca := context.WithCancel(waitctx.Background())
	switch meta.PluginMode(os.Getenv("OPNI_PLUGIN_MODE")) {
	case meta.ModeUnknown, meta.ModeGateway:
		tracing.Configure("plugin_topology_gateway")
		plugins.Serve(gateway.Scheme(ctx))
	case meta.ModeAgent:
		tracing.Configure("plugin_topology_agent")
		plugins.Serve(agent.Scheme(ctx))
	}
	ca()
	waitctx.Wait(ctx, 5*time.Second)
}