/*
.
*/

package controllers

import (
	"context"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	runtanzuv1alpha1 "github.com/pivotal/telemetryrecorder/api/v1alpha1"
)

// TelemetryRecordReconciler reconciles a TelemetryRecord object
type TelemetryRecordReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=run.tanzu.vmware.com,resources=telemetryrecords,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=run.tanzu.vmware.com,resources=telemetryrecords/status,verbs=get;update;patch

func (r *TelemetryRecordReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("telemetryrecord", req.NamespacedName)

	var telemetryRecord runtanzuv1alpha1.TelemetryRecord
	r.Get(ctx, req.NamespacedName, &telemetryRecord)
	log.Info("Getting TelemetryRecord request", "Telemetry Record", telemetryRecord)

	return ctrl.Result{}, nil
}

func (r *TelemetryRecordReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&runtanzuv1alpha1.TelemetryRecord{}).
		Complete(r)
}
