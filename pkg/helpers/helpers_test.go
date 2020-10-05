package helpers

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"os"
	"testing"
	"time"
)

func TestGetClientSet(t *testing.T) {
	t.Run("Test Unset Environment Returns Default", func(t *testing.T) {
		_, err := GetClientSet()
		if err == nil {
			t.Errorf("expected an error but got none")
		}
	})
}

func TestGetEnv(t *testing.T) {
	os.Setenv("SET_ENV", "1")
	t.Run("Test Unset Environment Returns Default", func(t *testing.T) {
		environment := GetEnv("UNSET_ENV", "default")
		if environment != "default" {
			t.Errorf("environment = %v; want default", environment)
		}
	})
	t.Run("Test Set Environment Returns Default", func(t *testing.T) {
		environment := GetEnv("SET_ENV", "2")
		if environment != "1" {
			t.Errorf("environment = %v; want 1", environment)
		}
	})
}

func TestCheckDeleteResourceAllowed(t *testing.T) {
	os.Setenv("DISALLOW_LIST", "secrets,Statefulsets")
	t.Run("Test Delete Resource Deployments", func(t *testing.T) {
		check := CheckDeleteResourceAllowed("deployments")
		if !check {
			t.Errorf("check = %v", check)
		}
	})
	t.Run("Test Delete Resource Secrets Fails", func(t *testing.T) {
		check := CheckDeleteResourceAllowed("secrets")
		if check {
			t.Errorf("check = %v", check)
		}
	})
}

func TestPassedTimeToLive(t *testing.T) {
	os.Setenv("WORKLOAD_TTL", "1")
	t.Run("Should return false for time thats not over ttl", func(t *testing.T) {
		creationTime := time.Now()
		passedCheck := passedTimeToLive(metav1.NewTime(creationTime))
		if passedCheck {
			t.Errorf("passedCheck = %v; want false", passedCheck)
		}
	})
	t.Run("Should return true for time thats over ttl", func(t *testing.T) {
		creationTime := time.Now().Add(time.Minute * time.Duration(2))
		passedCheck := passedTimeToLive(metav1.NewTime(creationTime))
		if passedCheck {
			t.Errorf("passedCheck = %v; want false", passedCheck)
		}
	})
}

func TestEphemeralChecks(t *testing.T) {
	os.Setenv("WORKLOAD_TTL", "1")
	t.Run("Should fail Ephemeral Checks", func(t *testing.T) {
		creationTime := time.Now().Add(time.Minute * time.Duration(2))
		shouldDelete := EphemeralChecks{
			Name:         "pod-1234",
			CreationTime: metav1.NewTime(creationTime),
			Delete:       false,
		}
		shouldDelete.RunChecks()
		if shouldDelete.Delete {
			t.Errorf("delete = %v; want true", shouldDelete.Delete)
		}
	})
	t.Run("Should pass Ephemeral Checks", func(t *testing.T) {
		creationTime := time.Now().Add(time.Minute * time.Duration(2))
		shouldNotDelete := EphemeralChecks{
			Name:         "ephemeral-1234",
			CreationTime: metav1.NewTime(creationTime),
			Delete:       false,
		}
		shouldNotDelete.RunChecks()
		if shouldNotDelete.Delete {
			t.Errorf("delete = %v; want false", shouldNotDelete.Delete)
		}
	})

}
