package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	pb "github.com/shivaraj-shanthaiah/code_orbit_apigateway/pkg/admin/adminpb"
)

func GetAllUserStatsHandler(c *gin.Context, client pb.AdminServiceClient) {
	timeout := time.Second * 100
	ctx, cancel := context.WithTimeout(c, timeout)
	defer cancel()

	response, err := client.AdminGetUserStats(ctx, &pb.AdUserStatsRequest{})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"Status":  http.StatusInternalServerError,
			"Message": "error in client response",
			"Error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"Status":  http.StatusAccepted,
		"Message": "All user stats fetched successfully",
		"Data":    response,
	})
}

func GetSubscriptionStatsHandler(c *gin.Context, client pb.AdminServiceClient) {
	timeout := time.Second * 100
	ctx, cancel := context.WithTimeout(c, timeout)
	defer cancel()

	response, err := client.AdminGetSubscriptionStats(ctx, &pb.AdSubscriptionStatsRequest{})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"Status":  http.StatusInternalServerError,
			"Message": "error in client response",
			"Error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"Status":  http.StatusAccepted,
		"Message": "Subscription stats fetched successfully",
		"Data":    response,
	})
}
