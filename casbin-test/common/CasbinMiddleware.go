package common

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// casbin中间件
func CasBinMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.RequestURI() //
		method := c.Request.Method
		log.Println(path, method)
		//验证url权限
		roleId := "admin"
		ok, _ := CasBinDB.Enforce(roleId, path, method)
		if ok {
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, map[string]string{
				"msg": "很遗憾,权限验证没有通过",
			})
			c.Abort()
			return
		}
	}
}
