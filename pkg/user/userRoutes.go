package user

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/shivaraj-shanthaiah/code_orbit_apigateway/middleware"
	"github.com/shivaraj-shanthaiah/code_orbit_apigateway/pkg/config"
	pb "github.com/shivaraj-shanthaiah/code_orbit_apigateway/pkg/user/userpb"
)

type User struct {
	cfg    *config.Config
	Client pb.UserServiceClient
}

func NewUserRoute(c *gin.Engine, cfg config.Config) {
	client, err := ClientDial(cfg)
	if err != nil {
		log.Fatalf("Error while connecting with grpc client :%v", err.Error())
	}

	userHandler := &User{
		cfg:    &cfg,
		Client: client,
	}

	apiVersion := c.Group("/api")
	user := apiVersion.Group("/user")
	{
		user.POST("/signup", userHandler.UserSignup)
		user.POST("/verify", userHandler.UserVerify)
		user.POST("/login", userHandler.UserLogin)
	}

	auth := user.Group("/auth")
	auth.Use(middleware.Authorization(cfg.SECRETKEY))
	{
		auth.GET("/view/profile", userHandler.ViewProfile)
		auth.PUT("/update/profile", userHandler.EditProfile)
		auth.PATCH("/change/password", userHandler.ChangePassword)

		auth.GET("/get/all/problems", userHandler.UGetAllProblems)
		auth.GET("/get/problem/:id", userHandler.UGetProblemByID)

		auth.POST("/submit/code", userHandler.SubmitCode)
		auth.GET("/get/stats", userHandler.GetUserStats)

		auth.GET("/get/all/plans", userHandler.GetAllPlans)
	}
}