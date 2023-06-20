package graph

import (
	"github.com/iBoBoTi/graphql-go-server/repository"
)

//go:generate go run github.com/99designs/gqlgen

type Resolver struct{
	 VideoRepo repository.VideoRepository
}
