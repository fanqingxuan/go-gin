package main

import (
	"flag"
	"fmt"
	"go-gin/handlers"

	"github.com/gin-gonic/gin"
)

var configFile = flag.String("f", "./.env.yaml", "the config file")

func main() {
	flag.Parse()
	server := gin.New()
	server.HandleMethodNotAllowed = true

	handlers.RegisterHandlers(server)

	port := ":8080"
	fmt.Printf("Starting server at localhost%s...\n", port)

	if err := server.Run(port); err != nil {
		fmt.Printf("Start server error,err=%v", err)
	}
}
