package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/iBoBoTi/graphql-go-server/http"
	"github.com/iBoBoTi/graphql-go-server/middleware"
	"github.com/iBoBoTi/graphql-go-server/repository"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	videoRepo := repository.New()

	server := gin.Default()

	server.Use(middleware.BasicAuth())
	server.GET("/", http.PlaygroundHandler())
	server.POST("/query", http.GraphqlHandler(videoRepo)) //graphql playground


	log.Printf("connect to http://localhost:%s/ for GraphQL playground", defaultPort)
	server.Run(":"+defaultPort)

	// http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	// http.Handle("/query", srv)

	// log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	// log.Fatal(http.ListenAndServe(":"+port, nil))
}
