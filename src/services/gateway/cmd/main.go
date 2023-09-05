package main

import (
	"fmt"
	"gateway/pkg/bootstrap"
	"github.com/joho/godotenv"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
)

var (
	_, b, _, _ = runtime.Caller(0)
	root       = filepath.Join(filepath.Dir(b), "..")
)

func main() {
	godotenv.Load(string(root) + `/.env`)
	//DB := database.InitDB()
	//models.Init(DB)
	bootstrap.Init()

	//credentials := handlers.AllowCredentials()
	//methods := handlers.AllowedMethods([]string{"GET, POST, PATCH, PUT, DELETE, OPTIONS"})
	//ttl := handlers.MaxAge(3600)
	//headers := handlers.AllowedHeaders([]string{"content-type"})
	//origins := handlers.AllowedOrigins([]string{"localhost:3000"})
	//cache.ConnectRedis(context.Background())
	fmt.Println("Removed cache")
	//mail.SendEmail("My subject", "This is test", "", []string{"abquddus.ctg@gmail.com", "jesse.miller.2022.smtp@gmail.com"}, "test.txt")

	PORT := os.Getenv("APP_PORT")
	fmt.Println("Listening port", PORT)

	http.ListenAndServe(":"+PORT, nil)

}
