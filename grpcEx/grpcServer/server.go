package main

import (
	"log"
	"net"

	pb "github.com/protobufs/grpcEx/protofiles"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

//used to create MoneyTransactionServer
type server struct {
}

//MakeTransaction implements MoneyTransactionServer.MakeTransaction
func (s *server) MakeTransaction(ctx context.Context, in *pb.TransactionRequest) (*pb.TransactionResponse, error) {

	log.Printf("Got request for money Transfer...")
	log.Printf("Amount: %f, From A/c:%s, To A/c:%s", in.Amount, in.From, in.To)
	//user in.Amount, in.Form, in.To and perform transaction logic
	return &pb.TransactionResponse{Confirmation: true}, nil
}

const (
	port = ":50051"
)

func main() {
	lis, err := net.Listen("tcp", port)
	s := grpc.NewServer()
	pb.RegisterMoneyTransactionServer(s, &server{})
	reflection.Register(s)
	if err = s.Serve(lis); err != nil {
		log.Fatal("Fatal to serve: %v", err)
	}
}
