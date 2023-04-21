package controllers_test

import (
	"context"
	"fmt"

	opnimonitoringv1beta1 "github.com/rancher/opni/apis/monitoring/v1beta1"
	"github.com/rancher/opni/pkg/otel"

	. "github.com/kralicky/kmatch"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	promoperatorv1 "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1"
	opnicorev1beta1 "github.com/rancher/opni/apis/core/v1beta1"
	opniloggingv1beta1 "github.com/rancher/opni/apis/logging/v1beta1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

var _ = Describe("Core Collector Controller", Ordered, Label("controller", "slow"), func() {
	var (
		ns                  string
		loggingConfig       *opniloggingv1beta1.CollectorConfig
		monitoringConfig    *opnimonitoringv1beta1.CollectorConfig
		metricsCollectorObj *opnicorev1beta1.Collector
		loggingCollectorObj *opnicorev1beta1.Collector
	)
	When("creating a collector resource for monitoring", func() {
		It("should succeed in creating the objects", func() {
			ns = makeTestNamespace()
			monitoringConfig = &opnimonitoringv1beta1.CollectorConfig{
				ObjectMeta: metav1.ObjectMeta{
					Name: otel.MetricsCrdName,
				},
				Spec: opnimonitoringv1beta1.CollectorConfigSpec{
					RemoteWriteEndpoint: "http://test-endpoint",
				},
			}
			metricsCollectorObj = &opnicorev1beta1.Collector{
				ObjectMeta: metav1.ObjectMeta{
					Name: "test-monitoring",
				},
				Spec: opnicorev1beta1.CollectorSpec{
					SystemNamespace: ns,
					AgentEndpoint:   "http://test-endpoint",
					MetricsConfig: &corev1.LocalObjectReference{
						Name: otel.MetricsCrdName,
					},
				},
			}
			serviceMonitor := &promoperatorv1.ServiceMonitor{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "opni-kube-state-metrics",
					Namespace: ns,
				},
				Spec: promoperatorv1.ServiceMonitorSpec{
					Endpoints: []promoperatorv1.Endpoint{
						{
							HonorLabels: true,
							Port:        "http",
						},
					},
					JobLabel: "app.kubernetes.io/name",
					Selector: metav1.LabelSelector{
						MatchLabels: map[string]string{
							"app.kubernetes.io/instance": "opni",
							"app.kubernetes.io/name":     "kube-state-metrics",
						},
					},
				},
			}

			service := &corev1.Service{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "some-service",
					Namespace: ns,
					Labels: map[string]string{
						"app.kubernetes.io/instance": "opni",
						"app.kubernetes.io/name":     "kube-state-metrics",
					},
				},
				Spec: corev1.ServiceSpec{
					Selector: map[string]string{
						"test": "A",
					},
					Ports: []corev1.ServicePort{
						{
							Name:       "http",
							Port:       8080,
							TargetPort: intstr.FromString("http"),
						},
					},
				},
			}

			pods := &corev1.PodList{
				Items: []corev1.Pod{
					{
						ObjectMeta: metav1.ObjectMeta{
							Name:      "pod-1",
							Namespace: ns,
							Labels: map[string]string{
								"test": "A",
							},
						},
						Spec: corev1.PodSpec{
							Containers: []corev1.Container{
								{
									Name:  "idk",
									Image: "nginx",
									Ports: []corev1.ContainerPort{
										{
											Name:          "http",
											ContainerPort: 8080,
											Protocol:      "TCP",
										},
									},
								},
							},
						},
					},
				},
			}

			podMonitor := &promoperatorv1.PodMonitor{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "opni-kube-state-metrics",
					Namespace: ns,
				},
				Spec: promoperatorv1.PodMonitorSpec{
					Selector: metav1.LabelSelector{
						MatchLabels: map[string]string{
							"test": "A",
						},
					},
					PodMetricsEndpoints: []promoperatorv1.PodMetricsEndpoint{
						{
							Port: "http",
						},
					},
				},
			}

			By("creating the otel collector configs for metrics")
			Expect(k8sClient.Create(context.Background(), monitoringConfig)).To(Succeed())
			Expect(k8sClient.Create(context.Background(), metricsCollectorObj)).To(Succeed())

			By("creating physical objects for discovery")
			Expect(k8sClient.Create(context.Background(), service)).To(Succeed())
			for _, pod := range pods.Items {
				Expect(k8sClient.Create(context.Background(), &pod)).To(Succeed())
			}

			By("creating the prometheus CRDs used by the collector operator")
			Expect(k8sClient.Create(context.Background(), serviceMonitor)).To(Succeed())
			Expect(k8sClient.Create(context.Background(), podMonitor)).To(Succeed())
		})

		It("should create the monitoring infra", func() {
			By("checking the agent config map has the hostmetrics & kubelets deps")
			Eventually(Object(&corev1.ConfigMap{
				ObjectMeta: metav1.ObjectMeta{
					Name:      fmt.Sprintf("%s-agent-config", metricsCollectorObj.Name),
					Namespace: ns,
				},
			})).Should(ExistAnd(
				HaveData("receivers.yaml", nil, "config.yaml", nil),
				HaveOwner(metricsCollectorObj),
			))

			By("checking the the aggregator config map embeds scrape configs and selectors")
			By("checking aggregator configmap")
			Eventually(Object(&corev1.ConfigMap{
				ObjectMeta: metav1.ObjectMeta{
					Name:      fmt.Sprintf("%s-aggregator-config", metricsCollectorObj.Name),
					Namespace: ns,
				},
			})).Should(ExistAnd(
				HaveData("aggregator.yaml", nil),
				HaveOwner(metricsCollectorObj),
			))
			By("checking daemonset exists and has hostmetrics requirements")
			Eventually(Object(&appsv1.DaemonSet{
				ObjectMeta: metav1.ObjectMeta{
					Name:      fmt.Sprintf("%s-collector-agent", metricsCollectorObj.Name),
					Namespace: ns,
				},
			})).Should(ExistAnd(
				HaveMatchingContainer(And(
					HaveName("otel-collector"),
					HaveVolumeMounts("collector-config"),
				)),
				HaveMatchingVolume(And(
					HaveVolumeSource("ConfigMap"),
					HaveName("collector-config"),
				)),
				HaveMatchingVolume(And(
					HaveVolumeSource("HostPath"),
					HaveName("root"),
				)),
				HaveOwner(metricsCollectorObj),
			))
			By("checking aggregator deployment")
			Eventually(Object(&appsv1.Deployment{
				ObjectMeta: metav1.ObjectMeta{
					Name:      fmt.Sprintf("%s-collector-aggregator", metricsCollectorObj.Name),
					Namespace: ns,
				},
			})).Should(ExistAnd(
				HaveMatchingContainer(And(
					HaveName("otel-collector"),
					HaveVolumeMounts("collector-config"),
				)),
				HaveMatchingVolume(And(
					HaveVolumeSource("ConfigMap"),
					HaveName("collector-config"),
				)),
				HaveOwner(metricsCollectorObj),
			))
			By("checking the metrics tls assets secrets exists")
			Eventually(Object(&corev1.Secret{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "opni-otel-tls-assets",
					Namespace: ns,
				},
			})).Should(ExistAnd(
				HaveOwner(metricsCollectorObj),
			))
		})
		//TODO : wip
		XIt("should mount required tls assets for service monitors based on prometheus crds", func() {
			By("creating a user workload & user defined secret for TLS")
			secretName := "test-tls-secret"
			tlsSecret := &corev1.Secret{
				ObjectMeta: metav1.ObjectMeta{
					Name:      secretName,
					Namespace: ns,
				},
				Data: map[string][]byte{
					"ca.crt":  []byte("test-ca"),
					"tls.crt": []byte("test-cert"),
					"tls.key": []byte("test-key"),
				},
			}

			service := &corev1.Service{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "tls-service",
					Namespace: ns,
					Labels: map[string]string{
						"app.kubernetes.io/instance": "opni",
						"app.kubernetes.io/name":     "kube-state-metrics",
					},
				},
				Spec: corev1.ServiceSpec{
					Selector: map[string]string{
						"test": "tls",
					},
					Ports: []corev1.ServicePort{
						{
							Name:       "http-metrics",
							Port:       8080,
							TargetPort: intstr.FromString("http-metrics"),
						},
					},
				},
			}

			deploy := &appsv1.Deployment{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "tls-deployment",
					Namespace: ns,
					Labels: map[string]string{
						//TODO
						"test": "tls",
					},
				},
				Spec: appsv1.DeploymentSpec{
					Selector: &metav1.LabelSelector{
						MatchLabels: map[string]string{
							"test": "tls",
						},
					},
					Template: corev1.PodTemplateSpec{
						ObjectMeta: metav1.ObjectMeta{
							Name:      "tls-pod",
							Namespace: ns,
							Labels: map[string]string{
								"test": "tls",
							},
						},
						Spec: corev1.PodSpec{
							Containers: []corev1.Container{
								{
									Name:  "tls-container",
									Image: "nginx",
									Ports: []corev1.ContainerPort{
										{
											Name:          "http-metrics",
											ContainerPort: 8080,
										},
									},
								},
							},
						},
					},
				},
			}
			Expect(k8sClient.Create(context.Background(), tlsSecret)).To(Succeed())
			Expect(k8sClient.Create(context.Background(), service)).To(Succeed())
			Expect(k8sClient.Create(context.Background(), deploy)).To(Succeed())

			By("creating a service monitor with the secret TLS config using safe TLS")
			serviceMonitorWithSecretTLS := &promoperatorv1.ServiceMonitor{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "tls-service-monitor",
					Namespace: ns,
				},
				Spec: promoperatorv1.ServiceMonitorSpec{
					Endpoints: []promoperatorv1.Endpoint{
						{
							Path:   "/metrics",
							Port:   "http-metrics",
							Scheme: "https",
							TLSConfig: &promoperatorv1.TLSConfig{
								SafeTLSConfig: promoperatorv1.SafeTLSConfig{
									CA: promoperatorv1.SecretOrConfigMap{
										Secret: &corev1.SecretKeySelector{
											LocalObjectReference: corev1.LocalObjectReference{
												Name: secretName,
											},
											Key: "ca.crt",
										},
									},
									Cert: promoperatorv1.SecretOrConfigMap{
										Secret: &corev1.SecretKeySelector{
											LocalObjectReference: corev1.LocalObjectReference{
												Name: secretName,
											},
											Key: "tls.crt",
										},
									},
									KeySecret: &corev1.SecretKeySelector{
										LocalObjectReference: corev1.LocalObjectReference{
											Name: secretName,
										},
										Key: "tls.key",
									},
								},
							},
						},
					},
					NamespaceSelector: promoperatorv1.NamespaceSelector{
						MatchNames: []string{ns},
					},
					Selector: metav1.LabelSelector{
						MatchLabels: map[string]string{
							"test": "tls",
						},
					},
				},
			}
			Expect(k8sClient.Create(context.Background(), serviceMonitorWithSecretTLS)).To(Succeed())

			Eventually(Object(&corev1.Secret{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "opni-otel-tls-assets",
					Namespace: ns,
				},
			})).Should(ExistAnd(
				HaveOwner(metricsCollectorObj),
				HaveData(
					"ca.crt", "test-ca",
					"tls.crt", "test-cert",
					"tls.key", "test-key",
				),
			))

			By("creating a user defined configmap for TLS")
			// TODO

			By("creating a service monitor with the configmap TLS config using safe TLS")
		})

		It("should mount required tls assets for pod monitors", func() {
			By("using safe TLS config")
		})
	})

	When("creating a collector resource for logging", func() {
		It("should succeed in creating the objects", func() {
			ns = makeTestNamespace()
			loggingConfig = &opniloggingv1beta1.CollectorConfig{
				ObjectMeta: metav1.ObjectMeta{
					Name: "test-logging-config",
				},
				Spec: opniloggingv1beta1.CollectorConfigSpec{
					Provider: opniloggingv1beta1.LogProviderGeneric,
				},
			}
			loggingCollectorObj = &opnicorev1beta1.Collector{
				ObjectMeta: metav1.ObjectMeta{
					Name: "test-logging",
				},
				Spec: opnicorev1beta1.CollectorSpec{
					SystemNamespace: ns,
					AgentEndpoint:   "http://test-endpoint",
					LoggingConfig: &corev1.LocalObjectReference{
						Name: loggingConfig.Name,
					},
				},
			}
			Expect(k8sClient.Create(context.Background(), loggingConfig)).To(Succeed())
			Expect(k8sClient.Create(context.Background(), loggingCollectorObj)).To(Succeed())
		})
		// TODO: handle non happy path cases
		It("should create log scraping infra", func() {
			By("checking agent configmap")
			Eventually(Object(&corev1.ConfigMap{
				ObjectMeta: metav1.ObjectMeta{
					Name:      fmt.Sprintf("%s-agent-config", loggingCollectorObj.Name),
					Namespace: ns,
				},
			})).Should(ExistAnd(
				HaveData("receivers.yaml", nil, "config.yaml", nil),
				HaveOwner(loggingCollectorObj),
			))
			By("checking aggregator configmap")
			Eventually(Object(&corev1.ConfigMap{
				ObjectMeta: metav1.ObjectMeta{
					Name:      fmt.Sprintf("%s-aggregator-config", loggingCollectorObj.Name),
					Namespace: ns,
				},
			})).Should(ExistAnd(
				HaveData("aggregator.yaml", nil),
				HaveOwner(loggingCollectorObj),
			))
			By("checking daemonset")
			Eventually(Object(&appsv1.DaemonSet{
				ObjectMeta: metav1.ObjectMeta{
					Name:      fmt.Sprintf("%s-collector-agent", loggingCollectorObj.Name),
					Namespace: ns,
				},
			})).Should(ExistAnd(
				HaveMatchingContainer(And(
					HaveName("otel-collector"),
					HaveVolumeMounts("collector-config", "varlogpods"),
				)),
				HaveMatchingVolume(And(
					HaveVolumeSource("ConfigMap"),
					HaveName("collector-config"),
				)),
				HaveMatchingVolume(And(
					HaveVolumeSource("HostPath"),
					HaveName("varlogpods"),
				)),
				HaveOwner(loggingCollectorObj),
			))
			By("checking deployment")
			Eventually(Object(&appsv1.Deployment{
				ObjectMeta: metav1.ObjectMeta{
					Name:      fmt.Sprintf("%s-collector-aggregator", loggingCollectorObj.Name),
					Namespace: ns,
				},
			})).Should(ExistAnd(
				HaveMatchingContainer(And(
					HaveName("otel-collector"),
					HaveVolumeMounts("collector-config"),
				)),
				HaveMatchingVolume(And(
					HaveVolumeSource("ConfigMap"),
					HaveName("collector-config"),
				)),
				HaveOwner(loggingCollectorObj),
			))
		})
	})
})
