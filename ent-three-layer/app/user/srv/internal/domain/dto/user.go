package dto

import (
	"ent-three-layer/app/user/srv/internal/domain/do"
)

type UserDTO struct {
	do.UserDO
}

type UserDTOList struct {
	TotalCount int64      `json:"totalCount,omitempty"` //总数
	Items      []*UserDTO `json:"data"`                 //数据
}
