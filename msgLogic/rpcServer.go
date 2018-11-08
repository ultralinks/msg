package msgLogic

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"msg/msgLogic/parseRequest"
	"msg/msgLogic/pb/msgLogic"
	"net"
)

type server struct{}

func (s *server) ParseMsg(ctx context.Context, in *msgLogic.ParseMsgRequest) (*msgLogic.ParseMsgResponse, error) {
	data := in.Data

	parseRequest.ParseRequest(data)

	result := &msgLogic.ParseMsgResponse{
		Status: "1",
	}
	return result, nil
}

func RunRpcServer() {
	fmt.Println("start msgLogic rpc server")
	address := "0.0.0.0:10009"
	lis, err := net.Listen("tcp", address)

	log.Println("rpc server start and listen", address)

	if err != nil {
		log.Println("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	msgLogic.RegisterMsgLogicServer(s, &server{})

	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Println("failed to serve: %v", err)
	}
}
