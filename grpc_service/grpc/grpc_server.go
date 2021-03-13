package grpc

import (
	"context"
	"github.com/Andrew161644/full_app/grpc_service/grpc/database/providers"
	"log"
)

type GrpcServer struct {
	DataBase *providers.Database
}

func (server *GrpcServer) mustEmbedUnimplementedDataServiceServer() {
	panic("implement me")
}

func (server *GrpcServer) GetUsers(ctx context.Context, in *GetUserRequest) (*GetUserResponse, error) {
	var user, err = server.DataBase.GetUserById(int(in.Id))
	if err != nil {
		log.Fatal(err)
	}
	return &GetUserResponse{
		User: MapToUserGrpc(user),
	}, nil
}
