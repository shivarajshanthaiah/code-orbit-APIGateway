package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	pb "github.com/shivaraj-shanthaiah/code_orbit_apigateway/pkg/user/userpb"
)

func UserGetProblemStatsHandler(c *gin.Context, client pb.UserServiceClient) {
	timeout := time.Second * 100
	ctx, cancel := context.WithTimeout(c, timeout)
	defer cancel()

	response, err := client.UserGetProblemStats(ctx, &pb.UProblemStatsRequest{})
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
		"Message": "Problem stats fetched successfully",
		"Data":    response,
	})
}

func UserGetLeaderboardHandler(c *gin.Context, adminClient pb.UserServiceClient) {
	timeout := time.Second * 100
	ctx, cancel := context.WithTimeout(c, timeout)
	defer cancel()

	response, err := adminClient.UserGetLeaderboardStats(ctx, &pb.ULeaderboardRequest{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "error in admin client response",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   response.Leaderboard,
	})
}
