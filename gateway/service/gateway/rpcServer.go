package gateway

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"msg/gateway"
	pbGateway "msg/gateway/pb/gateway"
	"net"
)

type server struct{}

func (s *server) ReceiveSendData(ctx context.Context, in *pbGateway.SendDataRequest) (*pbGateway.SendDataResponse, error) {
	token := in.Token
	data := in.Data

	//HubObj.Sendcast <- sendData
	gateway.HubObj.Sendcast <- &gateway.SendData{
		Token: token,
		Data:  data,
	}

	result := &pbGateway.SendDataResponse{
		Status: "1",
	}
	return result, nil
}

func RunRpcServer() {
	address := "0.0.0.0:10010"
	lis, err := net.Listen("tcp", address)

	if err != nil {
		log.Println("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	pbGateway.RegisterGatewayServer(s, &server{})

	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Println("failed to serve: %v", err)
	}
}
