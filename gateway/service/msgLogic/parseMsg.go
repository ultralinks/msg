package msgLogic

import (
	"context"
	"log"
	"msg/gateway/app"
	pbMsgLogic "msg/gateway/pb/msgLogic"
	"time"
)

func ParseMsg(request []byte) (linkKeys []string, data []byte) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	parseMsgRequest := pbMsgLogic.ParseMsgRequest{
		Data: request,
	}

	response, err := app.MsgLogicRpcClient.ParseMsg(ctx, &parseMsgRequest)
	if err != nil {
		log.Println("rpc parseMsg error", err)

		return nil, nil
	}

	return response.LinkKeys, response.Data
}
