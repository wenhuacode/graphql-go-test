package handler

import (
	"casbin-test/api"
	"casbin-test/common"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

//Casbin 权限管理

type CasbinInfo struct {
	Path   string `json:"path" form:"path"`
	Method string `json:"method" form:"method"`
}
type CasbinCreateRequest struct {
	RoleId      string       `json:"role_id" form:"role_id" description:"角色ID"`
	DomainId    string       `json:"domain_id" form:"domain_id" description:"组织ID"`
	CasbinInfos []CasbinInfo `json:"casbin_infos" description:"权限模型列表"`
}

func AddPolicy(c *gin.Context) {
	log.Printf("==========")
	var params CasbinCreateRequest
	c.ShouldBind(&params)

	for _, v := range params.CasbinInfos {

		log.Println(params.RoleId, v.Path, v.Method)
		err := api.AddPolicyApi(common.CasBinDB, params.RoleId, v.Path, v.Method)
		if err != nil {
			//	c.JSON(http.StatusOK,gin.H{
			//		"res":"bad",
			//	})
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"res": "ok",
	})
}

func TestListPolicics(c *gin.Context) {

	c.JSON(http.StatusOK, map[string]string{"msg": "权限正常通过"})
}
