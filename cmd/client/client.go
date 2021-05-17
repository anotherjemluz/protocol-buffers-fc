package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/codeedu/fc2-grpc/pb"
	"google.golang.org/grpc"
)

func main() {
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	// tratamento de erro
	if err != nil {
		log.Fatalf("Could not connect to gRPC Server: %v", err)
	}

	// fechar a conexão apenas ao final de todos os processos
	defer connection.Close()

	client := pb.NewUserServiceClient(connection)
	// AddUser(client)
	AddUserVerbose(client)
}

func AddUser(client pb.UserServiceClient) {
	// requisição teste
	req := &pb.User{
		Id:    "0",
		Name:  "Juniu",
		Email: "junio@ju.com",
	}

	res, err := client.AddUser(context.Background(), req)
	// tratamento de erro
	if err != nil {
		log.Fatalf("Could not make gRPC request: %v", err)
	}

	fmt.Println(res)
}

// stream
func AddUserVerbose(client pb.UserServiceClient) {

	req := &pb.User{
		Id:    "0",
		Name:  "Juniu",
		Email: "junio@ju.com",
	}

	responseStream, err := client.AddUserVerbose(context.Background(), req)
	// tratamento de erro
	if err != nil {
		log.Fatalf("Could not make gRPC request: %v", err)
	}

	for {
		stream, err := responseStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Could not receive the message: %v", err)
		}
		fmt.Println("Status:", stream.Status, " - ", stream.GetUser())
	}
}
