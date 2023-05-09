package kubernetes_test

import (
	"context"
	"testing"
	"time"

	"github.com/kralicky/kmatch"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/rancher/opni/apis"
	"github.com/rancher/opni/pkg/test"
	_ "github.com/rancher/opni/pkg/test/setup"
	"github.com/rancher/opni/pkg/test/testk8s"
	"github.com/rancher/opni/pkg/test/testruntime"
	"github.com/rancher/opni/pkg/util/waitctx"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/rest"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	k8sClient  client.Client
	restConfig *rest.Config
	scheme     *runtime.Scheme
	k8sManager ctrl.Manager
)

func TestAPIs(t *testing.T) {
	SetDefaultEventuallyTimeout(30 * time.Second)
	SetDefaultEventuallyPollingInterval(100 * time.Millisecond)
	SetDefaultConsistentlyDuration(2 * time.Second)
	SetDefaultConsistentlyPollingInterval(100 * time.Millisecond)
	RegisterFailHandler(Fail)
	RunSpecs(t, "Logging Plugin Suite")
}

var _ = BeforeSuite(func() {
	testruntime.IfIntegration(func() {
		env := test.Environment{}

		ctx, ca := context.WithCancel(waitctx.Background())
		var err error

		restConfig, scheme, err = testk8s.StartK8s(ctx, []string{}, apis.NewScheme())
		Expect(err).NotTo(HaveOccurred())

		k8sClient, err = client.New(restConfig, client.Options{
			Scheme: scheme,
		})
		Expect(err).NotTo(HaveOccurred())

		kmatch.SetDefaultObjectClient(k8sClient)

		k8sManager = testk8s.StartManager(ctx, restConfig, scheme)

		DeferCleanup(func() {
			By("tearing down the test environment")
			err := env.Stop()
			Expect(err).NotTo(HaveOccurred())

			ca()
			waitctx.Wait(ctx)
		})
	})
})
