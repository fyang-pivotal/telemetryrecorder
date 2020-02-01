/*
.
*/

package controllers

import (
	"context"
	"github.com/go-logr/logr"
	runtanzuv1alpha1 "github.com/pivotal/telemetryrecorder/api/v1alpha1"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// TelemetryRecordReconciler reconciles a TelemetryRecord object
type TelemetryRecordReconciler struct {
	client.Client
	Log           logr.Logger
	Scheme        *runtime.Scheme
	DynamicClient dynamic.Interface
}

// +kubebuilder:rbac:groups=run.tanzu.vmware.com,resources=telemetryrecords,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=run.tanzu.vmware.com,resources=telemetryrecords/status,verbs=get;update;patch

func (r *TelemetryRecordReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("telemetryrecord", req.NamespacedName)

	var telemetryRecord runtanzuv1alpha1.TelemetryRecord
	r.Get(ctx, req.NamespacedName, &telemetryRecord)
	log.Info("Getting TelemetryRecord request", "Telemetry Record", telemetryRecord)

	targetResource := schema.GroupVersionResource{Group: telemetryRecord.Spec.Apigroup, Version: telemetryRecord.Spec.ApiVersion, Resource: telemetryRecord.Spec.ResourceName}

	got, _ := r.DynamicClient.Resource(targetResource).Namespace(telemetryRecord.Spec.Namespaced).List(metav1.ListOptions{})
	log.Info("Getting dynamic client list", "crd", got)

	podList := v1.PodList{}
	r.List(ctx, &podList)
	log.Info("Getting pod list", "pods", podList)

	return ctrl.Result{}, nil
}

func (r *TelemetryRecordReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&runtanzuv1alpha1.TelemetryRecord{}).
		Complete(r)
}
