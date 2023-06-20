package http

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/iBoBoTi/graphql-go-server/graph"
	"github.com/iBoBoTi/graphql-go-server/repository"
)


func PlaygroundHandler() gin.HandlerFunc{
	h := playground.Handler("GraphQL playground", "/query")
	return func(c *gin.Context){
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func GraphqlHandler(videoRepo repository.VideoRepository) gin.HandlerFunc{
	
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		VideoRepo: videoRepo,
	}}))
	return func(c *gin.Context){
		srv.ServeHTTP(c.Writer, c.Request)
	}
}