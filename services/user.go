package services

import (
	"context"
	"fmt"
	"time"

	"github.com/codeedu/fc2-grpc/pb"
)

// type UserServiceClient interface {
// 	AddUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error)
// 	AddUserVerbose(ctx context.Context, in *User, opts ...grpc.CallOption) (UserService_AddUserVerboseClient, error)
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
