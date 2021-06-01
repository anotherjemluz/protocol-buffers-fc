package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

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
	// AddUserVerbose(client)
	// AddUsers(client)
	AddUserStreamBoth(client)
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

func AddUsers(client pb.UserServiceClient) {
	reqs := []*pb.User{
		&pb.User{
			Id:    "w1",
			Name:  "wesley 1",
			Email: "wes@ley1.com",
		},
		&pb.User{
			Id:    "w2",
			Name:  "wesley 2",
			Email: "wes@ley2.com",
		},
		&pb.User{
			Id:    "w3",
			Name:  "wesley 3",
			Email: "wes@ley3.com",
		},
		&pb.User{
			Id:    "w4",
			Name:  "wesley 4",
			Email: "wes@ley4.com",
		},
		&pb.User{
			Id:    "w5",
			Name:  "wesley 5",
			Email: "wes@ley5.com",
		},
	}

	// o context.Background garante que se a mensagem não chegar ele já para as requisições,
	// controlando o fluxo de dados
	stream, err := client.AddUsers(context.Background())
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	// percarre cada um dos items do array e começa a enviar as requisições
	for _, req := range reqs {
		stream.Send(req)
		time.Sleep(time.Second * 3)
	}

	// stream.closeAndRecv interrompe o encio dos dados e pede o envio de uma resposta
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error creating response: %v", err)
	}

	fmt.Println(res)
}

func AddUserStreamBoth(client pb.UserServiceClient) {

	stream, err := client.AddUserStreamBoth(context.Background())
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	reqs := []*pb.User{
		&pb.User{
			Id:    "w1",
			Name:  "wesley 1",
			Email: "wes@ley1.com",
		},
		&pb.User{
			Id:    "w2",
			Name:  "wesley 2",
			Email: "wes@ley2.com",
		},
		&pb.User{
			Id:    "w3",
			Name:  "wesley 3",
			Email: "wes@ley3.com",
		},
		&pb.User{
			Id:    "w4",
			Name:  "wesley 4",
			Email: "wes@ley4.com",
		},
		&pb.User{
			Id:    "w5",
			Name:  "wesley 5",
			Email: "wes@ley5.com",
		},
	}

	wait := make(chan int)

	go func() {

		for _, req := range reqs {
			fmt.Println("Sending user:", req.Name)
			stream.Send(req)
			time.Sleep(time.Second * 2)
		}

		stream.CloseSend()

	}()

	go func() {

		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatalf("Error receiving data: %v", err)
				break
			}
			fmt.Printf("Recebendo user %v com status: %v \n", res.GetUser().GetName(), res.GetStatus())
		}
		close(wait)
	}()

	<-wait

}
