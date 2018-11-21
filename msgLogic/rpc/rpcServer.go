package rpc

import (
	"fmt"
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"msg/msgLogic/parseRequest"
	"msg/msgLogic/pb/msgLogic"
)

type server struct{}

func (s *server) ParseMsg(ctx context.Context, in *msgLogic.ParseMsgRequest) (*msgLogic.ParseMsgResponse, error) {
	data := in.Data

	linkKeys, data, err := parseRequest.ParseRequest(data)
	if err != nil {
		return nil, err
	}

	result := &msgLogic.ParseMsgResponse{
		LinkKeys: linkKeys,
		Data:     data,
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
