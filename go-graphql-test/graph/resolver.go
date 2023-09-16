package graph

import (
	"go-graphql-test/services/authservice"
	"go-graphql-test/services/emailservice"
	"go-graphql-test/services/userservice"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	UserService  userservice.UserService
	AuthService  authservice.AuthService
	EmailService emailservice.EmailService
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
