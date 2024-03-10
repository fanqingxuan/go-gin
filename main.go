package main

import (
	"flag"
	"fmt"
	"go-gin/handlers"
	"go-gin/middlewares"
	"go-gin/svc"

	"github.com/gin-gonic/gin"
)

var configFile = flag.String("f", "./.env.yaml", "the config file")

func main() {
	flag.Parse()

	server := gin.New()
	server.HandleMethodNotAllowed = true
	serverCtx := svc.NewServiceContext()

	middlewares.RegisterGlobalMiddlewares(server)

	handlers.RegisterHandlers(server, serverCtx)

	port := ":8080"
	fmt.Printf("Starting server at localhost%s...\n", port)

	if err := server.Run(port); err != nil {
		fmt.Printf("Start server error,err=%v", err)
	}
}
