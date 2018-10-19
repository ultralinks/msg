package gateway

import (
	"sync"
)

// store user and client map
var UserClientMap = NewUserClientMap()

type userClientMap struct {
	rwMutex      sync.RWMutex
	token2Client map[string][]*Client
	client2Token map[*Client]string
}

func NewUserClientMap() *userClientMap {
	return &userClientMap{
		token2Client: make(map[string][]*Client),
		client2Token: make(map[*Client]string),
	}
}

func (userClient *userClientMap) Run(){

}

func (userClient *userClientMap) Join(token string, client *Client) {
	userClient.rwMutex.Lock()
	defer userClient.rwMutex.Unlock()

	userClient.token2Client[token] = append(userClient.token2Client[token], client)
	userClient.client2Token[client] = token

	return
}

func (userClient *userClientMap) Leave(client *Client) {
	userClient.rwMutex.Lock()
	defer userClient.rwMutex.Unlock()

	token := userClient.client2Token[client]
	delete(userClient.client2Token, client)

	for k, v := range userClient.token2Client[token] {
		if v == client {
			userClient.token2Client[token] = append(userClient.token2Client[token][:k], userClient.token2Client[token][k+1:]...)
		}
	}

	if len(userClient.token2Client[token]) == 0 {
		delete(userClient.token2Client, token)
	}

	return
}
