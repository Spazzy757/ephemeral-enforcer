package workloadkiller

import (
	"context"
	"github.com/spazzy757/ephemeral-enforcer/pkg/helpers"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"log"
	"sync"
)

/*
KillWorkloads runs delete on all of the following in the namespace:
- deployments
- statefulsets
- services
- secrets
- configmaps
*/
func KillWorkloads(clientset kubernetes.Interface) {
	// Look for namespace or default to default namespace
	namespace := helpers.GetEnv("NAMESPACE", "default")
	// Wait Group To handle the waiting for all deletes to complete
	var wg sync.WaitGroup
	wg.Add(6)
	// Delete all Deployments
	if helpers.CheckDeleteResourceAllowed("deployments") {
		go deleteDeployments(clientset, &namespace, &wg)
	}
	// Delete all Statefulsets
	if helpers.CheckDeleteResourceAllowed("statefulsets") {
		go deleteStatefulsets(clientset, &namespace, &wg)
	}
	// Delete Services
	if helpers.CheckDeleteResourceAllowed("services") {
		go deleteServices(clientset, &namespace, &wg)
	}
	// Delete All Secrets
	if helpers.CheckDeleteResourceAllowed("secrets") {
		go deleteSecrets(clientset, &namespace, &wg)
	}
	// Delete All Configmaps
	if helpers.CheckDeleteResourceAllowed("configmaps") {
		go deleteConfigMaps(clientset, &namespace, &wg)
	}
	// Delete All Daemonsets
	if helpers.CheckDeleteResourceAllowed("daemonsets") {
		go deleteDaemonSets(clientset, &namespace, &wg)
	}
	// wait for processes to finish
	wg.Wait()
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

func deleteDeployments(clientset kubernetes.Interface, namespace *string, wg *sync.WaitGroup) {
	defer wg.Done()
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

func deleteStatefulsets(clientset kubernetes.Interface, namespace *string, wg *sync.WaitGroup) {
	defer wg.Done()
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

func deleteServices(clientset kubernetes.Interface, namespace *string, wg *sync.WaitGroup) {
	defer wg.Done()
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

func deleteSecrets(clientset kubernetes.Interface, namespace *string, wg *sync.WaitGroup) {
	defer wg.Done()
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

func deleteConfigMaps(clientset kubernetes.Interface, namespace *string, wg *sync.WaitGroup) {
	defer wg.Done()
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

func deleteDaemonSets(clientset kubernetes.Interface, namespace *string, wg *sync.WaitGroup) {
	defer wg.Done()
	client := clientset.AppsV1().DaemonSets(*namespace)
	daemonsets, err := client.List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatal("Error:", err.Error())
	}
	checks := []helpers.EphemeralChecks{}
	for _, element := range daemonsets.Items {
		checks = append(checks, helpers.EphemeralChecks{
			Name:         element.Name,
			CreationTime: element.CreationTimestamp,
			Delete:       false,
		})
	}
	deleteList := getDeleteList(checks)
	log.Printf("There are %d daemonsets scheduled for deletion\n", len(deleteList))
	deletePolicy := metav1.DeletePropagationForeground
	for _, element := range deleteList {
		if err := client.Delete(context.TODO(), element, metav1.DeleteOptions{
			PropagationPolicy: &deletePolicy,
		}); err != nil {
			log.Fatal(err)
		}
	}
}