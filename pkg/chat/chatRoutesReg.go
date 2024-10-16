package chat

import (
	"github.com/gin-gonic/gin"
	"github.com/shivaraj-shanthaiah/code_orbit_apigateway/pkg/chat/handler"
)

func (c *Chat) Chat(ctx *gin.Context) {
	handler.HandleWebSocketConnection(ctx, c.client, c.userClient)
}

func (c *Chat) ChatScreen(ctx *gin.Context) {
	handler.ChatScreen(ctx, c.client)
}
