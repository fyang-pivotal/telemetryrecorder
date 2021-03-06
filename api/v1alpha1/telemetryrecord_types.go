/*
.
*/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// TelemetryRecordSpec defines the desired state of TelemetryRecord
type TelemetryRecordSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	ApiGroup     string   `json:"apiGroup,omitempty"`
	ApiVersion   string   `json:"apiVersion,omitempty"`
	ResourceName string   `json:"resourceName,omitempty"`
	Namespaced   bool     `json:"namespaced,omitempty"`
	Fields       []string `json:"fields,omitempty"`
}

// TelemetryRecordStatus defines the observed state of TelemetryRecord
type TelemetryRecordStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true

// TelemetryRecord is the Schema for the telemetryrecords API
type TelemetryRecord struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TelemetryRecordSpec   `json:"spec,omitempty"`
	Status TelemetryRecordStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TelemetryRecordList contains a list of TelemetryRecord
type TelemetryRecordList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TelemetryRecord `json:"items"`
}

func init() {
	SchemeBuilder.Register(&TelemetryRecord{}, &TelemetryRecordList{})
}
