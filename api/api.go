package api

import (
	"github.com/Exam4/4th-month-exam-Auth-service/api/handler"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewGin(h handler.Handler) *gin.Engine {
	r := gin.Default()

	r.GET("/api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	auth := r.Group("/auth")
	{
		auth.POST("/register", h.Register)
		auth.POST("/login", h.Login)
		auth.POST("/logout", h.Logout)
		auth.POST("/forgot", h.ForgotPassword)
		auth.POST("/reset", h.ResetPassword)
	}


	user := r.Group("/user")
	{
		user.GET("/profile/:id", h.GetProfile)
		user.PUT("/profile", h.UpdateProfile)
		user.PUT("/password", h.ChangePassword)
	}

	return r
}
