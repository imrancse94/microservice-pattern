package pb

import (
	"fmt"
	"github.com/imrancse94/microservice-pattern/src/protobuf"
	"golang.org/x/net/context"
)

type Server struct {
	protobuf.UnsafeAuthServiceServer
}

func (s *Server) Login(ctx context.Context, req *protobuf.LoginRequest) (res *protobuf.LoginResponse, err error) {

	fmt.Println("Email", req.Email)
	fmt.Println("Password", req.Password)
	/*req := mux.Vars(r)
	fmt.Println("Test Mux", req["id"], req["name"])*/
	//request := requests.LoginRequest{}
	//Helper.Request(r, &request)
	//userData, message := Services.Login(request)
	userData := &protobuf.LoginResponse{
		Status: 100,
		Data:   nil,
		Error:  "No error",
	}

	return userData, nil

}

func (s *Server) Register(ctx context.Context, req *protobuf.RegisterRequest) (res *protobuf.RegisterResponse, error error) {

	userData := &protobuf.RegisterResponse{
		Status: 100,
		Error:  "No error",
	}

	return userData, nil
}

func (s *Server) Validate(ctx context.Context, req *protobuf.ValidateRequest) (res *protobuf.ValidateResponse, error error) {

	userData := &protobuf.ValidateResponse{
		Status: 100,
		Error:  "No error",
	}

	return userData, nil
}
