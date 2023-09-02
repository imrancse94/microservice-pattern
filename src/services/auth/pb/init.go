package pb

import (
	"fmt"
	"golang.org/x/net/context"
)

type Server struct {
	UnsafeAuthServiceServer
}

func (s *Server) Login(ctx context.Context, req *LoginRequest) (res *LoginResponse, err error) {

	fmt.Println("Email", req.Email)
	fmt.Println("Password", req.Password)
	/*req := mux.Vars(r)
	fmt.Println("Test Mux", req["id"], req["name"])*/
	//request := requests.LoginRequest{}
	//Helper.Request(r, &request)
	//userData, message := Services.Login(request)
	userData := &LoginResponse{
		Status: 100,
		Data:   nil,
		Error:  "No error",
	}

	return userData, nil

}

func (s *Server) Register(ctx context.Context, req *RegisterRequest) (res *RegisterResponse, error error) {

	userData := &RegisterResponse{
		Status: 100,
		Error:  "No error",
	}

	return userData, nil
}

func (s *Server) Validate(ctx context.Context, req *ValidateRequest) (res *ValidateResponse, error error) {

	userData := &ValidateResponse{
		Status: 100,
		Error:  "No error",
	}

	return userData, nil
}
