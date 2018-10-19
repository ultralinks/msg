package openapi

// receive sendData from msgLogic
func ReceiveSendData(sendData SendData) {
	HubObj.sendcast <- sendData
}
