package rpc

import (
	"fmt"
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	userLogic "msg/userLogic/app"
	pb "msg/userLogic/pb/userLogic"
	"msg/userLogic/service/app"
)

type server struct{}

func (s *server) FetchApp(ctx context.Context, in *pb.AppRequest) (*pb.AppReply, error) {
	re := pb.AppReply{}
	app.FetchByKey(in.AppKey, &re)
	return &re, nil
}

func RunRpcServer() {
	fmt.Println("start msgLogic rpc server")
	address := userLogic.Config.Rpc.Domain + ":" + userLogic.Config.Rpc.Port
	lis, err := net.Listen("tcp", address)

	log.Println("rpc server start and listen", address)

	if err != nil {
		log.Println("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterUserServer(s, &server{})

	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Println("failed to serve: %v", err)
	}
}
