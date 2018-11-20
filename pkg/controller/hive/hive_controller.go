package hive

import (
	"context"
	hivev1alpha1 "github.com/openshift/hive-operator/pkg/apis/hive/v1alpha1"
	"io"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/yaml"
	"log"
	"os"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

/**
* USER ACTION REQUIRED: This is a scaffold file intended for the user to modify with their own Controller
* business logic.  Delete these comments after modifying this file.*
 */

// Add creates a new Hive Controller and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcileHive{client: mgr.GetClient(), scheme: mgr.GetScheme()}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("hive-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to primary resource Hive
	err = c.Watch(&source.Kind{Type: &hivev1alpha1.Hive{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	// TODO(user): Modify this to be the types you create that are owned by the primary resource
	// Watch for changes to secondary resource Pods and requeue the owner Hive
	err = c.Watch(&source.Kind{Type: &corev1.Pod{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &hivev1alpha1.Hive{},
	})
	if err != nil {
		return err
	}

	return nil
}

var _ reconcile.Reconciler = &ReconcileHive{}

// ReconcileHive reconciles a Hive object
type ReconcileHive struct {
	// This client, initialized using mgr.Client() above, is a split client
	// that reads objects from the cache and writes to the apiserver
	client client.Client
	scheme *runtime.Scheme
}

//Register components to the cluster
func registerComponents(componentPath string, r *ReconcileHive) {
	componentObject := unstructured.Unstructured{}
	file, err := os.Open(componentPath)
	if err != nil {
		panic(err.Error())
	}
	decoder := yaml.NewYAMLOrJSONDecoder(file, 65536)
	for {
		log.Print("Inside creation of resourse")
		err = decoder.Decode(&componentObject)
		if err == io.EOF {
			break
		}
		if err != nil && err != io.EOF {
			panic(err.Error())
		}
		r.client.Create(context.TODO(), &componentObject)

		if err != nil && !errors.IsAlreadyExists(err) {
			log.Print("Failed to create resource: %v", err)
		}
	}
}

// Reconcile reads that state of the cluster for a Hive object and makes changes based on the state read
// and what is in the Hive.Spec
// TODO(user): Modify this Reconcile function to implement your Controller logic.  This example creates
// a Pod as an example
// Note:
// The Controller will requeue the Request to be processed again if the returned error is non-nil or
// Result.Requeue is true, otherwise upon completion it will remove the work from the queue.
func (r *ReconcileHive) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	log.Printf("Reconciling Hive %s/%s\n", request.Namespace, request.Name)

	// Fetch the Hive instance
	instance := &hivev1alpha1.Hive{}
	err := r.client.Get(context.TODO(), request.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}

	// Define a new Pod object
	/*pod := newPodForCR(instance)

	// Set Hive instance as the owner and controller
	if err := controllerutil.SetControllerReference(instance, pod, r.scheme); err != nil {
		return reconcile.Result{}, err
	}

	// Check if this Pod already exists
	found := &corev1.Pod{}
	err = r.client.Get(context.TODO(), types.NamespacedName{Name: pod.Name, Namespace: pod.Namespace}, found)
	if err != nil && errors.IsNotFound(err) {
		log.Printf("Creating a new Pod %s/%s\n", pod.Namespace, pod.Name)
		err = r.client.Create(context.TODO(), pod)
		if err != nil {
			return reconcile.Result{}, err
		}

		// Pod created successfully - don't requeue
		return reconcile.Result{}, nil
	} else if err != nil {
		return reconcile.Result{}, err
	}

	// Pod already exists - don't requeue
	log.Printf("Skip reconcile: Pod %s/%s already exists", found.Namespace, found.Name)*/

	//Registering cluster-deployment CRD
	//u := v1alpha1.CustomResourceDefinitions{}
	/*u := unstructured.Unstructured{}
	f, err := os.Open("deploy/config/manager.yaml")
	if err != nil {
		panic(err.Error())
	}
	decoder := yaml.NewYAMLOrJSONDecoder(f, 65536)
	for {
		log.Print("Inside creation of cluster-deployment")
		//u = v1alpha1.CustomResourceDefinitions{}
		err = decoder.Decode(&u)
		if err == io.EOF {
			break
		}
		if err != nil && err != io.EOF {
			panic(err.Error())
		}
		//err = sdk.Create(&u)
		r.client.Create(context.TODO(), &u)

		if err != nil && !errors.IsAlreadyExists(err) {
			log.Print("Failed to create deployment.yaml: %v", err)
		}
	}
	log.Print("Outside creation of deployment")*/
	//return &u
	managerComponent := "deploy/config/manager.yaml"
	roleComponent := "deploy/config/rbac-role.yaml"
	registerComponents(managerComponent, r)
	registerComponents(roleComponent, r)

	return reconcile.Result{}, nil
}

// newPodForCR returns a busybox pod with the same name/namespace as the cr
func newPodForCR(cr *hivev1alpha1.Hive) *corev1.Pod {
	labels := map[string]string{
		"app": cr.Name,
	}
	return &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      cr.Name + "-pod",
			Namespace: cr.Namespace,
			Labels:    labels,
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
}
