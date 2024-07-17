package main

import (
	"log"
	"net"

	"github.com/Exam4/4th-month-exam-Auth-service/genproto/auth"
	"github.com/Exam4/4th-month-exam-Auth-service/genproto/user"
	"github.com/Exam4/4th-month-exam-Auth-service/service"
	"github.com/Exam4/4th-month-exam-Auth-service/storage/postgres"
	"google.golang.org/grpc"
)

func main() {
	dbPostgres, err := postgres.NewPostgresStorage()
	if err != nil {
		log.Fatal("Error while connect db:", err.Error())
	}

	liss, err := net.Listen("tcp", ":9999")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	auth.RegisterAuthServiceServer(s, service.NewAuthService(dbPostgres))
	user.RegisterUserServiceServer(s, service.NewUserService(dbPostgres))

	log.Printf("server listening at %v", liss.Addr())
	if err := s.Serve(liss); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
