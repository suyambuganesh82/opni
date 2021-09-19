/*
Copyright 2021.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"

	nvidiav1 "github.com/NVIDIA/gpu-operator/api/v1"
	v1beta1 "github.com/rancher/opni/apis/v1beta1"
	"github.com/rancher/opni/pkg/resources/gpuadapter"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

// GpuPolicyAdapterReconciler reconciles a GpuPolicyAdapter object
type GpuPolicyAdapterReconciler struct {
	client.Client
	scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=opni.io,resources=gpupolicyadapters,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=opni.io,resources=gpupolicyadapters/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=opni.io,resources=gpupolicyadapters/finalizers,verbs=update
//+kubebuilder:rbac:groups=nvidia.opni.io,resources=*,verbs=get;list;watch;create;update;patch;delete

func (r *GpuPolicyAdapterReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	// Look up the object from the request.
	gpuAdapter := v1beta1.GpuPolicyAdapter{}
	err := r.Get(ctx, req.NamespacedName, &gpuAdapter)
	if err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	return gpuadapter.ReconcileGPUAdapter(ctx, r.Client, &gpuAdapter)
}

// SetupWithManager sets up the controller with the Manager.
func (r *GpuPolicyAdapterReconciler) SetupWithManager(mgr ctrl.Manager) error {
	r.Client = mgr.GetClient()
	r.scheme = mgr.GetScheme()
	return ctrl.NewControllerManagedBy(mgr).
		For(&v1beta1.GpuPolicyAdapter{}).
		Owns(&nvidiav1.ClusterPolicy{}).
		Complete(r)
}
