package auth

import (
	"fmt"
	"gateway/grpc/auth/config"
	pb "github.com/imrancse94/microservice-pattern/src/protobuf"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func InitServiceClient() pb.AuthServiceClient {
	c := config.Config()
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(fmt.Sprintf("localhost:%s", c.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewAuthServiceClient(cc)
}
