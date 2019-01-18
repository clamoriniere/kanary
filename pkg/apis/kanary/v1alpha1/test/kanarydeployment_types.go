package v1alpha1_test

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	kanaryv1alpha1 "github.com/amadeusitgroup/kanary/pkg/apis/kanary/v1alpha1"
)

// NewKanaryDeployment returns new KanaryDeploymentInstance
func NewKanaryDeployment(name, namespace, serviceName string, replicas int32, options *NewKanaryDeploymentOptions) *kanaryv1alpha1.KanaryDeployment {
	kd := &kanaryv1alpha1.KanaryDeployment{
		TypeMeta: metav1.TypeMeta{
			Kind:       "KanaryDeployment",
			APIVersion: kanaryv1alpha1.SchemeGroupVersion.String(),
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
	}

	kd.Spec.Template.Spec.Selector = &metav1.LabelSelector{}
	kd.Spec.Template.Spec.Replicas = kanaryv1alpha1.NewInt32(replicas)
	kd.Spec.ServiceName = serviceName

	if options != nil {
		if options.StartTime != nil {
			kd.CreationTimestamp = *options.StartTime
		}
		if options.Scale != nil {
			kd.Spec.Scale = *options.Scale
		}
		if options.Traffic != nil {
			kd.Spec.Traffic = *options.Traffic
		}
		if options.Validation != nil {
			kd.Spec.Validation = *options.Validation
		}
		if options.Status != nil {
			kd.Status = *options.Status
		}
	}

	kd = kanaryv1alpha1.DefaultKanaryDeployment(kd)
	kd.Spec.ServiceName = serviceName

	return kd
}

// NewKanaryDeploymentOptions used to provide creation options
type NewKanaryDeploymentOptions struct {
	StartTime  *metav1.Time
	Scale      *kanaryv1alpha1.KanaryDeploymentSpecScale
	Traffic    *kanaryv1alpha1.KanaryDeploymentSpecTraffic
	Validation *kanaryv1alpha1.KanaryDeploymentSpecValidation
	Status     *kanaryv1alpha1.KanaryDeploymentStatus
}