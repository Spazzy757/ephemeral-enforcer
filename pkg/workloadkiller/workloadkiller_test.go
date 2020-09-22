package workloadkiller

import (
	"context"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	fake "k8s.io/client-go/kubernetes/fake"
	"log"
	"os"
	"testing"
)

func TestDeleteDeployments(t *testing.T) {
	fakeClientSet := fake.NewSimpleClientset(&appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:        "ephemeral",
			Namespace:   "default",
			Annotations: map[string]string{},
		},
	}, &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:        "other",
			Namespace:   "default",
			Annotations: map[string]string{},
		},
	})
	t.Run("Test Delete Deployments", func(t *testing.T) {
		namespace := "default"
		os.Setenv("EPHEMERAL_ENFORCER_NAME", "ephemeral")
		deleteDeployments(fakeClientSet, &namespace)
		deployments, err := fakeClientSet.AppsV1().Deployments(namespace).List(
			context.TODO(),
			metav1.ListOptions{},
		)
		if err != nil {
			log.Fatal("Error:", err.Error())
		}
		if len(deployments.Items) != 1 {
			t.Errorf("Expected No Deployments but got %v", len(deployments.Items))
		}
	})
}

func TestDeleteStatefulsets(t *testing.T) {
	fakeClientSet := fake.NewSimpleClientset(&appsv1.StatefulSet{
		ObjectMeta: metav1.ObjectMeta{
			Name:        "ephemeral",
			Namespace:   "default",
			Annotations: map[string]string{},
		},
	}, &appsv1.StatefulSet{
		ObjectMeta: metav1.ObjectMeta{
			Name:        "other",
			Namespace:   "default",
			Annotations: map[string]string{},
		},
	})
	t.Run("Test Delete Statefulsetss", func(t *testing.T) {
		namespace := "default"
		os.Setenv("EPHEMERAL_ENFORCER_NAME", "ephemeral")
		deleteStatefulsets(fakeClientSet, &namespace)
		statefulsets, err := fakeClientSet.AppsV1().StatefulSets(namespace).List(
			context.TODO(),
			metav1.ListOptions{},
		)
		if err != nil {
			log.Fatal("Error:", err.Error())
		}
		if len(statefulsets.Items) != 1 {
			t.Errorf("Expected No StatefulSets but got %v", len(statefulsets.Items))
		}
	})
}
