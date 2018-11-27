package rpc

import (
	"context"
	"log"
	"time"

	pbMsgLogic "msg/gateway/pb/msgLogic"
)

func ParseMsg(request []byte) (linkKeys []string, data []byte) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	parseMsgRequest := pbMsgLogic.ParseMsgRequest{
		Data: request,
	}

	response, err := MsgLogicRpcClient.ParseMsg(ctx, &parseMsgRequest)
	if err != nil {
		log.Println("rpc parseMsg error", err)

		return nil, nil
	}

	return response.LinkKeys, response.Data
}

func GetLinkKeyByToken(token string) (string, error) {
	if token == "" {
		return "", nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	getLinkByTokenRequest := pbMsgLogic.GetLinkByTokenRequest{
		Token: token,
	}
	response, err := MsgLogicRpcClient.GetLinkByToken(ctx, &getLinkByTokenRequest)

	return response.LinkKey, err
}
