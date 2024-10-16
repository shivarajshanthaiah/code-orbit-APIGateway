package chat

import (
	"log"

	"github.com/gin-gonic/gin"
	pb "github.com/shivaraj-shanthaiah/code_orbit_apigateway/pkg/chat/chatpb"
	"github.com/shivaraj-shanthaiah/code_orbit_apigateway/pkg/config"
	"github.com/shivaraj-shanthaiah/code_orbit_apigateway/pkg/user"
	userpb "github.com/shivaraj-shanthaiah/code_orbit_apigateway/pkg/user/userpb"
)

type Chat struct {
	cfg        *config.Config
	userClient userpb.UserServiceClient
	client     pb.ChatServiceClient
}

func NewChatRoutes(c *gin.Engine, cfg config.Config) {
	client, err := ClientDial(cfg)
	if err != nil {
		log.Fatalf("error not connected with grpc client : %v", err.Error())
	}

	userClient, err := user.ClientDial(cfg)
	if err != nil {
		log.Fatalf("error not connected with grpc user client : %v", err.Error())
	}

	chatHandler := &Chat{
		cfg:        &cfg,
		client:     client,
		userClient: userClient,
	}

	apiVersion := c.Group("/api/v1")

	user := apiVersion.Group("/user")
	{
		user.GET("/chat", chatHandler.Chat)
	}
	c.GET("/chat", chatHandler.ChatScreen)
}
