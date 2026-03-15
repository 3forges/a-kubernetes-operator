package main

// package operator

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/informers"
	v1Informers "k8s.io/client-go/informers/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	log.Printf("PESTO-OPERATOR - Now creating Kubernetes Client")
	clientset, err := getClient()
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("PESTO-OPERATOR - Now setting up informer with clientset: %s", clientset)
	// var pestoInformer v1.PodInformer
	var pestoInformer = setupInformer(clientset)

	/**
	 * We’re creating a context.Context that is cancelled on
	 * an os.Interrupt signal.
	 * This allows us to prevent the application from
	 * exiting until it receives an interrupt signal.
	 *
	 * Its Done channel is passed to podInformer.Informer.Run(), to
	 * keep the informer alive until execution is cancelled.
	 **/
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	/*
		pestoInformer.Run(ctx.Done())
	*/
	pestoInformer.Informer().Run(ctx.Done())
}

func setupInformer(clientset *kubernetes.Clientset) v1Informers.PodInformer {
	informerFactory := informers.NewSharedInformerFactory(clientset, time.Minute)
	podInformer := informerFactory.Core().V1().Pods()

	podInformer.Informer().AddEventHandler(
		cache.ResourceEventHandlerFuncs{
			AddFunc: func(obj interface{}) {
				var catchedPod = obj.(*v1.Pod)
				log.Printf("PESTO-OPERATOR - A new pod was created and detected by the operator Name : %s", catchedPod.Name)
				log.Printf("PESTO-OPERATOR - A new pod was created and detected by the operator Labels : %s", fmt.Sprintf("%v", catchedPod.Labels))
				log.Printf("PESTO-OPERATOR - A new pod was created and detected by the operator Kind : %s", catchedPod.Kind)
				log.Printf("PESTO-OPERATOR - A new pod was created and detected by the operator APIVersion : %s", catchedPod.APIVersion)
				log.Printf("PESTO-OPERATOR - A new pod was created and detected by the operator CreationTimestamp : %s", catchedPod.CreationTimestamp)
				log.Printf("PESTO-OPERATOR - A new pod was created and detected by the operator Status : %s", catchedPod.Status)
				log.Printf("PESTO-OPERATOR - A new pod was created and detected by the operator Spec.Containers[0].Image : %s", catchedPod.Spec.Containers[0].Image)

			},
			UpdateFunc: func(oldObj, newObj interface{}) {
				var catchedPodOldObj = oldObj.(*v1.Pod)
				var catchedPodNewObj = newObj.(*v1.Pod)
				log.Printf("PESTO-OPERATOR - A pod was update is on the way and detected by the operator : %s", catchedPodNewObj.Name)
				log.Printf("PESTO-OPERATOR - A pod was update is on the way and detected by the operator : %s", catchedPodNewObj.Labels)
				log.Printf("PESTO-OPERATOR - The pod that is being updated shall become Name: %s", catchedPodOldObj.Name)
				log.Printf("PESTO-OPERATOR - The pod that is being updated shall become Labels: %s", catchedPodOldObj.Labels)
			},
			DeleteFunc: func(obj interface{}) {
				// var catchedPod = obj.(*v1.Pod)
				//panic: interface conversion: interface {} is *v1.Pod, not *client.V1Pod [recovered, repanicked]
				//

				var catchedPod = obj.(*v1.Pod)
				log.Printf("PESTO-OPERATOR - An existing pod was deleted and that was detected by the operator Name : %s", catchedPod.Name)
				// Handle pod deletion event
				log.Printf("PESTO-OPERATOR - An existing pod was deleted and that was detected by the operator Namespace : %s", catchedPod.Namespace)
				log.Printf("PESTO-OPERATOR - An existing pod was deleted and that was detected by the operator Labels : %s", catchedPod.Labels)
			},
		},
	)
	return podInformer
}

func getClient() (*kubernetes.Clientset, error) {
	kubeConfig := os.Getenv("KUBECONFIG")

	var clusterConfig *rest.Config
	var err error
	if kubeConfig != "" {
		clusterConfig, err = clientcmd.BuildConfigFromFlags("", kubeConfig)
	} else {
		clusterConfig, err = rest.InClusterConfig()
	}
	if err != nil {
		log.Fatalln(err)
	}

	/*
		clusterClient, err := dynamic.NewForConfig(clusterConfig)
		if err != nil {
			log.Fatalln(err)
		}
	*/

	clientset, err := kubernetes.NewForConfig(clusterConfig)
	if err != nil {
		log.Fatalln(fmt.Errorf("failed to initialise clientset from config: %s", err))
		return nil, err
	}

	return clientset, nil
}
