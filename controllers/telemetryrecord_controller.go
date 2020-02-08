/*
.
*/

package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-logr/logr"
	runtanzuv1alpha1 "github.com/pivotal/telemetryrecorder/api/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
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

	targetResource := schema.GroupVersionResource{
		Group:    telemetryRecord.Spec.ApiGroup,
		Version:  telemetryRecord.Spec.ApiVersion,
		Resource: telemetryRecord.Spec.ResourceName,
	}

	got, _ := r.DynamicClient.Resource(targetResource).Namespace("default").List(metav1.ListOptions{})
	log.Info("Getting dynamic client list", "crd", got, "targetResource", targetResource)

	// List target resources see https://github.com/kubernetes/client-go/blob/5be5d5753fd2079867ab0adc3f01f3b1c9337ecf/examples/dynamic-create-update-delete-deployment/main.go#L168
	res := make([]map[string]interface{}, 0)
	for _, item := range got.Items {
		crdValues := make(map[string]interface{})
		for _, field := range telemetryRecord.Spec.Fields {
			s, _, _ := unstructured.NestedFieldCopy(item.Object, "spec", "output_properties", field)
			crdValues[field] = s
		}
		res = append(res, crdValues)
	}

	bytes, _ := json.MarshalIndent(res, "", "    ")
	fmt.Printf("\n ================ \n Instrumented values are %v \n ================ \n \n ", string(bytes))
	return ctrl.Result{}, nil
}

func (r *TelemetryRecordReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&runtanzuv1alpha1.TelemetryRecord{}).
		Complete(r)
}
