package alerting_test

import (
	"context"
	"fmt"
	"github.com/prometheus/common/model"
	"github.com/rancher/opni/pkg/alerting/metrics"
	"github.com/rancher/opni/pkg/metrics/unmarshal"
	"github.com/rancher/opni/plugins/alerting/pkg/alerting"
	"github.com/rancher/opni/plugins/cortex/pkg/apis/cortexadmin"
	"math/rand"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/rancher/opni/pkg/alerting/condition"
	alertingv1alpha "github.com/rancher/opni/pkg/apis/alerting/v1alpha"
	managementv1 "github.com/rancher/opni/pkg/apis/management/v1"
	"github.com/rancher/opni/pkg/util/waitctx"

	"github.com/rancher/opni/pkg/test"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/emptypb"
)

var _ = Describe("Alerting Conditions integration tests", Ordered, Label(test.Unit, test.Slow), func() {
	ctx := context.Background()
	When("The alerting condition API is passed invalid input it should be robust", func() {
		Specify("Create Alert Condition API should be robust to invalid input", func() {
			toTestCreateCondition := []InvalidInputs{
				{
					req: &alertingv1alpha.AlertCondition{},
					err: fmt.Errorf("invalid input"),
				},
			}

			for _, invalidInput := range toTestCreateCondition {
				_, err := alertingClient.CreateAlertCondition(ctx, invalidInput.req.(*alertingv1alpha.AlertCondition))
				Expect(err).To(HaveOccurred())
			}

		})

		Specify("Get Alert Condition API should be robust to invalid input", func() {
			//TODO

		})

		Specify("Update Alert Condition API should be robust to invalid input", func() {
			//TODO
		})

		Specify("List Alert Condition API should be robust to invalid input", func() {
			//TODO
		})

		Specify("Delete Alert Condition API should be robust to invalid input", func() {
			//TODO
		})
	})

	When("We mock out backend metrics for alerting", func() {
		Specify("We should be able to mock out kubernetes pod metrics", func() {
			Expect(kubernetesTempMetricServerPort).NotTo(Equal(0))
			Expect(curTestState.mockPods).NotTo(BeEmpty())
			// non-deterministically sample mock pods to set
			mp := curTestState.mockPods[rand.Intn(len(curTestState.mockPods))]
			for _, state := range metrics.KubePodStates {
				Eventually(func() error {
					mp.phase = state
					setMockKubernetesPodState(kubernetesTempMetricServerPort, mp)
					resp, err := adminClient.Query(ctx, &cortexadmin.QueryRequest{
						Tenants: []string{"agent"},
						Query: fmt.Sprintf(
							"kube_pod_status_phase{pod=\"%s\",namespace=\"%s\",phase=\"%s\", uid=\"%s\"}",
							mp.podName,
							mp.namespace,
							state,
							mp.uid,
						),
					})
					if err != nil {
						return err
					}
					q, err := unmarshal.UnmarshalPrometheusResponse(resp.Data)
					if err != nil {
						return err
					}
					var v model.Vector
					switch q.V.Type() {
					case model.ValVector:
						v = q.V.(model.Vector)

					default:
						return fmt.Errorf("cannot unmarshal prometheus response into vector type")
					}
					if len(v) == 0 {
						return fmt.Errorf("no data found")
					}
					for _, sample := range v {
						if sample.Value != 1 {
							return fmt.Errorf("expected 1 sample, got %f", sample.Value)
						}
					}

					return nil
				}, time.Second*10, time.Second).Should(Succeed())
			}
		})
	})

	When("The alerting plugin starts...", func() {
		It("Should be able to CRUD [kubernetes] type alert conditions", func() {
			alertDuration := durationpb.New(time.Second * 30)
			// non-deterministically sample mock pods to set
			mp := curTestState.mockPods[rand.Intn(len(curTestState.mockPods))]
			resp, err := alertingClient.CreateAlertCondition(ctx, &alertingv1alpha.AlertCondition{
				Name:     "kubernetes-pod-stuck",
				Severity: alertingv1alpha.Severity_WARNING,
				Labels:   []string{mp.uid},
				AlertType: &alertingv1alpha.AlertTypeDetails{
					Type: &alertingv1alpha.AlertTypeDetails_KubeState{
						KubeState: &alertingv1alpha.AlertConditionKubeState{
							ClusterId: "agent",
							Object:    mp.podName,
							Namespace: mp.namespace,
							State:     mp.phase,
							For:       alertDuration,
						},
					},
				},
			})
			Expect(err).NotTo(HaveOccurred())
			Expect(resp).NotTo(BeNil())

			ruleGroupId := alerting.CortexRuleIdFromUuid(resp.Id)
			test.ExpectRuleGroupToExist(
				adminClient,
				ctx,
				"agent",
				ruleGroupId,
				time.Second,
				time.Second*15,
			)

			// check alert actually evaluates to true in cortex
			Eventually(func() error {
				resp, err := adminClient.Query(ctx, &cortexadmin.QueryRequest{
					Tenants: []string{"agent"},
					Query:   fmt.Sprintf("max(ALERTS{alert_name=\"%s\"}) or on() vector(0)", resp.Id),
				})
				if err != nil {
					return err
				}
				_, err = unmarshal.UnmarshalPrometheusResponse(resp.Data)
				if err != nil {
					return err
				}

				return nil
			}, alertDuration.AsDuration(), time.Second*5).Should(Succeed())

			// this code block checks that alert has been triggered in opni alerting
			// FIXME: cortex alerting rules are not yet connected to opni alerting
			//Eventually(func() error {
			//	filteredLogs, err := alertingClient.ListAlertLogs(ctx, &alertingv1alpha.ListAlertLogRequest{
			//		Labels: []string{mp.uid},
			//	})
			//	if err != nil {
			//		return err
			//	}
			//	if len(filteredLogs.Items) == 0 {
			//		return fmt.Errorf("no logs found")
			//	}
			//	return nil
			//}, alertDuration.AsDuration(), time.Second*10).Should(Succeed())
		})

		It("Should be able to CRUD [composition] type alert conditions", func() {
			// TODO : when implemented
		})

		It("Should be able to CRUD [control flow] type alert conditions", func() {
			// TODO: when implemented
		})

		XIt("Should be CRUD [system] type alert conditions", func() {
			conditions, err := alertingClient.ListAlertConditions(ctx, &alertingv1alpha.ListAlertConditionRequest{})
			Expect(err).To(Succeed())
			Expect(conditions.Items).To(HaveLen(0))

			client := env.NewManagementClient()
			token, err := client.CreateBootstrapToken(context.Background(), &managementv1.CreateBootstrapTokenRequest{
				Ttl: durationpb.New(time.Hour),
			})
			Expect(err).NotTo(HaveOccurred())
			info, err := client.CertsInfo(context.Background(), &emptypb.Empty{})
			Expect(err).To(Succeed())

			// on agent startup expect an alert condition to be created
			ctxca, ca := context.WithCancel(waitctx.FromContext(context.Background()))
			p, _ := env.StartAgent("agent", token, []string{info.Chain[len(info.Chain)-1].Fingerprint}, test.WithContext(ctxca))
			Expect(p).NotTo(Equal(0))
			time.Sleep(time.Second)
			newConds, err := alertingClient.ListAlertConditions(ctx, &alertingv1alpha.ListAlertConditionRequest{})
			Expect(err).To(Succeed())
			Expect(newConds.Items).To(HaveLen(1))

			// kill the agent
			ca()
			time.Sleep(time.Millisecond * 100)
			logs, err := alertingClient.ListAlertLogs(ctx, &alertingv1alpha.ListAlertLogRequest{})
			Expect(err).To(Succeed())
			Expect(logs.Items).NotTo(HaveLen(0))
			filteredLogs, err := alertingClient.ListAlertLogs(ctx, &alertingv1alpha.ListAlertLogRequest{
				Labels: condition.OpniDisconnect.Labels,
			})
			Expect(err).To(Succeed())
			Expect(filteredLogs.Items).ToNot(HaveLen(0))
		})
	})
})
