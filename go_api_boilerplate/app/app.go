package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/yhagio/go_api_boilerplate/common/hmachash"
	"github.com/yhagio/go_api_boilerplate/common/randomstring"
	"github.com/yhagio/go_api_boilerplate/configs"
	pwdDomain "github.com/yhagio/go_api_boilerplate/domain/passwordreset"
	"github.com/yhagio/go_api_boilerplate/domain/user"

	"github.com/yhagio/go_api_boilerplate/gql"
	"github.com/yhagio/go_api_boilerplate/infra/mailgunclient"
	"github.com/yhagio/go_api_boilerplate/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/yhagio/go_api_boilerplate/docs" // docs is generated by Swag CLI

	"github.com/yhagio/go_api_boilerplate/controllers"
	"github.com/yhagio/go_api_boilerplate/repositories/passwordreset"
	"github.com/yhagio/go_api_boilerplate/repositories/userrepo"
	"github.com/yhagio/go_api_boilerplate/services/authservice"
	"github.com/yhagio/go_api_boilerplate/services/emailservice"
	"github.com/yhagio/go_api_boilerplate/services/userservice"

	_ "github.com/lib/pq" // For Postgres setup
)

var (
	router = gin.Default()
)

// @title Go API Boilerplate Swagger
// @version 1.0
// @description This is Go API Boilerplate
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name MIT
// @license.url https://github.com/yhagio/github.com/yhagio/go_api_boilerplate/blob/master/LICENSE

// @host localhost:3000
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func Run() {
	/*
		====== Swagger setup ============
		(http://localhost:3000/swagger/index.html)
	*/
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	/*
		====== Setup configs ============
	*/
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	config := configs.GetConfig()

	// Connects to PostgresDB
	db, err := gorm.Open(
		config.Postgres.Dialect(),
		config.Postgres.GetPostgresConnectionInfo(),
	)
	if err != nil {
		panic(err)
	}

	// Migration
	// db.DropTableIfExists(&user.User{})
	db.AutoMigrate(&user.User{}, &pwdDomain.PasswordReset{})
	defer db.Close()

	/*
		=======  Setup Common ============
	*/
	rds := randomstring.NewRandomString()
	hm := hmachash.NewHMAC(config.HMACKey)

	/*
		====== Setup infra ==============
	*/
	emailClient := mailgunclient.NewMailgunClient(config)

	/*
		====== Setup repositories =======
	*/
	userRepo := userrepo.NewUserRepo(db)
	pwdRepo := passwordreset.NewPasswordResetRepo(db)

	/*
		====== Setup services ===========
	*/
	userService := userservice.NewUserService(userRepo, pwdRepo, rds, hm, config.Pepper)
	authService := authservice.NewAuthService(config.JWTSecret)
	emailservice := emailservice.NewEmailService(emailClient)

	/*
		====== Setup controllers ========
	*/
	userCtl := controllers.NewUserController(userService, authService, emailservice)

	/*
		====== Setup middlewares ========
	*/
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	/*
		====== Setup routes =============
	*/
	router.GET("/ping", func(c *gin.Context) { c.String(http.StatusOK, "pong") })

	router.GET("/graphql", gql.PlaygroundHandler("/query"))
	router.POST("/query",
		middlewares.SetUserContext(config.JWTSecret),
		gql.GraphqlHandler(userService, authService, emailservice))

	api := router.Group("/api")

	api.POST("/register", userCtl.Register)
	api.POST("/login", userCtl.Login)
	api.POST("/forgot_password", userCtl.ForgotPassword)
	api.POST("/update_password", userCtl.ResetPassword)

	user := api.Group("/users")

	user.GET("/:id", userCtl.GetByID)

	account := api.Group("/account")
	account.Use(middlewares.RequireLoggedIn(config.JWTSecret))
	{
		account.GET("/profile", userCtl.GetProfile)
		account.PUT("/profile", userCtl.Update)
	}

	// Run
	// port := fmt.Sprintf(":%s", viper.Get("APP_PORT"))
	port := fmt.Sprintf(":%s", config.Port)
	router.Run(port)
}
