package grpc

import (
	. "github.com/Andrew161644/full_app/grpc_service/grpc/database/models"
)

func MapToUserModel(user User) UserModel {
	return UserModel{
		ID:     int(user.ID),
		Name:   user.Name,
		RoleId: int(user.RoleId),
	}
}
func MapToUserGrpc(user UserModel) *User {
	return &User{
		ID:     int64(user.ID),
		Name:   user.Name,
		RoleId: int64(user.RoleId),
	}
}
