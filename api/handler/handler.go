package handler

import (
	"github.com/Exam4/4th-month-exam-Auth-service/genproto/auth"
	"github.com/Exam4/4th-month-exam-Auth-service/genproto/user"
	"github.com/go-redis/redis/v8"
)



type Handler struct{
	Auth auth.AuthServiceClient
	User user.UserServiceClient
	Redis *redis.Client
}

func NewHandler(auth auth.AuthServiceClient, user user.UserServiceClient, redis *redis.Client) *Handler {
	return &Handler{
		Auth: auth,
		User: user,
		Redis: redis,
	}
}