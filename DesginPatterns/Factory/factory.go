/*
The Factory Pattern is a creational design pattern that encapsulates object creation. Instead of directly creating concrete objects, the client calls a factory function, which returns the appropriate implementation, usually through an interface. In Go, factory functions commonly use the NewX() naming convention. The pattern helps reduce coupling and makes it easier to add or replace implementations.
*/

package main

import (
	"fmt"
)

type s3 interface {
	Upload()
}

type AWSClient struct{}
type NoobaaClient struct{}
type MiniIOClient struct{}

func (aws *AWSClient) Upload() {
	fmt.Println("I am in aws upload...!!")
}

func (n *NoobaaClient) Upload() {
	fmt.Println("I am in noobaa upload...!!")
}

func (m *MiniIOClient) Upload() {
	fmt.Println("I am in Minio Upload...!!")
}

func NewS3(provider string) (s3, error) {
	switch provider {
	case "aws":
		return &AWSClient{}, nil
	case "noobaa":
		return &NoobaaClient{}, nil
	case "minio":
		return &MiniIOClient{}, nil
	default:
		return nil, fmt.Errorf("I am in default..")
	}

}

func main() {
	store, err := NewS3("minio")
	if err != nil {
		return
	}
	store.Upload()

}
