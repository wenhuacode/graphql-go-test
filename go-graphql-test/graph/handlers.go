package graph

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"go-graphql-test/services/authservice"
	"go-graphql-test/services/emailservice"
	"go-graphql-test/services/userservice"
)

// GraphqlHandler defines the GQLGen GraphQL server handler
func GrapgqlHandler(
	us userservice.UserService,
	as authservice.AuthService,
	es emailservice.EmailService) gin.HandlerFunc {
	conf := Config{
		Resolvers: &Resolver{
			UserService:  us,
			AuthService:  as,
			EmailService: es,
		},
	}
	exec := NewExecutableSchema(conf)
	h := handler.NewDefaultServer(exec)
	return func(c *gin.Context) { h.ServeHTTP(c.Writer, c.Request) }
}

// PlaygroundHandler Defines the Playground handler to expose our playground
func PlayGroundHandler(path string) gin.HandlerFunc {
	h := playground.Handler("GraphQL Playground", path)
	return func(c *gin.Context) { h.ServeHTTP(c.Writer, c.Request) }
}
