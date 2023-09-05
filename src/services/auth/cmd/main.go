package main

import (
	"auth/controllers"
	"auth/models"
	"auth/pkg/bootstrap"
	"auth/pkg/cache"
	"auth/pkg/database"
	"context"
	"fmt"
	"github.com/imrancse94/microservice-pattern/src/protobuf/auth"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"log"
	"net"
	"path/filepath"
	"runtime"
)

var (
	_, b, _, _ = runtime.Caller(0)
	root       = filepath.Join(filepath.Dir(b), "..")
)

func main() {
	godotenv.Load(string(root) + `/.env`)
	DB := database.InitDB()
	models.Init(DB)
	bootstrap.Init()

	//credentials := handlers.AllowCredentials()
	//methods := handlers.AllowedMethods([]string{"GET, POST, PATCH, PUT, DELETE, OPTIONS"})
	//ttl := handlers.MaxAge(3600)
	//headers := handlers.AllowedHeaders([]string{"content-type"})
	//origins := handlers.AllowedOrigins([]string{"localhost:3000"})
	cache.ConnectRedis(context.Background())

	//mail.SendEmail("My subject", "This is test", "", []string{"abquddus.ctg@gmail.com", "jesse.miller.2022.smtp@gmail.com"}, "test.txt")

	/*PORT := os.Getenv("APP_PORT")
	fmt.Println("Listening port", PORT)

	http.ListenAndServe(":"+PORT, nil)*/
	fmt.Println("This is auth")

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 9000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	s := controllers.Server{}
	auth.RegisterAuthServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

	fmt.Println("err", err)

}
