package grpcserver

import (
	"context"
	"fmt"
	"log"
	"net"
	"practice/k8sjobs/statusProto"

	"google.golang.org/grpc"
)

type server struct {
	statusProto.UnimplementedStatusServiceServer
	mp map[string][]string
}

func StartGRPCServer(addr string, k8sStatusMap map[string][]string) error {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Error %v", err)
	}
	fmt.Printf("Server is listening on %v ...", addr)

	s := grpc.NewServer()
	statusProto.RegisterStatusServiceServer(s, &server{mp: k8sStatusMap})
	return s.Serve(lis)
}

func (s *server) Status(ctx context.Context, request *statusProto.StatusRequest) (*statusProto.StatusList, error) {
	podList := statusProto.ItemInstances{}
	requestType := request.ItemType
	for _, str := range s.mp[requestType] {
		v := statusProto.Item{
			Name: str,
		}
		podList.Items = append(podList.Items, &v)
	}
	ans := statusProto.StatusList{}
	ans.ListItems = append(ans.ListItems, &podList)
	response := &statusProto.StatusList{
		ListItems: ans.ListItems,
	}
	return response, nil
}
