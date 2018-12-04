package rpc

import (
	"fmt"
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"msg/msgLogic/app"
	"msg/msgLogic/parseRequest"
	"msg/msgLogic/pb/msgLogic"
	linkService "msg/msgLogic/service/link"
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

func (s *server) GetLinkByToken(ctx context.Context, r *msgLogic.GetLinkByTokenRequest) (*msgLogic.GetLinkByTokenResponse, error) {
	link, err := linkService.GetByToken(r.Token)
	res := &msgLogic.GetLinkByTokenResponse{
		LinkId:  link.Id,
		LinkKey: link.Key,
		Nick:    link.Nick,
		Avt:     link.Avt,
		AppId:   link.AppId,
	}
	return res, err
}

func RunRpcServer() {
	fmt.Println("start msgLogic rpc server")
	address := app.Config.Rpc.Domain + ":" + app.Config.Rpc.Port
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
