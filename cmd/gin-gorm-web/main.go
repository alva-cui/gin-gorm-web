// @title Gin GORM Web API
// @version 1.0
// @description This is a sample Gin GORM web server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

package main

import (
	"gin-gorm-web/internal/config"
	"gin-gorm-web/internal/controller"
	"gin-gorm-web/internal/dao"
	"gin-gorm-web/internal/middlewares"
	_ "gin-gorm-web/docs" // 导入docs包以初始化Swagger

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	// 连接数据库
	config.ConnectDatabase()

	// 自动迁移模型
	config.DB.AutoMigrate(&dao.User{})

	// 创建gin引擎
	r := gin.New()

	// 使用中间件
	r.Use(gin.Recovery())
	r.Use(middlewares.LoggerMiddleware())
	r.Use(middlewares.ErrorHandlerMiddleware())

	// 设置路由
	controller.SetupRoutes(r)

	// 添加Swagger路由
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 启动服务器
	r.Run(":8080")
}