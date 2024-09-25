package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"bff/internal/gen/greeter"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	r := gin.Default()

	// Define a route that will trigger the gRPC call
	r.GET("/sayhello/:name", func(c *gin.Context) {
		name := c.Param("name")

		// Call the gRPC client
		message, err := sendHelloToPythonServer(name)
		if err != nil {
			log.Printf("failed to call gRPC server: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to send hello"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": message})
	})

	// Run the server
	r.Run(":8080")
}

// sendHelloToPythonServer calls the Python server via gRPC
func sendHelloToPythonServer(name string) (string, error) {
	// Connect to the gRPC server (adjust the address as necessary)
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return "", fmt.Errorf("did not connect: %v", err)
	}
	defer conn.Close()

	// Create a client for the Greeter service
	client := greeter.NewGreeterClient(conn)

	// Prepare the request
	req := &greeter.HelloRequest{Name: name}

	// Call the SayHello method
	resp, err := client.SayHello(context.Background(), req)
	if err != nil {
		return "", fmt.Errorf("could not greet: %v", err)
	}

	// Return the message
	return resp.Message, nil
}
