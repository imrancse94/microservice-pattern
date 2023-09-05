package auth

import (
	"fmt"
	pb "github.com/imrancse94/microservice-pattern/src/protobuf/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func InitServiceClient() pb.AuthServiceClient {
	c := Config()
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(fmt.Sprintf("%s:%s", c.Domain, c.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewAuthServiceClient(cc)
}
