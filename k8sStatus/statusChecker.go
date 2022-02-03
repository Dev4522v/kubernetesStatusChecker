package k8sStatus

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func ConnectToK8s() (*kubernetes.Clientset, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, err
	}

	// var kubeconfig *string
	// if home := homedir.HomeDir(); home != "" {
	// 	kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	// } else {
	// 	kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	// }
	// flag.Parse()

	// // use the current context in kubeconfig
	// config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	// if err != nil {
	// 	panic(err.Error())
	// }

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Println("Failed to build config")
		return nil, err
	}

	return clientset, nil
}

func GetServiceDetails(clientset kubernetes.Interface) ([]string, error) {
	services, err := clientset.CoreV1().Services("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		fmt.Println("Error getting the number of services")
		return []string{}, err
	}
	serviceNames := []string{}
	for _, service := range services.Items {
		serviceNames = append(serviceNames, service.Name)
	}
	return serviceNames, nil
}

func GetPodDetails(clientset kubernetes.Interface) ([]string, error) {
	pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		fmt.Println("Error getting the number of pods")
		return []string{}, err
	}
	podNames := []string{}
	for _, pod := range pods.Items {
		podNames = append(podNames, pod.Name)
	}
	return podNames, nil
}
