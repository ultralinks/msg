package gateway

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type server struct{}

func (s *server) ReceiveSendData(ctx context.Context, in *SendDataRequest) (*SendDataResponse, error) {
	token := in.Token
	data := in.Data

	HubObj.sendcast <- SendData{
		token: token,
		data:  data,
	}

	result := &SendDataResponse{
		Status: "1",
	}
	return result, nil
}

func RunRpcServer() {
	fmt.Println("start gateway rpc server")
	address := "0.0.0.0:10000"
	lis, err := net.Listen("tcp", address)

	log.Println("rpc server start and listen", address)

	if err != nil {
		log.Println("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	RegisterGatewayServer(s, &server{})

	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Println("failed to serve: %v", err)
	}
}
