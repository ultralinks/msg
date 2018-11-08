package msgLogic

import (
	"context"
	"log"
	"msg/gateway/app"
	pbMsgLogic "msg/gateway/pb/msgLogic"
	"time"
)

func ParseMsg(request []byte) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	parseMsgRequest := pbMsgLogic.ParseMsgRequest{
		Data: request,
	}

	log.Println("parseMsg start")
	app.MsgLogicRpcClient.ParseMsg(ctx, &parseMsgRequest)
	log.Println("parseMsg end")

	return
}
