package graph

import (
	"context"
	"errors"
	"go-graphql-test/domain/user"
	"go-graphql-test/graph/model"
)

// Register is the resolver for the register field.
func (r *mutationResolver) Register(ctx context.Context, input model.RegisterLogin) (*model.RegisterLoginOutput, error) {
	userDomain := &user.User{
		Email:    input.Email,
		Password: input.Password,
	}

	err := r.UserService.Create(userDomain)
	if err != nil {
		return nil, err
	}

	token, err := r.AuthService.IssueToken(*userDomain)
	if err != nil {
		return nil, err
	}

	return &model.RegisterLoginOutput{
		Token: token,
		User: &model.User{
			ID:        int(userDomain.ID),
			FirstName: userDomain.FirstName,
			LastName:  userDomain.LastName,
			Email:     userDomain.Email,
			Role:      userDomain.Role,
			Active:    userDomain.Active,
		},
	}, nil

}

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, input model.RegisterLogin) (*model.RegisterLoginOutput, error) {
	usr, err := r.UserService.GetByEmail(input.Email)
	if err != nil {
		return nil, err
	}
	err = r.UserService.ComparePassword(input.Password, usr.Password)
	if err != nil {
		return nil, err
	}
	token, err := r.AuthService.IssueToken(*usr)
	if err != nil {
		return nil, err
	}

	return &model.RegisterLoginOutput{
		Token: token,
		User: &model.User{
			ID:        int(usr.ID),
			FirstName: usr.FirstName,
			LastName:  usr.LastName,
			Email:     usr.Email,
			Role:      usr.Role,
			Active:    usr.Active,
		},
	}, nil

}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, input model.UpdateUser) (*model.User, error) {
	userId := ctx.Value("user_id")
	if userId == nil {
		return nil, errors.New("unauthorized: Token is invlaid")
	}

	usr, err := r.UserService.GetByID(userId.(uint))
	if err != nil {
		return nil, err
	}

	if input.Email != "" {
		usr.Email = input.Email
	}
	if input.FirstName != nil {
		usr.FirstName = *input.FirstName
	}
	if input.LastName != nil {
		usr.LastName = *input.LastName
	}
	err = r.UserService.Update(usr)
	if err != nil {
		return nil, err
	}
	return &model.User{
		ID:        int(usr.ID),
		FirstName: usr.FirstName,
		LastName:  usr.LastName,
		Email:     usr.Email,
		Role:      usr.Role,
		Active:    usr.Active,
	}, nil
}

// ForgotPassword is the resolver for the forgotPassword field.
func (r *mutationResolver) ForgotPassword(ctx context.Context, email string) (bool, error) {
	if email == "" {
		return false, errors.New("email is required")
	}

	token, err := r.UserService.InitiateResetPassowrd(email)
	if err != nil {
		return false, err
	}

	if err = r.EmailService.ResetPassword(email, token); err != nil {
		return false, err
	}
	return true, nil
}

// ResetPassword is the resolver for the resetPassword field.
func (r *mutationResolver) ResetPassword(ctx context.Context, resetToken string, password string) (*model.RegisterLoginOutput, error) {
	if resetToken == "" {
		return nil, errors.New("token is required")
	}

	if password == "" {
		return nil, errors.New("new password is required")
	}

	usr, err := r.UserService.CompleteUpdatePassword(resetToken, password)
	if err != nil {
		return nil, err
	}

	token, err := r.AuthService.IssueToken(*usr)
	if err != nil {
		return nil, err
	}

	return &model.RegisterLoginOutput{
		Token: token,
		User: &model.User{
			ID:        int(usr.ID),
			FirstName: usr.FirstName,
			LastName:  usr.LastName,
			Email:     usr.Email,
			Role:      usr.Role,
			Active:    usr.Active,
		},
	}, nil
}
