package v1alpha2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

const (
	version   = "v1alpha2"
	groupName = "devexp.runtime.redhat.com"
)

var (
	GroupName = groupName
	// SchemeGroupVersion is the group version used to register these objects.
	GroupVersion  = schema.GroupVersion{Group: GroupName, Version: version}
	schemeBuilder = runtime.NewSchemeBuilder(addKnownTypes)
	// Install is a function which adds this version to a scheme
	Install = schemeBuilder.AddToScheme
)

// Adds the list of known types to api.Scheme.
func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(GroupVersion,
		&Component{},
		&ComponentList{},
		&Link{},
		&LinkList{},
		&Capability{},
		&CapabilityList{},
	)
	metav1.AddToGroupVersion(scheme, GroupVersion)
	return nil
}
