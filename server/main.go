package main

import (
	"log"
	"net"

	pb "github.com/sadek-39/grpc/github.com/sadek-39/grpc/invoicer"
	"github.com/sadek-39/grpc/server/services"
	"google.golang.org/grpc"
)


func main() {
	// Listen on a specified port
	listener, err := net.Listen("tcp", ":8089")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	log.Println("Server is listening on port 8089")

	// Create a new gRPC server
	grpcServer := grpc.NewServer()

	// Register the Invoicer service
	pb.RegisterInvoicerServer(grpcServer, &services.Server{})

	// Start serving requests
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
