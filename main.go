package main

import (
	"flag"
	"fmt"

	grpcserver "practice/k8sjobs/gRPCServer"
	httpserver "practice/k8sjobs/httpServer"
	"practice/k8sjobs/k8sStatus"
)

func main() {
	serverType := flag.String("server", "G", "Specifies the type of the server")

	clientset, err := k8sStatus.ConnectToK8s()
	if err != nil {
		panic(err.Error())
	}

	var k8sStatusMap = make(map[string][]string)

	k8sStatusMap["pods"], err = k8sStatus.GetPodDetails(clientset)
	if err != nil {
		panic(err.Error())
	}
	k8sStatusMap["services"], err = k8sStatus.GetServiceDetails(clientset)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("Number of pods: %v \n", len(k8sStatusMap["pods"]))
	fmt.Printf("Number of services: %v \n", len(k8sStatusMap["services"]))
	flag.Parse()
	fmt.Println(*serverType)
	if *serverType == "G" {
		grpcserver.StartGRPCServer(":50051", k8sStatusMap)
	} else if *serverType == "R" {
		hro := httpserver.HTTPRequestObj{}
		if err := httpserver.StartHTTPServer(&hro, k8sStatusMap); err != nil {
			panic(err)
		}
	} else {
		fmt.Println("Unsupported server type")
	}
	fmt.Println("Program ended")
}
