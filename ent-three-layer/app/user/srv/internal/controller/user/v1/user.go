package v1

import (
	"context"
	proto "ent-three-layer/api/user/v1"
	v1 "ent-three-layer/api/user/v1"
	"ent-three-layer/app/user/srv/internal/domain/dto"
	v12 "ent-three-layer/app/user/srv/internal/service/v1"
	v13 "ent-three-layer/pkg/common/meta/v1"
	"ent-three-layer/pkg/log"
)

type userServer struct {
	proto.UnimplementedUserServer
	srv v12.ServiceFactory
}

func NewUserService(srv v12.ServiceFactory) v1.UserServer {
	return &userServer{srv: srv}
}

func DTOToResponse(userDTO dto.UserDTO) *proto.UserInfoResponse {
	//在grpc的message中字段有默认值，你不能随便赋值nil进去，容易出错
	//这里要搞清， 哪些字段是有默认值
	userInfoRsp := proto.UserInfoResponse{
		Id:       userDTO.ID.String(),
		PassWord: userDTO.Password,
		NickName: userDTO.Nickname,
		Gender:   userDTO.Gender,
		Role:     int32(userDTO.Role),
		Mobile:   userDTO.Mobile,
	}
	if &userDTO.Birthday != nil {
		userInfoRsp.BirthDay = uint64(userDTO.Birthday.Unix())
	}
	return &userInfoRsp
}

func (us *userServer) GetUserList(ctx context.Context, info *proto.PageInfo) (*proto.UserListResponse, error) {
	log.Info("GetUserList is called")
	srvOpts := v13.ListMeta{
		Page:     int(info.Pn),
		PageSize: int(info.PSize),
	}

	var orderBy []v13.OrderMeta
	for _, i2 := range info.Sort {
		orderBy = append(orderBy, v13.OrderMeta{
			Field: i2.Field,
			Order: int(i2.Order),
		})
	}

	dtoList, err := us.srv.User().List(ctx, orderBy, srvOpts)
	if err != nil {
		return nil, err
	}

	var rsp proto.UserListResponse
	for _, value := range dtoList.Items {
		userRsp := DTOToResponse(*value)
		rsp.Data = append(rsp.Data, userRsp)
	}
	return &rsp, nil
}

var _ v1.UserServer = &userServer{}
