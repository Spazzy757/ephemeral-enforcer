package helpers

import (
	//"flag"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	//"k8s.io/client-go/tools/clientcmd"
	"os"
	//	"path/filepath"
	"strconv"
	"strings"
	"time"
)

/*
GetClientSet Generates a clientset from the Kubeconfig
*/
func GetClientSet() (kubernetes.Interface, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, err
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	return clientset, nil
}

/*
GetEnv looks up an env key or returns a default
*/
func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

/*
EphemeralChecks Struct that holds the information of the resources that could be deleted
*/
type EphemeralChecks struct {
	Name         string
	CreationTime metav1.Time
	Delete       bool
}

/*
RunChecks Run checks to see if resource should be deleted
*/
func (e *EphemeralChecks) RunChecks() {
	// Run The check to see if the age is past the TTL
	if passedTimeToLive(e.CreationTime) {
		e.Delete = true
	}
}

func passedTimeToLive(creationTime metav1.Time) bool {
	ttl, _ := strconv.Atoi(os.Getenv("WORKLOAD_TTL"))
	// Convert to a negative
	ttl = 0 - ttl
	//
	previousTime := time.Now().Add(time.Minute * time.Duration(ttl))
	ttlTime := metav1.NewTime(previousTime)
	return creationTime.Before(&ttlTime)
}

/*
CheckDeleteResourceAllowed Checks if the resource is in the disallow list
*/
func CheckDeleteResourceAllowed(resourceType string) bool {
	disallowList := GetEnv("DISALLOW_LIST", "")
	for _, element := range strings.Split(disallowList, ",") {
		if strings.Contains(strings.ToLower(resourceType), strings.ToLower(element)) && element != "" {
			return false
		}
	}
	return true
}
