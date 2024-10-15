package db

import (
  "fmt"

  client "github.com/weaviate/weaviate-go-client/v4/weaviate"
)


type Weaviate struct {
	Conn *client.Client
}

func NewWeaviate() *Weaviate {
  config := client.Config{
	Scheme: "http",
	Host:   "localhost:9000",
  }
  c, err := client.NewClient(config)
  if err != nil {
	fmt.Printf("Error occurred %v", err)
	return nil
  }
  return &Weaviate{Conn: c}
}

