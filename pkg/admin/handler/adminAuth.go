package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	pb "github.com/shivaraj-shanthaiah/code_orbit_apigateway/pkg/admin/adminpb"
	"github.com/shivaraj-shanthaiah/code_orbit_apigateway/pkg/model"
)

func AdminLoginHandler(c *gin.Context, client pb.AdminServiceClient) {
	timeout := time.Second * 100
	ctx, cancel := context.WithTimeout(c, timeout)
	defer cancel()

	var admin model.Login
	if err := c.BindJSON(&admin); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status":  http.StatusBadRequest,
			"Message": "Error while binding JSON",
			"Error":   err.Error(),
		})
		return
	}

	response, err := client.AdminLoginRequest(ctx, &pb.AdminLogin{
		Email:    admin.Email,
		Password: admin.Password,
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
		"Message": "Login successful",
		"Data":    response,
	})
}