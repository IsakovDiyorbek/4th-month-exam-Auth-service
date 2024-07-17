package main

import (
	"fmt"
	"log"

	"github.com/Exam4/4th-month-exam-Auth-service/api"
	"github.com/Exam4/4th-month-exam-Auth-service/api/handler"
	_ "github.com/Exam4/4th-month-exam-Auth-service/docs"
	"github.com/Exam4/4th-month-exam-Auth-service/genproto/auth"
	"github.com/Exam4/4th-month-exam-Auth-service/genproto/user"
	"github.com/go-redis/redis/v8"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	userConn, err := grpc.NewClient(fmt.Sprintf("localhost%s", ":9999"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Error while connecting: ", err.Error())
	}
	
	defer userConn.Close()

	redisDb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	auth := auth.NewAuthServiceClient(userConn)
	userService := user.NewUserServiceClient(userConn)

	h := handler.Handler{Auth: auth, User: userService, Redis: redisDb}

	r := api.NewGin(h)
	
	fmt.Println("Server started on port: 8080")
	r.Run(":8080")
}
