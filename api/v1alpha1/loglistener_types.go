package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type LogListenerSpec struct {
	Trigger string `json:"trigger"`
	// Size string `json:"-"`
}

// LogListenerStatus defines the observed state of LogListener
type LogListenerStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Trigger string `json:"trigger"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// LogListener is the Schema for the loglisteners API
type LogListener struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   LogListenerSpec   `json:"spec,omitempty"`
	Status LogListenerStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// LogListenerList contains a list of LogListener
type LogListenerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LogListener `json:"items"`
}

func init() {
	SchemeBuilder.Register(&LogListener{}, &LogListenerList{})
}
