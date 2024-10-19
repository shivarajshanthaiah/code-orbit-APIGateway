package chat

import (
	"github.com/gin-gonic/gin"
	"github.com/shivaraj-shanthaiah/code_orbit_apigateway/pkg/chat/handler"
)

func (c *Chat) Chat(ctx *gin.Context) {
	handler.HandleWebSocketConnection(ctx, c.client, c.userClient)
}

func (c *Chat) AddComment(ctx *gin.Context) {
	handler.AddComment(ctx, c.client)
}

func (c *Chat) ReplyToComment(ctx *gin.Context) {
	handler.ReplyToComment(ctx, c.client)
}

func (c *Chat) GetComments(ctx *gin.Context) {
	handler.GetCommentsForProblem(ctx, c.client)
}

func (c *Chat) GetUserComments(ctx *gin.Context) {
	handler.GetUserComments(ctx, c.client)
}
