package gateway

import (
	"sync"
)

// store user and client map
var LinkClientMap = NewLinkClientMap()

type linkClientMap struct {
	rwMutex        sync.RWMutex
	linkKey2Client map[string][]*Client
	client2LinkKey map[*Client]string
}

func NewLinkClientMap() *linkClientMap {
	return &linkClientMap{
		linkKey2Client: make(map[string][]*Client),
		client2LinkKey: make(map[*Client]string),
	}
}

func (userClient *linkClientMap) Run() {

}

func (userClient *linkClientMap) Join(token string, client *Client) {
	userClient.rwMutex.Lock()
	defer userClient.rwMutex.Unlock()

	userClient.linkKey2Client[token] = append(userClient.linkKey2Client[token], client)
	userClient.client2LinkKey[client] = token

	return
}

func (userClient *linkClientMap) Leave(client *Client) {
	userClient.rwMutex.Lock()
	defer userClient.rwMutex.Unlock()

	token := userClient.client2LinkKey[client]
	delete(userClient.client2LinkKey, client)

	for k, v := range userClient.linkKey2Client[token] {
		if v == client {
			userClient.linkKey2Client[token] = append(userClient.linkKey2Client[token][:k], userClient.linkKey2Client[token][k+1:]...)
		}
	}

	if len(userClient.linkKey2Client[token]) == 0 {
		delete(userClient.linkKey2Client, token)
	}

	return
}
