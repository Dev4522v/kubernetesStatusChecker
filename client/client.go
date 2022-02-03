package main

import (
	"context"
	"fmt"
	"log"
	"practice/k8sjobs/statusProto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	fmt.Println("Hello client .....")

	opts := grpc.WithTransportCredentials(insecure.NewCredentials())
	cc, err := grpc.Dial("192.168.64.2:30001", opts)
	if err != nil {
		log.Fatal(err)
	}
	defer cc.Close()

	client := statusProto.NewStatusServiceClient(cc)

	var needItem string

	for needItem != "Q" {
		fmt.Println("Which item do you want?")
		fmt.Println("Press P for Pods")
		fmt.Println("Press S for Services")
		fmt.Println("Press Q to QUIT")
		fmt.Scanln(&needItem)
		var request statusProto.StatusRequest
		if needItem == "S" {
			request = statusProto.StatusRequest{
				ItemType: "services",
			}
			fmt.Println("")
			fmt.Println("Services running are: ")
			fmt.Println("")
		}
		if needItem == "P" {
			request = statusProto.StatusRequest{
				ItemType: "pods",
			}
			fmt.Println("")
			fmt.Println("Pods running are: ")
			fmt.Println("")
		}
		if needItem == "Q" {
			break
		}
		resp, _ := client.Status(context.Background(), &request)
		fmt.Println(resp)

		for _, item := range resp.ListItems {
			for _, itemName := range item.Items {
				fmt.Println(itemName.Name)
			}
		}
		fmt.Println("")
	}

}
