package workloadkiller

import (
	"context"
	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	fake "k8s.io/client-go/kubernetes/fake"
	"log"
	"os"
	"sync"
	"testing"
)

func TestKillWorkloads(t *testing.T) {
	os.Setenv("EPHEMERAL_ENFORCER_NAME", "ephemeral")
	fakeClientSet := fake.NewSimpleClientset(
		&appsv1.Deployment{
			ObjectMeta: metav1.ObjectMeta{
				Name:        "ephemeral",
				Namespace:   "default",
				Annotations: map[string]string{},
			},
		},
		&appsv1.Deployment{
			ObjectMeta: metav1.ObjectMeta{
				Name:        "other",
				Namespace:   "default",
				Annotations: map[string]string{},
			},
		},
		&appsv1.StatefulSet{
			ObjectMeta: metav1.ObjectMeta{
				Name:        "other",
				Namespace:   "default",
				Annotations: map[string]string{},
			},
		},
	)
	t.Run("Test Delete Deployments", func(t *testing.T) {
		namespace := "default"
		KillWorkloads(fakeClientSet)
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

func TestDeleteDeployments(t *testing.T) {
	os.Setenv("EPHEMERAL_ENFORCER_NAME", "ephemeral")
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
		var wg sync.WaitGroup
		wg.Add(1)
		deleteDeployments(fakeClientSet, &namespace, &wg)
		wg.Wait()
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

func TestDeleteStatefulSets(t *testing.T) {
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
	t.Run("Test Delete StatefulSets", func(t *testing.T) {
		namespace := "default"
		os.Setenv("EPHEMERAL_ENFORCER_NAME", "ephemeral")
		var wg sync.WaitGroup
		wg.Add(1)
		deleteStatefulsets(fakeClientSet, &namespace, &wg)
		wg.Wait()
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

func TestDeleteServices(t *testing.T) {
	fakeClientSet := fake.NewSimpleClientset(&v1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:        "ephemeral",
			Namespace:   "default",
			Annotations: map[string]string{},
		},
	}, &v1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:        "other",
			Namespace:   "default",
			Annotations: map[string]string{},
		},
	})
	t.Run("Test Delete Services", func(t *testing.T) {
		namespace := "default"
		os.Setenv("EPHEMERAL_ENFORCER_NAME", "ephemeral")
		var wg sync.WaitGroup
		wg.Add(1)
		deleteServices(fakeClientSet, &namespace, &wg)
		wg.Wait()
		services, err := fakeClientSet.CoreV1().Services(namespace).List(
			context.TODO(),
			metav1.ListOptions{},
		)
		if err != nil {
			log.Fatal("Error:", err.Error())
		}
		if len(services.Items) != 1 {
			t.Errorf("Expected No Services but got %v", len(services.Items))
		}
	})
}

func TestDeleteSecrets(t *testing.T) {
	fakeClientSet := fake.NewSimpleClientset(&v1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:        "ephemeral",
			Namespace:   "default",
			Annotations: map[string]string{},
		},
	}, &v1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:        "other",
			Namespace:   "default",
			Annotations: map[string]string{},
		},
	})
	t.Run("Test Delete Services", func(t *testing.T) {
		namespace := "default"
		os.Setenv("EPHEMERAL_ENFORCER_NAME", "ephemeral")
		var wg sync.WaitGroup
		wg.Add(1)
		deleteSecrets(fakeClientSet, &namespace, &wg)
		wg.Wait()
		secrets, err := fakeClientSet.CoreV1().Secrets(namespace).List(
			context.TODO(),
			metav1.ListOptions{},
		)
		if err != nil {
			log.Fatal("Error:", err.Error())
		}
		if len(secrets.Items) != 1 {
			t.Errorf("Expected No Services but got %v", len(secrets.Items))
		}
	})
}

func TestDeleteConfigMaps(t *testing.T) {
	fakeClientSet := fake.NewSimpleClientset(&v1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:        "ephemeral",
			Namespace:   "default",
			Annotations: map[string]string{},
		},
	}, &v1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:        "other",
			Namespace:   "default",
			Annotations: map[string]string{},
		},
	})
	t.Run("Test Delete ConfigMaps", func(t *testing.T) {
		namespace := "default"
		os.Setenv("EPHEMERAL_ENFORCER_NAME", "ephemeral")
		var wg sync.WaitGroup
		wg.Add(1)
		deleteConfigMaps(fakeClientSet, &namespace, &wg)
		wg.Wait()
		configmaps, err := fakeClientSet.CoreV1().ConfigMaps(namespace).List(
			context.TODO(),
			metav1.ListOptions{},
		)
		if err != nil {
			log.Fatal("Error:", err.Error())
		}
		if len(configmaps.Items) != 1 {
			t.Errorf("Expected 1 ConfigMap but got %v", len(configmaps.Items))
		}
	})
}
