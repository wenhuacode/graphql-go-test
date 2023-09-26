package do

import (
	"ent-three-layer/app/user/srv/internal/data/v1/ent"
)

type UserDO struct {
	ent.User
}

func (UserDO) TableName() string {
	return "user"
}

type UserDOList struct {
	TotalCount int64     `json:"totalCount,omitempty"`
	Items      []*UserDO `json:"items"`
}
