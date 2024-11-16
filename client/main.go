package main

import (
	"context"
	"log"

	"github.com/sadek-39/grpc/github.com/sadek-39/grpc/invoicer"
	"google.golang.org/grpc"
)

func main() {
    conn, err := grpc.Dial(":8089", grpc.WithInsecure())
    if err != nil {
        log.Fatalf("Failed to connect: %v", err)
    }
    defer conn.Close()

    client := invoicer.NewInvoicerClient(conn)
    response, err := client.Create(context.Background(), &invoicer.CreateRequest{
        Amount: &invoicer.Amount{Amount: 100, Currency: "USD"},
        From:   "Sender",
        To:     "Receiver",
    })
    if err != nil {
        log.Fatalf("Error calling Create: %v", err)
    }

    log.Printf("Create Response: %v", response)

	response, err = client.Update(context.Background(), &invoicer.CreateRequest{
        Amount: &invoicer.Amount{Amount: 100, Currency: "USD"},
        From:   "Sender",
        To:     "Receiver",
    })
    if err != nil {
        log.Fatalf("Error calling Create: %v", err)
    }

    log.Printf("update Response: %v", response)
}