package main

import (
	"log"
	"net"

	"github.com/codeedu/fc2-grpc/pb"
	"github.com/codeedu/fc2-grpc/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// o pacote net já vem embutido no go, e ele tem uma função
	// de listerner para escutar eventos

	lis, err := net.Listen("tcp", "localhost:50051")

	// escutando a porta
	if err != nil {
		log.Fatal("Could not connect: %v", err)
	}

	// criar o servidor
	grpcServer := grpc.NewServer()
	// registra o serviço
	pb.RegisterUserServiceServer(grpcServer, services.NewUserService())
	reflection.Register(grpcServer)

	// servidor escutando
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("Could not serve: %v", err)
	}
}
