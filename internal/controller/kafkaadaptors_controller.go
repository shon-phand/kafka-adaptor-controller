/*
Copyright 2025 ShonPhand.

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

package controller

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	apiv1alpha1 "github.com/shon-phand/kafka-adaptor-controller.git/api/v1alpha1"
)

// KafkaAdaptorsReconciler reconciles a KafkaAdaptors object
type KafkaAdaptorsReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=api.core.shon.io,resources=kafkaadaptors,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=api.core.shon.io,resources=kafkaadaptors/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=api.core.shon.io,resources=kafkaadaptors/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the KafkaAdaptors object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.19.1/pkg/reconcile
func (r *KafkaAdaptorsReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	l := log.FromContext(ctx)

	// TODO(user): your logic here

	ka := &apiv1alpha1.KafkaAdaptors{}

	if err := r.Get(ctx, req.NamespacedName, ka); err != nil {
		l.Error(err, "unable to fetch KafkaAdaptors")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}
	l.Info("KafkaAdaptors fetched", "KafkaAdaptors", ka)
	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *KafkaAdaptorsReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&apiv1alpha1.KafkaAdaptors{}).
		Named("kafkaadaptors").
		Complete(r)
}
