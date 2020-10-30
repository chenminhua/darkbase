package client

import (
	"log"
	"testing"
)

const (
	address = "localhost:12345"
)

func TestClient(t *testing.T) {
	client, err := NewClient(address)
	if client == nil {
		panic("client is nil")
	}
	err = client.Put("tb1", "foo", []byte("bar"))
	if err != nil {
		log.Fatalf("put err: %v", err)
	}
	r, err := client.Get("tb1", "foo")
	if err != nil {
		log.Fatalf("put err: %v", err)
	}
	log.Printf("get response : %v", r)
}
