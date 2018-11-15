package main

import (
	"flag"
	"log"
	"runtime"

	"github.com/openshift/hive-operator/pkg/apis"
	"github.com/openshift/hive-operator/pkg/controller"
	"github.com/operator-framework/operator-sdk/pkg/k8sutil"
	sdkVersion "github.com/operator-framework/operator-sdk/version"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/runtime/signals"
)

func printVersion() {
	log.Printf("Go Version: %s", runtime.Version())
	log.Printf("Go OS/Arch: %s/%s", runtime.GOOS, runtime.GOARCH)
	log.Printf("operator-sdk Version: %v", sdkVersion.Version)
}

func main() {
	printVersion()
	flag.Parse()

	namespace, err := k8sutil.GetWatchNamespace()
	if err != nil {
		log.Fatalf("failed to get watch namespace: %v", err)
	}

	// Get a config to talk to the apiserver
	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatal(err)
	}

	// Create a new Cmd to provide shared dependencies and start components
	mgr, err := manager.New(cfg, manager.Options{Namespace: namespace})
	if err != nil {
		log.Fatal(err)
	}

	log.Print("Registering Components.")

	// Setup Scheme for all resources
	if err := apis.AddToScheme(mgr.GetScheme()); err != nil {
		log.Fatal(err)
	}

	//Changes to register cluster-deployment CRD
	/*clusterDeploymentBytes, err := assets.Asset("deploy/crds/hive_v1alpha1_hive_cr.yaml")
	if err != nil {
		log.Fatal("Failed to get watch asset: %v", err)
	}
	clusterDeploymentCRD := resourceread.ReadCustomResourceDefinitionV1Beta1OrDie(clusterDeploymentBytes)

	resourceClient, _, err := k8sclient.GetResourceClient("v1alpha1", "hive", namespace)
	if err != nil {
		log.Fatal(err)
	}

	//c := config.GetConfig()

	_, updated, err := resourceapply.ApplyCustomResourceDefinition(
		resourceClient , clusterDeploymentCRD)
	if err != nil {
		log.Fatal("Failed to get register cluster deployment to kubernetes: %v", err)
	}

	log.Printf("Created cluster-deployment: %s", updated)*/


	// Setup all Controllers
	if err := controller.AddToManager(mgr); err != nil {
		log.Fatal(err)
	}

	log.Print("Starting the Cmd.")

	// Start the Cmd
	log.Fatal(mgr.Start(signals.SetupSignalHandler()))
	//creating cluster-deployment.yaml
	/*f, err := os.Open("deploy/cluster-deployment.yaml")
	if err != nil {
		panic(err.Error())
	}
	decoder := yaml.NewYAMLOrJSONDecoder(f, 65536)
	for {
		u := v1beta1.CustomResourceDefinition{}
		err = decoder.Decode(&u)
		if err == io.EOF {
			break
		}
		if err != nil && err != io.EOF {
			panic(err.Error())
		}
		//u.SetNamespace(namespace)
		err = sdk.Create(&u)
		if err != nil && !errors.IsAlreadyExists(err) {
			log.Fatal(err)
		}
	}*/
}
