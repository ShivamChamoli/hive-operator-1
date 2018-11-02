package main

import (
	"context"
	"runtime"

	"github.com/openshift/machine-config-operator/lib/resourceapply"
	"github.com/openshift/machine-config-operator/pkg/operator/assets"

	"github.com/official-hive-operator/hive-operator-1/lib/resourceread"
	"github.com/official-hive-operator/hive-operator-1/pkg/stub"
	"github.com/operator-framework/operator-sdk/pkg/sdk"
	"github.com/operator-framework/operator-sdk/pkg/util/k8sutil"
	sdkVersion "github.com/operator-framework/operator-sdk/version"
	apiextclientset "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"

	"time"

	"github.com/sirupsen/logrus"
)

func printVersion() {
	logrus.Infof("Go Version: %s", runtime.Version())
	logrus.Infof("Go OS/Arch: %s/%s", runtime.GOOS, runtime.GOARCH)
	logrus.Infof("operator-sdk Version: %v", sdkVersion.Version)
}

func main() {
	printVersion()
	
	resource := "hive.openshift.io/v1alpha1"
	kind := "Hive"
	namespace, err := k8sutil.GetWatchNamespace()
	if err != nil {
		logrus.Fatalf("Failed to get watch namespace: %v", err)
	}

	//code for getting data from bindata and registering it to kubernetes
	clusterDeploymentBytes, err := assets.Asset("deploy/cluster-deployment.yaml")
	if err != nil {
		logrus.Fatalf("Failed to get watch asset: %v", err)
	}
	clusterDeploymentCRD := resourceread.ReadCustomResourceDefinitionV1Beta1OrDie(clusterDeploymentBytes)

	_, updated, err := resourceapply.ApplyCustomResourceDefinition(apiextclientset.Interface.ApiextensionsV1beta1(clusterDeploymentCRD), clusterDeploymentCRD)
	if err != nil {
		logrus.Fatalf("Failed to get register cluster deployment to kubernetes: %v", err)
	}

	logrus.Infof("Created cluster-deployment: %s", updated)

	/*if updated {
		if err := client.waitForCustomResourceDefinition(clusterDeploymentCRD); err != nil {
			logrus.Fatalf("Failed to create cluster deployment CRD: %v", err)
		}
	}*/

	
	//code for parsing a CRD and registering it to the kubernetes api
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
			logrus.Errorf("Failed to create crd.yaml: %v", err)
		}
	}*/

	resyncPeriod, _ := time.ParseDuration("5s")
	logrus.Infof("Watching %s, %s, %s, %d", resource, kind, namespace, resyncPeriod)
	sdk.Watch(resource, kind, namespace, time.Duration(5)*time.Second)
	sdk.Handle(stub.NewHandler())
	sdk.Run(context.TODO())
}
