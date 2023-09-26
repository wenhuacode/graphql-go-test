package dto

import (
	"ent-three-layer/app/user/srv/internal/domain/do"
)

type userDTO struct {
	do.UserDO
}

type UserDTOList struct {
	TotalCount int64      `json:"totalCount,omitempty"` //总数
	Items      []*userDTO `json:"data"`                 //数据
}
