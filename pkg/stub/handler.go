package stub

import (
	"context"
	"fmt"

	"github.com/openshift-hive/hive-operator/pkg/apis/hive/v1alpha1"

	"github.com/operator-framework/operator-sdk/pkg/sdk"
	"github.com/sirupsen/logrus"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/intstr"
)

func NewHandler() sdk.Handler {
	return &Handler{}
}

type Handler struct {
	// Fill me
}

func (h *Handler) Handle(ctx context.Context, event sdk.Event) error {
	switch o := event.Object.(type) {
	case *v1alpha1.Hive:
		if event.Deleted {
			//deleted event logging
			logrus.Infof("Deleted event")
		} else {
			//create a deployment and then create
			//a service, sanity check!!
			logrus.Infof("event created/updated")
			//create the deployment
			dep := newHiveDeployment(o)
			err := sdk.Create(dep)
			if err != nil && !errors.IsAlreadyExists(err) {
				logrus.Errorf("Failed to create Deployment: %v", err)
				return err
			}
			err = sdk.Get(dep)
			if err != nil {
				return fmt.Errorf("failed to get deployment: %v", err)
			}
			//check if the spec size is the same
			size := o.Spec.Size
			if *dep.Spec.Replicas != size {
				dep.Spec.Replicas = &size
				err = sdk.Update(dep)
				if err != nil {
					return fmt.Errorf("failed to update deployment: %v", err)
				}
			}
			/*podNames := getPodNames(podList.Items)
			if !reflect.DeepEqual(podNames, memcached.Status.Nodes) {
				memcached.Status.Nodes = podNames
				err := sdk.Update(memcached)
				if err != nil {
					return fmt.Errorf("failed to update memcached status: %v", err)
				}
			}*/
			//create the service
			err = sdk.Create(newHiveService(o))
			if err != nil && !errors.IsAlreadyExists(err) {
				logrus.Errorf("Failed to create Hive service: %v", err)
				return err
			}
		}
	}
	return nil
}

//Deployment with 3 replicas
func newHiveDeployment(cr *v1alpha1.Hive) *appsv1.Deployment {
	labels := map[string]string{
		"app": "hive-operator",
	}
	replicas := cr.Spec.Size
	//deployment present in apps/v1 not corev1
	//need metav1 for including the TypeMeta, ObjectMeta
	return &appsv1.Deployment{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Deployment",
			APIVersion: "apps/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "hive-deployment",
			Namespace: cr.Namespace,
			OwnerReferences: []metav1.OwnerReference{
				*metav1.NewControllerRef(cr, schema.GroupVersionKind{
					Group:   v1alpha1.SchemeGroupVersion.Group,
					Version: v1alpha1.SchemeGroupVersion.Version,
					Kind:    "Hive",
				}),
			},
			Labels: labels,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: labels,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: labels,
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{{
						Image: "luksa/kubia:v2",
						Name:  "hive-operator",
					}},
				},
			},
		},
	}
}

//Service
func newHiveService(cr *v1alpha1.Hive) *corev1.Service {
	labels := map[string]string{
		"app": "hive-operator",
	}
	//service is present in corev1 just like pod
	return &corev1.Service{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Service",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "hive-service",
			Namespace: cr.Namespace,
			OwnerReferences: []metav1.OwnerReference{
				*metav1.NewControllerRef(cr, schema.GroupVersionKind{
					Group:   v1alpha1.SchemeGroupVersion.Group,
					Version: v1alpha1.SchemeGroupVersion.Version,
					Kind:    "Hive",
				}),
			},
			Labels: labels,
		},
		Spec: corev1.ServiceSpec{
			Type: "NodePort",
			Selector: map[string]string{
				"app": "hive-operator",
			},
			Ports: []corev1.ServicePort{
				{
					Protocol: corev1.Protocol("TCP"),
					TargetPort: intstr.IntOrString{
						StrVal: "8080"},
					Port: 8080,
				},
			},
		},
	}
}

// newbusyBoxPod demonstrates how to create a busybox pod
/*func newbusyBoxPod(cr *v1alpha1.Hive) *corev1.Pod {
	labels := map[string]string{
		"app": "busy-box",
	}
	apps.Deployment{}
	return &corev1.Pod{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Pod",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "busy-box",
			Namespace: cr.Namespace,
			OwnerReferences: []metav1.OwnerReference{
				*metav1.NewControllerRef(cr, schema.GroupVersionKind{
					Group:   v1alpha1.SchemeGroupVersion.Group,
					Version: v1alpha1.SchemeGroupVersion.Version,
					Kind:    "Hive",
				}),
			},
			Labels: labels,
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Name:    "busybox",
					Image:   "busybox",
					Command: []string{"sleep", "3600"},
				},
			},
		},
	}
}*/
