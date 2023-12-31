package do

import (
	"ent-three-layer/app/user/srv/internal/data/v1/ent"
)

type UserDO struct {
	ent.User
}

type UserDOList struct {
	TotalCount int64     `json:"totalCount,omitempty"`
	Items      []*UserDO `json:"items"`
}

func (UserDO) TableName() string {
	return "user"
}
