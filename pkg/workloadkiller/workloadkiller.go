package workloadkiller

import (
	"context"
	"github.com/spazzy757/ephemeral-enforcer/pkg/helpers"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"log"
	"os"
)

/*
KillWorkloads runs delete on all of the following in the namespace:
- deployments
- statefulsets
- services
- secrets
- configmaps
*/
func KillWorkloads(kubeconfig *rest.Config) {

	// creates the clientset
	clientset, err := kubernetes.NewForConfig(kubeconfig)
	if err != nil {
		panic(err.Error())
	}
	namespace := os.Getenv("NAMESPACE")
	// get pods in all the namespaces by omitting namespace
	// Or specify namespace to get pods in particular namespace
	deleteDeployments(clientset, &namespace)
	deleteStatefulsets(clientset, &namespace)
	deleteServices(clientset, &namespace)
	deleteSecrets(clientset, &namespace)
	deleteConfigMaps(clientset, &namespace)
}

func getDeleteList(resourceList []helpers.EphemeralChecks) []string {
	var deleteList []string
	for _, element := range resourceList {
		element.RunChecks()
		if element.Delete {
			deleteList = append(deleteList, element.Name)
		}
	}
	return deleteList
}

func deleteDeployments(clientset *kubernetes.Clientset, namespace *string) {
	client := clientset.AppsV1().Deployments(*namespace)
	deployments, err := client.List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatal("Error:", err.Error())
	}
	checks := []helpers.EphemeralChecks{}
	for _, element := range deployments.Items {
		checks = append(checks, helpers.EphemeralChecks{
			Name:         element.Name,
			CreationTime: element.CreationTimestamp,
			Delete:       false,
		})
	}
	deleteList := getDeleteList(checks)
	log.Printf("There are %d deployments scheduled for deletion\n", len(deleteList))
	deletePolicy := metav1.DeletePropagationForeground
	for _, element := range deleteList {
		if err := client.Delete(context.TODO(), element, metav1.DeleteOptions{
			PropagationPolicy: &deletePolicy,
		}); err != nil {
			log.Fatal(err)
		}
	}

}

func deleteStatefulsets(clientset *kubernetes.Clientset, namespace *string) {
	client := clientset.AppsV1().StatefulSets(*namespace)
	statefulsets, err := client.List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatal("Error:", err.Error())
	}
	checks := []helpers.EphemeralChecks{}
	for _, element := range statefulsets.Items {
		checks = append(checks, helpers.EphemeralChecks{
			Name:         element.Name,
			CreationTime: element.CreationTimestamp,
			Delete:       false,
		})
	}
	deleteList := getDeleteList(checks)
	log.Printf("There are %d statefulsets scheduled for deletion\n", len(deleteList))
	deletePolicy := metav1.DeletePropagationForeground
	for _, element := range deleteList {
		if err := client.Delete(context.TODO(), element, metav1.DeleteOptions{
			PropagationPolicy: &deletePolicy,
		}); err != nil {
			log.Fatal(err)
		}
	}
}

func deleteServices(clientset *kubernetes.Clientset, namespace *string) {
	client := clientset.CoreV1().Services(*namespace)
	services, err := client.List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatal("Error:", err.Error())
	}
	checks := []helpers.EphemeralChecks{}
	for _, element := range services.Items {
		checks = append(checks, helpers.EphemeralChecks{
			Name:         element.Name,
			CreationTime: element.CreationTimestamp,
			Delete:       false,
		})
	}
	deleteList := getDeleteList(checks)
	log.Printf("There are %d services scheduled for deletion\n", len(deleteList))
	deletePolicy := metav1.DeletePropagationForeground
	for _, element := range deleteList {
		if err := client.Delete(context.TODO(), element, metav1.DeleteOptions{
			PropagationPolicy: &deletePolicy,
		}); err != nil {
			log.Fatal(err)
		}
	}
}

func deleteSecrets(clientset *kubernetes.Clientset, namespace *string) {
	client := clientset.CoreV1().Secrets(*namespace)
	secrets, err := client.List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatal("Error:", err.Error())
	}
	checks := []helpers.EphemeralChecks{}
	for _, element := range secrets.Items {
		checks = append(checks, helpers.EphemeralChecks{
			Name:         element.Name,
			CreationTime: element.CreationTimestamp,
			Delete:       false,
		})
	}
	deleteList := getDeleteList(checks)
	log.Printf("There are %d secrets scheduled for deletion\n", len(deleteList))
	deletePolicy := metav1.DeletePropagationForeground
	for _, element := range deleteList {
		if err := client.Delete(context.TODO(), element, metav1.DeleteOptions{
			PropagationPolicy: &deletePolicy,
		}); err != nil {
			log.Fatal(err)
		}
	}
}

func deleteConfigMaps(clientset *kubernetes.Clientset, namespace *string) {
	client := clientset.CoreV1().ConfigMaps(*namespace)
	configmaps, err := client.List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatal("Error:", err.Error())
	}
	checks := []helpers.EphemeralChecks{}
	for _, element := range configmaps.Items {
		checks = append(checks, helpers.EphemeralChecks{
			Name:         element.Name,
			CreationTime: element.CreationTimestamp,
			Delete:       false,
		})
	}
	deleteList := getDeleteList(checks)
	log.Printf("There are %d configmaps scheduled for deletion\n", len(deleteList))
	deletePolicy := metav1.DeletePropagationForeground
	for _, element := range deleteList {
		if err := client.Delete(context.TODO(), element, metav1.DeleteOptions{
			PropagationPolicy: &deletePolicy,
		}); err != nil {
			log.Fatal(err)
		}
	}
}
