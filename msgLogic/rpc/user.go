package rpc

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "msg/msgLogic/pb/userLogic"
)

func FetchApp(appKey string) *pb.AppReply {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	// 调用方法
	reqBody := new(pb.AppRequest)
	reqBody.AppKey = appKey

	start := time.Now()
	app, err := UserRpcClient.FetchApp(ctx, reqBody)
	if err != nil {
		log.Println("rpc fetch app error", err)
	}
	elapsed := time.Since(start)
	fmt.Println("message rpc execute duration", elapsed)
	return app
}
