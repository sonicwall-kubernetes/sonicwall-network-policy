package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog/v2"

	swnpv1alpha1 "github.com/sonicwall-kubernetes/sonicwall-network-policy/pkg/apis/sonicwallnetworkpolicy/v1alpha1"
	swnpclientset "github.com/sonicwall-kubernetes/sonicwall-network-policy/pkg/generated/clientset/versioned"
	swnpinformers "github.com/sonicwall-kubernetes/sonicwall-network-policy/pkg/generated/informers/externalversions"
)

func getKubeConfigFullPath() string {
	var kubeconfig string
	if envvar := os.Getenv("KUBECONFIG"); len(envvar) > 0 {
		kubeconfig = envvar
	} else {
		kubeconfig = filepath.Join(os.Getenv("HOME"), ".kube/config")
	}
	return kubeconfig
}

func handleAddSonicwallNetworkPolicy(new interface{}) {
	np := new.(*swnpv1alpha1.SonicwallNetworkPolicy)
	fmt.Println("add", np.Name)
}

func handleUpdateSonicwallNetworkPolicy(old, new interface{}) {
	np := new.(*swnpv1alpha1.SonicwallNetworkPolicy)
	fmt.Println("modify", np.Name)
}

func handleDeleteSonicwallNetworkPolicy(obj interface{}) {
	np := obj.(*swnpv1alpha1.SonicwallNetworkPolicy)
	fmt.Println("delete", np.Name)
}

func main() {
	config, err := rest.InClusterConfig()
	if err != nil {
		kubeconfig := getKubeConfigFullPath()
		config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
		if err != nil {
			panic(err.Error())
		}
	}

	swnpClient, err := swnpclientset.NewForConfig(config)
	if err != nil {
		klog.Fatalf("Error building swnp clientset: %s", err.Error())
	}

	swnpInformerFactory := swnpinformers.NewSharedInformerFactory(swnpClient, time.Minute*5)

	swnpInformer := swnpInformerFactory.K8s().V1alpha1().SonicwallNetworkPolicies()

	swnpInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    handleAddSonicwallNetworkPolicy,
		UpdateFunc: handleUpdateSonicwallNetworkPolicy,
		DeleteFunc: handleDeleteSonicwallNetworkPolicy,
	})

	swnpInformerFactory.Start(wait.NeverStop)
	swnpInformerFactory.WaitForCacheSync(wait.NeverStop)

	<-make(chan int)
}
