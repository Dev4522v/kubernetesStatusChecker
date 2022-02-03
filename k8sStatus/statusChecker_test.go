package k8sStatus_test

import (
	"context"
	"fmt"
	"testing"

	k8sStatus "practice/k8sjobs/k8sStatus"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
)

func TestGetPodDetails(t *testing.T) {
	clientSet := fake.NewSimpleClientset()
	pod := &v1.Pod{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Pod",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-pod",
			Namespace: "default",
		},
		Spec: v1.PodSpec{
			Containers: []v1.Container{
				{
					Name:            "nginx",
					Image:           "nginx",
					ImagePullPolicy: "Always",
				},
			},
		},
	}
	_, err := clientSet.CoreV1().Pods(pod.Namespace).Create(context.TODO(), pod, metav1.CreateOptions{})
	if err != nil {
		t.Fatalf("GetNumberOfPods testcase failed with error: %v", err)
	}
	val, err1 := k8sStatus.GetPodDetails(clientSet)
	if err1 != nil {
		t.Fatalf("GetNumberOfPods testcase failed with error: %v", err1)
	}
	fmt.Println(val)
	fmt.Println("Working fine")
}
