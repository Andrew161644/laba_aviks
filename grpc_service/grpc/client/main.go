package main

import (
	"context"
	"fmt"
	. "github.com/Andrew161644/full_app/grpc_service/grpc"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err := grpc.Dial(":8081", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	c := NewDataServiceClient(conn)
	res, err := c.GetUsers(context.Background(), &GetUserRequest{
		Id: 2,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res.User)
}
