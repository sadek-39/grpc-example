package services

import (
	"context"
	"log"

	pb "github.com/sadek-39/grpc/github.com/sadek-39/grpc/invoicer"
)

// server implements the Invoicer service.
type Server struct {
	pb.UnimplementedInvoicerServer
}

// Create handles the Create RPC and generates a response.
func (s *Server) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	log.Printf("Received Create request: from=%s, to=%s, amount=%d %s",
		req.From, req.To, req.Amount.Amount, req.Amount.Currency)

	// Example response
	response := &pb.CreateResponse{
		Pdf:  "invoice_123.pdf",
		Docx: "invoice_123.docx",
	}
	return response, nil
}

// Update handles the Update RPC and generates a response.
func (s *Server) Update(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	log.Printf("Received Update request: from=%s, to=%s, amount=%d %s",
		req.From, req.To, req.Amount.Amount, req.Amount.Currency)

	// Example response
	response := &pb.CreateResponse{
		Pdf:  "updated_invoice_123.pdf",
		Docx: "updated_invoice_123.docx",
	}
	return response, nil
}