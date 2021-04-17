package main

import (
	"context"
	"log"

	pb "github.com/protobufs/grpcEx/protofiles"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	//Set up a connection to the server
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal("Did not connect: %v", err)
	}
	defer conn.Close()

	//Create a client
	c := pb.NewMoneyTransactionClient(conn)

	//Prepare data. Get this from clients like Frontend
	from := "1234"
	to := "5678"
	amount := float32(1250.75)

	//make a server request
	r, err := c.MakeTransaction(context.Background(), &pb.TransactionRequest{From: from, To: to, Amount: amount})
	if err != nil {
		log.Fatalf("Could not transaction: %v", err)
	}
	log.Printf("Transaction confirmed: %t", r.Confirmation)

}
