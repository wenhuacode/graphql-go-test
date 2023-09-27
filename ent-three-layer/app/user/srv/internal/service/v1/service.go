package v1

import v1 "ent-three-layer/app/user/srv/internal/data/v1"

type ServiceFactory interface {
	User() UserSrv
}

type service struct {
	data v1.DataFactory
}

func (s *service) User() UserSrv {
	return newUser(s)
}

func NewService(store v1.DataFactory) *service {
	return &service{data: store}
}

var _ ServiceFactory = &service{}
