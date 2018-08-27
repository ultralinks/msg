package test

import (
	"testing"
	"msg/test/client"
)

func TestLogin(t *testing.T) {
	t.Log("test login start")
	client := client.NewClient()
	client.WriteMessage(1, []byte("{type:\"login\",msgType:\"text\",msg:\"123456789\"}"))
	client.Close()
}
