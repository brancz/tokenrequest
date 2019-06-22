package main

import (
	"flag"
	"fmt"

	authenticationv1 "k8s.io/api/authentication/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/klog"
)

func main() {
	klog.InitFlags(nil)
	flag.Parse()

	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	treq := &authenticationv1.TokenRequest{
		Spec: authenticationv1.TokenRequestSpec{
			Audiences: []string{"api"},
		},
	}

	_, err = clientset.CoreV1().ServiceAccounts("default").CreateToken("default", treq)
	if err != nil {
		panic(err)
	}
	fmt.Println("successfully created token")
}
