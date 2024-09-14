package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	pb "github.com/shivaraj-shanthaiah/code_orbit_apigateway/pkg/admin/adminpb"
)

func BlockUserHandler(c *gin.Context, client pb.AdminServiceClient) {
	timeout := time.Second * 100
	ctx, cancel := context.WithTimeout(c, timeout)
	defer cancel()

	userId := c.Param("id")

	response, err := client.AdminBlockUser(ctx, &pb.UserID{
		ID: string(userId),
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status":  http.StatusBadRequest,
			"Message": "error in client response",
			"Error":   err.Error(),
		})
		return
	}
	
	c.JSON(http.StatusAccepted, gin.H{
		"Status":  http.StatusAccepted,
		"Message": "User blocked succesfully",
		"Data":    response,
	})
}


func UnBlockUserHandler(c *gin.Context, client pb.AdminServiceClient) {
	timeout := time.Second * 100
	ctx, cancel := context.WithTimeout(c, timeout)
	defer cancel()

	userId := c.Param("id")

	response, err := client.AdminUnBlockUser(ctx, &pb.UserID{
		ID: string(userId),
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status":  http.StatusBadRequest,
			"Message": "error in client response",
			"Error":   err.Error(),
		})
		return
	}
	
	c.JSON(http.StatusAccepted, gin.H{
		"Status":  http.StatusAccepted,
		"Message": "User unblocked succesfully",
		"Data":    response,
	})
}