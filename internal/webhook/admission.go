package webhook

import (
	v1 "github/setera/pkg/api/setera.com/v1"

	admissionv1 "k8s.io/api/admission/v1"
	corev1 "k8s.io/api/core/v1"

	"k8s.io/apimachinery/pkg/runtime"
	//"k8s.io/apimachinery/pkg/runtime/serializer"
)

var (
	runtimeScheme = runtime.NewScheme()
	//codecFactory  = serializer.NewCodecFactory(runtimeScheme)
	//deserializer  = codecFactory.UniversalDeserializer()
)

func init() {

	corev1.AddToScheme(runtimeScheme)
	admissionv1.AddToScheme(runtimeScheme)
	v1.AddToScheme(runtimeScheme)

}
