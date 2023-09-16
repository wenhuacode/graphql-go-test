package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-graphql-test/common/hmachash"
	"go-graphql-test/common/randomstring"
	pwdDomain "go-graphql-test/domain/passwordreset"
	"go-graphql-test/domain/user"
	"go-graphql-test/global"
	"go-graphql-test/graph"
	"go-graphql-test/infra/mailgunclient"
	"go-graphql-test/initialize"
	"go-graphql-test/middlewares"
	"go-graphql-test/repositories/passwordreset"
	"go-graphql-test/repositories/userrepo"
	"go-graphql-test/services/authservice"
	"go-graphql-test/services/emailservice"
	"go-graphql-test/services/userservice"
	"net/http"
)

var (
	router = gin.Default()
)

func Run() {

	//初始化
	initialize.InitLogger()
	initialize.InitConfig()
	initialize.InitDB()

	err := global.DB.AutoMigrate(&user.User{}, &pwdDomain.PasswordReset{})
	if err != nil {
		panic(err)
	}

	/*
		=======  Setup Common ============
	*/
	rds := randomstring.NewRandomString()
	hm := hmachash.NewHMAC("1123")

	/*
		====== Setup infra ==============
	*/
	emailClient := mailgunclient.NewMailgunClient(global.Config)

	/*
		====== Setup repositories =======
	*/
	userRepo := userrepo.NewUserRepo(global.DB)
	pwdRepo := passwordreset.NewPasswordResetRepo(global.DB)

	/*
		====== Setup services ===========
	*/
	userService := userservice.NewUserService(userRepo, pwdRepo, rds, hm, global.Config.Pepper)
	authService := authservice.NewAuthService(global.Config.JWTSecret)
	emailservice := emailservice.NewEmailService(emailClient)

	/*
		====== Setup middlewares ========
	*/
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	/*
		====== Setup routes =============
	*/
	router.GET("/ping", func(c *gin.Context) { c.String(http.StatusOK, "pong") })

	router.GET("/graphql", graph.PlayGroundHandler("/query"))
	router.POST("/query",
		middlewares.SetUserContext(global.Config.JWTSecret),
		graph.GrapgqlHandler(userService, authService, emailservice))

	port := fmt.Sprintf(":%s", global.Config.Port)
	router.Run(port)
}
