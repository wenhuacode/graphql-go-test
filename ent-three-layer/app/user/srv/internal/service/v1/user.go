package v1

import (
	"context"
	v1 "ent-three-layer/app/user/srv/internal/data/v1"

	"ent-three-layer/app/user/srv/internal/domain/dto"
	v2 "ent-three-layer/pkg/common/meta/v1"
)

type UserSrv interface {
	List(ctx context.Context, orderby []v2.OrderMeta, opts v2.ListMeta) (*dto.UserDTOList, error)
	Create(ctx context.Context, user *dto.UserDTO) error
	Update(ctx context.Context, user *dto.UserDTO) error
	GetByID(ctx context.Context, ID uint64) (*dto.UserDTO, error)
	GetByMobile(ctx context.Context, mobile string) (*dto.UserDTO, error)
}

type userService struct {
	data v1.DataFactory
}

func newUser(srv *service) *userService {
	return &userService{
		data: srv.data,
	}
}

func (us *userService) List(ctx context.Context, orderby []v2.OrderMeta, opts v2.ListMeta) (*dto.UserDTOList, error) {
	doList, err := us.data.User().List(ctx, orderby, opts)
	if err != nil {
		return nil, err
	}

	var userDTOList dto.UserDTOList
	for _, value := range doList.Items {
		projectDTO := dto.UserDTO{*value}
		userDTOList.Items = append(userDTOList.Items, &projectDTO)
	}
	userDTOList.TotalCount = doList.TotalCount
	//业务逻辑3
	return &userDTOList, nil
}

func (us *userService) Create(ctx context.Context, user *dto.UserDTO) error {
	//TODO implement me
	panic("implement me")
}

func (us *userService) Update(ctx context.Context, user *dto.UserDTO) error {
	//TODO implement me
	panic("implement me")
}

func (us *userService) GetByID(ctx context.Context, ID uint64) (*dto.UserDTO, error) {
	//TODO implement me
	panic("implement me")
}

func (us *userService) GetByMobile(ctx context.Context, mobile string) (*dto.UserDTO, error) {
	//TODO implement me
	panic("implement me")
}
