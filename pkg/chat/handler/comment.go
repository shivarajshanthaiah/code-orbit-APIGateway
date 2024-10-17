package handler

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	pb "github.com/shivaraj-shanthaiah/code_orbit_apigateway/pkg/chat/chatpb"
	"github.com/shivaraj-shanthaiah/code_orbit_apigateway/pkg/model"
)

func AddComment(c *gin.Context, client pb.ChatServiceClient) {
	timeout := time.Second * 100
	ctx, cancel := context.WithTimeout(c, timeout)
	defer cancel()

	id, ok := c.Get("user_id")
	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status":  http.StatusBadRequest,
			"Message": "User ID not found in context",
			"Error":   "",
		})
		return
	}

	userID, ok := id.(string)
	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status":  http.StatusBadRequest,
			"Message": "Error while user id conversion",
			"Error":   "",
		})
		return
	}

	var comment model.Comment
	if err := c.BindJSON(&comment); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status":  http.StatusBadRequest,
			"Message": "Error binding JSON request body",
			"Error":   err.Error(),
		})
		return
	}
	fmt.Println(comment)

	response, err := client.AddComment(ctx, &pb.CommentRequest{
		ProblemId: uint32(comment.ProblemID),
		UserId:    userID,
		Content:   comment.Content,
	})

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"Status":  http.StatusInternalServerError,
			"Message": "error in client response",
			"Error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"Status":  http.StatusCreated,
		"Message": "comment added successfully",
		"Data":    response,
	})
}

func ReplyToComment(c *gin.Context, client pb.ChatServiceClient) {
	timeout := time.Second * 100
	ctx, cancel := context.WithTimeout(c, timeout)
	defer cancel()

	id, ok := c.Get("user_id")
	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status":  http.StatusBadRequest,
			"Message": "User ID not found in context",
			"Error":   "",
		})
		return
	}

	userID, ok := id.(string)
	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status":  http.StatusBadRequest,
			"Message": "Error while user id conversion",
			"Error":   "",
		})
		return
	}

	var comment model.Comment
	if err := c.BindJSON(&comment); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status":  http.StatusBadRequest,
			"Message": "Error binding JSON request body",
			"Error":   err.Error(),
		})
		return
	}

	response, err := client.ReplyToComment(ctx, &pb.ReplyRequest{
		CommentId: comment.ID,
		UserId:    userID,
		Content:   comment.Content,
	})

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"Status":  http.StatusInternalServerError,
			"Message": "Error in client response",
			"Error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"Status":  http.StatusCreated,
		"Message": "Reply added successfully",
		"Data":    response,
	})
}

func GetCommentsForProblem(c *gin.Context, client pb.ChatServiceClient) {
	timeout := time.Second * 100
	ctx, cancel := context.WithTimeout(c, timeout)
	defer cancel()

	problemIDParam := c.Param("id")
	problemID, err := strconv.Atoi(problemIDParam)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status":  http.StatusBadRequest,
			"Message": "Invalid problem ID",
			"Error":   err.Error(),
		})
		return
	}

	response, err := client.FetchComments(ctx, &pb.FetchCommentsRequest{
		ProblemId: uint32(problemID),
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"Status":  http.StatusInternalServerError,
			"Message": "Error fetching comments",
			"Error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Status":   http.StatusOK,
		"Comments": response.Comments,
	})
}

func GetUserComments(c *gin.Context, client pb.ChatServiceClient) {
	timeout := time.Second * 100
	ctx, cancel := context.WithTimeout(c, timeout)
	defer cancel()

	id, ok := c.Get("user_id")
	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status":  http.StatusBadRequest,
			"Message": "User ID not found in context",
			"Error":   "",
		})
		return
	}

	userID, ok := id.(string)
	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status":  http.StatusBadRequest,
			"Message": "Error while user id conversion",
			"Error":   "",
		})
		return
	}

	response, err := client.FetchUserComments(ctx, &pb.FetchUserCommentsRequest{
		UserId: userID,
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"Status":  http.StatusInternalServerError,
			"Message": "Error fetching user comments",
			"Error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Status":   http.StatusOK,
		"Comments": response.Comments,
	})
}
