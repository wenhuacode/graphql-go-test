package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.37

import (
	"context"
	"errors"
	"go-graphql-test/graph/model"
)

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id int) (*model.User, error) {
	user, err := r.UserService.GetByID(uint(id))
	if err != nil {
		return nil, err
	}
	return &model.User{
		ID:        int(user.ID),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Role:      user.Role,
		Active:    user.Active,
	}, nil
}

// UserProfile is the resolver for the userProfile field.
func (r *queryResolver) UserProfile(ctx context.Context) (*model.User, error) {
	userID := ctx.Value("user_id")
	if userID == nil {
		return nil, errors.New("Unauthorized: Token is invlaid")
	}

	user, err := r.UserService.GetByID(userID.(uint))
	if err != nil {
		return nil, err
	}
	return &model.User{
		ID:        int(user.ID),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Role:      user.Role,
		Active:    user.Active,
	}, nil
}
