package services

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/codeedu/fc2-grpc/pb"
)

// type UserServiceClient interface {
// 	AddUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error)
// 	AddUserVerbose(ctx context.Context, in *User, opts ...grpc.CallOption) (UserService_AddUserVerboseClient, error)
// AddUsers(ctx context.Context, opts ...grpc.CallOption) (UserService_AddUsersClient, error)
// }

type UserService struct {
	pb.UnimplementedUserServiceServer
}

func NewUserService() *UserService {
	return &UserService{}
}

// registra um metodo
func (*UserService) AddUser(ctx context.Context, req *pb.User) (*pb.User, error) {

	// Insert - Database
	fmt.Println(req.Name)

	// retorno do metodo
	return &pb.User{
		Id:    "123",
		Name:  req.GetName(),
		Email: req.GetEmail(),
	}, nil

}

// registra um metodo stream
func (*UserService) AddUserVerbose(req *pb.User, stream pb.UserService_AddUserVerboseServer) error {

	// envio
	stream.Send(&pb.UserResultStream{
		Status: "Init",
		User:   &pb.User{},
	})

	// soneca
	time.Sleep(time.Second * 3)

	// envio
	stream.Send(&pb.UserResultStream{
		Status: "Inserting",
		User:   &pb.User{},
	})

	// soneca
	time.Sleep(time.Second * 3)

	// envio
	stream.Send(&pb.UserResultStream{
		Status: "User has been inserted",
		User: &pb.User{
			Id:    "123",
			Name:  req.GetName(),
			Email: req.GetEmail(),
		},
	})

	// soneca
	time.Sleep(time.Second * 3)

	// envio
	stream.Send(&pb.UserResultStream{
		Status: "Completed",
		User: &pb.User{
			Id:    "123",
			Name:  req.GetName(),
			Email: req.GetEmail(),
		},
	})

	// soneca
	time.Sleep(time.Second * 3)

	return nil

}

func (*UserService) AddUsers(stream pb.UserService_AddUsersServer) error {

	// cria uma lista vazia de usuários
	users := []*pb.User{}

	// começa a receber o streaming num loop infinito
	for {
		req, err := stream.Recv()

		// se parar de receber (se o cliente para de enviar informação)
		if err == io.EOF {
			// envia a lista de usuários em seu estado atual
			return stream.SendAndClose(&pb.Users{
				User: users,
			})
		}

		// tratamento de erros
		if err != nil {
			log.Fatalf("Error receiving stream: %v", err)
		}

		// ao fim do for ele vai pegar esse user abaixo e atualizar a lista de usuários usando append
		users = append(users, &pb.User{
			Id:    req.GetId(),
			Name:  req.GetName(),
			Email: req.GetEmail(),
		})

		fmt.Println("Adding", req.GetName())
	}
}
