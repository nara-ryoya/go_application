package presenter

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/nara-ryoya/go_application/api/internal/controller/system"
	"github.com/nara-ryoya/go_application/api/internal/controller/user"
)

const latest = "/v1"

type Server struct {}

func (*s Server) Run(ctx *context.Context) {
	r := gin.Default()
	v1 := r.Group(latest)

	// 死活管理機能
	{
		systemHandler := system.NewSystemhandler()
		v1.GET("/health", systemHandler.Health)
	}

	// ユーザー管理機能
	{
		userHandler := user.NewUserHandler()
		v1.GET("",userHandler.GetUsers)
		v1.GET("/:id",userHandler.GetUserById)
		v1.POST("",userHandler.EditUser)
		v1.DELETE("/:id",userHandler.DeleteUser)
	}

	err := r.Run()

	if err != nil {
		return err
	}

	return nil
}

func NewServer() *Server {
	return &Server{}
}