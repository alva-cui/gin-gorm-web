package controller

import (
	"gin-gorm-web/internal/service"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	// 用户路由
	r.GET("/users", service.GetUsers)
	r.GET("/users/:id", service.GetUser)
	r.POST("/users", service.CreateUser)
	r.PUT("/users/:id", service.UpdateUser)
	r.DELETE("/users/:id", service.DeleteUser)
}