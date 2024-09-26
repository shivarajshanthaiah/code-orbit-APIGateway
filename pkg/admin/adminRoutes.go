package admin

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/shivaraj-shanthaiah/code_orbit_apigateway/middleware"
	pb "github.com/shivaraj-shanthaiah/code_orbit_apigateway/pkg/admin/adminpb"
	"github.com/shivaraj-shanthaiah/code_orbit_apigateway/pkg/config"
)

type Admin struct {
	cfg    *config.Config
	Client pb.AdminServiceClient
}

func NewAdminRoute(c *gin.Engine, cfg config.Config) {
	client, err := ClientDial(cfg)
	if err != nil {
		log.Fatalf("error not conntected with gRPC lient : %v", err.Error())
	}

	adminHandler := &Admin{
		cfg:    &cfg,
		Client: client,
	}

	api := c.Group("/api")

	admin := api.Group("/admin")
	{
		admin.POST("/login", adminHandler.AdminLogin)
	}

	auth := admin.Group("/auth")
	auth.Use(middleware.AdminAuthorization(cfg.SECRETKEY, "admin"))
	{
		auth.PATCH("/user/block/:id", adminHandler.BlockUser)
		auth.PATCH("/user/unblock/:id", adminHandler.UnblockUser)
		auth.GET("/user/list", adminHandler.GetAllUsers)
		auth.GET("/user/:id", adminHandler.GetUserByID)

		auth.POST("/add/problem", adminHandler.InsertProblem)
		auth.GET("/get/all/problem", adminHandler.GetAllProblems)
		auth.PUT("/edit/problem/:id", adminHandler.EditProblem)
		auth.PATCH("/upgrade/problem/:id", adminHandler.UpgradeProblem)
		
		auth.POST("/insert/testcases", adminHandler.InsertTestCases)
		auth.PUT("/update/testcases", adminHandler.UpdateTestCases)
		auth.GET("/get/problem/:id", adminHandler.GetProblemWithTestCases)
		
	}
}
