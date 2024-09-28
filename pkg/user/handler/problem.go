package handler

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	pb "github.com/shivaraj-shanthaiah/code_orbit_apigateway/pkg/user/userpb"
	"google.golang.org/grpc/metadata"
)

func UserGetAllProblemsHandler(c *gin.Context, client pb.UserServiceClient) {
	timeout := time.Second * 100
	ctx, cancel := context.WithTimeout(c, timeout)
	defer cancel()

	response, err := client.UserGetAllProblems(ctx, &pb.UserNoParam{})
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
		"Message": "All problems fetched successfully",
		"Data":    response,
	})
}

// func GetProblemWithTestCasesHandler(c *gin.Context, client pb.UserServiceClient) {
// 	timeout := time.Second * 100
// 	ctx, cancel := context.WithTimeout(c, timeout)
// 	defer cancel()

// 	problemIDParam := c.Param("id")
// 	problemID, err := strconv.Atoi(problemIDParam)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"Status":  http.StatusBadRequest,
// 			"Message": "Invalid problem ID",
// 			"Error":   err.Error(),
// 		})
// 		return
// 	}

// 	response, err := client.UserGetProblemWithTestCases(ctx, &pb.UserProblemId{
// 		ID: int32(problemID),
// 	})

// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"Status":  http.StatusInternalServerError,
// 			"Message": "Error fetching problem details",
// 			"Error":   err.Error(),
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"Status":  http.StatusOK,
// 		"Message": "Problem with test cases fetched successfully",
// 		"Data":    response,
// 	})
// }

func GetProblemWithTestCasesHandler(c *gin.Context, client pb.UserServiceClient) {
	timeout := time.Second * 100
	ctx, cancel := context.WithTimeout(c, timeout)
	defer cancel()

	userID := c.GetString("user_id")
	if userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"Status":  http.StatusUnauthorized,
			"Message": "User not authenticated",
		})
		return
	}

	log.Println("UserID in handler: ", userID)

	//gRPC metadata to pass the user ID
	md := metadata.New(map[string]string{
		"user_id": userID,
	})
	ctx = metadata.NewOutgoingContext(ctx, md)

	problemIDParam := c.Param("id")
	problemID, err := strconv.Atoi(problemIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status":  http.StatusBadRequest,
			"Message": "Invalid problem ID",
			"Error":   err.Error(),
		})
		return
	}

	response, err := client.UserGetProblemWithTestCases(ctx, &pb.UserProblemId{
		ID: int32(problemID),
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Status":  http.StatusInternalServerError,
			"Message": "Error fetching problem details",
			"Error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Status":  http.StatusOK,
		"Message": "Responded successfully",
		"Data":    response,
	})
}
