package integration_test

import (
	"context"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	managementv1 "github.com/rancher/opni/pkg/apis/management/v1"
	"github.com/rancher/opni/pkg/test"
	"github.com/rancher/opni/pkg/util"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/emptypb"
)

var _ = Describe("Gateway unary interceptor tests", Ordered, Label("integration"), func() {
	var env *test.Environment
	var client managementv1.ManagementClient
	var ctx context.Context
	BeforeAll(func() {
		env = &test.Environment{
			TestBin: "../../../testbin/bin",
		}
		Expect(env.Start()).To(Succeed())
		client = env.NewManagementClient(test.WithClientCaching("10Mi", time.Minute))
		ctx = env.Context()
	})

	When("we use the gateway caching interceptors", func() {
		It("should allow clients to opt-in to caching", func() {
			By("making a request with caching")
			clusterList, err := client.ListClusters(
				util.WithGrpcClientCaching(ctx, 15*time.Second),
				&managementv1.ListClustersRequest{},
			)
			Expect(err).To(Succeed())
			Expect(clusterList).NotTo(BeNil())
			Expect(clusterList.Items).To(HaveLen(0))

			By("bootstrapping an agent")
			certsInfo, err := client.CertsInfo(context.Background(), &emptypb.Empty{})
			Expect(err).NotTo(HaveOccurred())
			fingerprint := certsInfo.Chain[len(certsInfo.Chain)-1].Fingerprint
			Expect(fingerprint).NotTo(BeEmpty())
			token, err := client.CreateBootstrapToken(context.Background(), &managementv1.CreateBootstrapTokenRequest{
				Ttl: durationpb.New(time.Minute),
			})
			Expect(err).NotTo(HaveOccurred())
			env.StartAgent("test-cluster-id", token, []string{fingerprint})

			By("having the client being able to bypass the cache")
			Eventually(func() int {
				clusterList, err = client.ListClusters(
					// ctx,
					util.WithBypassCache(ctx),
					&managementv1.ListClustersRequest{},
				)
				if err != nil {
					return -1
				}
				return len(clusterList.Items)
			}, time.Second*5, time.Second*1).Should(Equal(1))

			By("verifying that the cache is still valid")
			clusterList, err = client.ListClusters(
				ctx,
				&managementv1.ListClustersRequest{},
			)
			Expect(err).To(Succeed())
			Expect(clusterList).NotTo(BeNil())
			Expect(clusterList.Items).To(HaveLen(0))

			By("verifying that the cache expires after the TTL")

			Eventually(func() int {
				return len(util.Must(client.ListClusters(
					ctx,
					&managementv1.ListClustersRequest{})).Items,
				)
			}, 15*time.Second, time.Second).Should(Equal(1))
		})
	})
})
