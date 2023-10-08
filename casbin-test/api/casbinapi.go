package api

import (
	"github.com/casbin/casbin/v2"
)

func AddPolicyApi(db *casbin.Enforcer, roleId string, path, method string) error {
	_, err := db.AddPolicy(roleId, path, method)
	if err != nil {
		return err
	}
	return nil
}
