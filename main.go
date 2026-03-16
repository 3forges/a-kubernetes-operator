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
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/dynamic/dynamicinformer"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	log.Printf("PESTO-OPERATOR - Now creating Kubernetes Client")
	client, err := getClient()
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("PESTO-OPERATOR - Now setting up informer with clientset: %s", client)
	// var pestoInformer v1.PodInformer
	var pestoInformer = setupInformer(client)

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
	pestoInformer.Run(ctx.Done())
}

// (*dynamic.DynamicClient, error)
func setupInformer(dynClient *dynamic.DynamicClient) cache.SharedIndexInformer {
	// informerFactory := informers.NewSharedInformerFactory(clientset, time.Minute)
	// podInformer := informerFactory.Core().V1().Pods()

	// here very important: the resource MUST be plural form or the informer reflector will fail
	resource := schema.GroupVersionResource{Group: "stable.pesto.io", Version: "v1", Resource: "baovaults"}

	factory := dynamicinformer.NewFilteredDynamicSharedInformerFactory(dynClient, time.Minute, v1.NamespaceAll, nil)
	pestoInformer := factory.ForResource(resource).Informer()

	pestoInformer.AddEventHandler(
		cache.ResourceEventHandlerFuncs{
			AddFunc: func(obj interface{}) {

				// var catchedResource = obj.(*v1.Pod)
				var catchedResource = obj.(*unstructured.Unstructured)
				log.Printf("PESTO-OPERATOR - A new deployment was created and detected by the operator Name : %s", catchedResource.GetName())
				log.Printf("PESTO-OPERATOR - A new deployment was created and detected by the operator Labels : %s", fmt.Sprintf("%v", catchedResource.GetLabels()))
				log.Printf("PESTO-OPERATOR - A new deployment was created and detected by the operator Kind : %s", catchedResource.GetKind())
				log.Printf("PESTO-OPERATOR - A new deployment was created and detected by the operator APIVersion : %s", catchedResource.GetAPIVersion())
				log.Printf("PESTO-OPERATOR - A new deployment was created and detected by the operator CreationTimestamp : %s", catchedResource.GetCreationTimestamp())
				log.Printf("PESTO-OPERATOR - A new deployment was created and detected by the operator Namespace : %s", catchedResource.GetNamespace())
				log.Printf("PESTO-OPERATOR - A new deployment was created and detected by the operator catchedResource.GetManagedFields()[0].Subresource : %s", catchedResource.GetManagedFields()[0].Subresource)
				log.Printf("PESTO-OPERATOR - A new deployment was created and detected by the operator catchedResource.GetManagedFields()[0] : %s", catchedResource.GetManagedFields())
				log.Printf("PESTO-OPERATOR - A new deployment was created and detected by the operator catchedResource.GetManagedFields()[0] : %s", catchedResource.GetManagedFields()[0])
				log.Printf("PESTO-OPERATOR - A new deployment was created and detected by the operator catchedResource.GetManagedFields()[0].SwaggerDoc() : %s", catchedResource.GetManagedFields()[0].SwaggerDoc())
				log.Printf("PESTO-OPERATOR - A new deployment was created and detected by the operator catchedResource.GetManagedFields()[0].FieldsV1 : %s", catchedResource.GetManagedFields()[0].FieldsV1)
				log.Printf("PESTO-OPERATOR - A new deployment was created and detected by the operator catchedResource.GetManagedFields()[0].FieldsV1.Raw : %s", catchedResource.GetManagedFields()[0].FieldsV1.Raw)
				log.Printf("PESTO-OPERATOR - A new deployment was created and detected by the operator catchedResource.GetManagedFields()[0].FieldsType : %s", catchedResource.GetManagedFields()[0].FieldsType)
			},
			UpdateFunc: func(oldObj, newObj interface{}) {
				// var catchedResourceOldObj = oldObj.(*v1.Pod)
				// var catchedResourceNewObj = newObj.(*v1.Pod)
				var catchedResourceOldObj = oldObj.(*unstructured.Unstructured)
				var catchedResourceNewObj = newObj.(*unstructured.Unstructured)

				log.Printf("PESTO-OPERATOR - A deployment update is on the way and detected by the operator : %s", catchedResourceNewObj.GetName())
				log.Printf("PESTO-OPERATOR - A deployment update is on the way and detected by the operator : %s", catchedResourceNewObj.GetLabels())
				log.Printf("PESTO-OPERATOR - A new deployment update is on the way and detected by the operator catchedResourceNewObj.GetManagedFields()[0] : %s", catchedResourceNewObj.GetManagedFields())
				log.Printf("PESTO-OPERATOR - A new deployment update is on the way and detected by the operator catchedResourceNewObj.GetManagedFields()[0] : %s", catchedResourceNewObj.GetManagedFields()[0])
				log.Printf("PESTO-OPERATOR - A new deployment update is on the way and detected by the operator catchedResourceNewObj.GetGenerateName() : %s", catchedResourceNewObj.GetGenerateName())
				log.Printf("PESTO-OPERATOR - A new deployment update is on the way and detected by the operator catchedResourceNewObj.GetManagedFields()[0].SwaggerDoc() : %s", catchedResourceNewObj.GetManagedFields()[0].SwaggerDoc())
				log.Printf("PESTO-OPERATOR - A new deployment update is on the way and detected by the operator catchedResourceNewObj.GetManagedFields()[0].FieldsV1 : %s", catchedResourceNewObj.GetManagedFields()[0].FieldsV1)
				log.Printf("PESTO-OPERATOR - A new deployment update is on the way and detected by the operator catchedResourceNewObj.GetManagedFields()[0].FieldsV1.Raw : %s", catchedResourceNewObj.GetManagedFields()[0].FieldsV1.Raw)
				log.Printf("PESTO-OPERATOR - A new deployment update is on the way and detected by the operator catchedResourceNewObj.GetManagedFields()[0].FieldsType : %s", catchedResourceNewObj.GetManagedFields()[0].FieldsType)
				// log.Printf("PESTO-OPERATOR - The deployment that is being updated shall become Name: %s", catchedResourceOldObj.GetName())
				// log.Printf("PESTO-OPERATOR - The deployment that is being updated shall become Labels: %s", catchedResourceOldObj.GetLabels())
			},
			DeleteFunc: func(obj interface{}) {
				// var catchedResource = obj.(*v1.Pod)
				//panic: interface conversion: interface {} is *v1.Pod, not *client.V1Pod [recovered, repanicked]
				//

				// var catchedResource = obj.(*v1.Pod)
				var catchedResource = obj.(*unstructured.Unstructured)
				log.Printf("PESTO-OPERATOR - An existing deployment was deleted and that was detected by the operator Name : %s", catchedResource.GetName())
				// Handle pod deletion event
				log.Printf("PESTO-OPERATOR - An existing deployment was deleted and that was detected by the operator Namespace : %s", catchedResource.GetNamespace())
				log.Printf("PESTO-OPERATOR - An existing deployment was deleted and that was detected by the operator Labels : %s", catchedResource.GetLabels())
			},
		},
	)
	return pestoInformer
}

func getClient() (*dynamic.DynamicClient, error) {
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
	clusterClient, err := dynamic.NewForConfig(clusterConfig)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	/*
		clusterClient, err := dynamic.NewForConfig(clusterConfig)
		if err != nil {
			log.Fatalln(err)
		}

		clientset, err := kubernetes.NewForConfig(clusterConfig)
		if err != nil {
			log.Fatalln(fmt.Errorf("failed to initialise clientset from config: %s", err))
			return nil, err
		}
	*/

	return clusterClient, nil
}
