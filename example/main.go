package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nzlov/gin-swagger"
	"github.com/nzlov/gin-swagger/swaggerFiles"

	"github.com/nzlov/gin-swagger/example/api"
	_ "github.com/nzlov/gin-swagger/example/docs"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name X-Appkey

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 192.168.1.200:6060
// @BasePath /
func main() {
	r := gin.New()
	r.Use(gin.Logger())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/testapi/get-string-by-int/:some_id", api.GetStringByInt)

	r.Run(":6060")
}
