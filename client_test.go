package client

import (
	"testing"
)

var client = New(
	&Config{
		Addr:     "http://127.0.0.1:8848",
		Username: "nacos",
		Password: "nacos",
	})

func TestClient_Health(t *testing.T) {
	err := client.Health()
	if err != nil {
		t.Error(err)
	}
}

func TestClient_Auth(t *testing.T) {
	err := client.Login()
	if err != nil {
		t.Error(err)
	}
}
