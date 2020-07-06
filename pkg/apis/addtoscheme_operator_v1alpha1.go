package apis

import (
	k8sv1alpha1 "github.com/tektoncd/operator/pkg/apis/kuberntes/operator/v1alpha1"
	openshiftv1alpha1 "github.com/tektoncd/operator/pkg/apis/openshift/operator/v1alpha1"
)

func init() {
	// Register the types with the Scheme so the components can map objects to GroupVersionKinds and back
	AddToSchemes = append(AddToSchemes, k8sv1alpha1.SchemeBuilder.AddToScheme)

	AddToSchemes = append(AddToSchemes, openshiftv1alpha1.SchemeBuilder.AddToScheme)
}
