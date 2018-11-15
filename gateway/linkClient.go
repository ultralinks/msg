package gateway

import (
	"sync"
)

// store user and client map
var LinkDeviceMap = NewLinkDeviceMap()

type linkDeviceMap struct {
	rwMutex        sync.RWMutex
	linkKey2Device map[string][]*Client
	device2LinkKey map[*Client]string
}

func NewLinkDeviceMap() *linkDeviceMap {
	return &linkDeviceMap{
		linkKey2Device: make(map[string][]*Client),
		device2LinkKey: make(map[*Client]string),
	}
}

func (userClient *linkDeviceMap) Run() {

}

func (userClient *linkDeviceMap) Join(device string, client *Client) {
	userClient.rwMutex.Lock()
	defer userClient.rwMutex.Unlock()

	userClient.linkKey2Device[device] = append(userClient.linkKey2Device[device], client)
	userClient.device2LinkKey[client] = device

	return
}

func (userClient *linkDeviceMap) Leave(client *Client) {
	userClient.rwMutex.Lock()
	defer userClient.rwMutex.Unlock()

	device := userClient.device2LinkKey[client]
	delete(userClient.device2LinkKey, client)

	for k, v := range userClient.linkKey2Device[device] {
		if v == client {
			userClient.linkKey2Device[device] = append(userClient.linkKey2Device[device][:k], userClient.linkKey2Device[device][k+1:]...)
		}
	}

	if len(userClient.linkKey2Device[device]) == 0 {
		delete(userClient.linkKey2Device, device)
	}

	return
}
