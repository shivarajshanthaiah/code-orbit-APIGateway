package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shivaraj-shanthaiah/code_orbit_apigateway/pkg/model"
	pb "github.com/shivaraj-shanthaiah/code_orbit_apigateway/pkg/user/userpb"
)

func SubmitCodeHandler(c *gin.Context, client pb.UserServiceClient) {
	timeout := time.Second * 100
	ctx, cancel := context.WithTimeout(c, timeout)
	defer cancel()

	id, ok := c.Get("user_id")
	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status":  http.StatusBadRequest,
			"Message": "Error while fetching id from context",
			"Error":   ""})
		return
	}

	var submissionRequest model.Submission

	// var submissionRequest struct {
	// 	ProblemID int    `json:"problem_id"`
	// 	Language  string `json:"language"`
	// 	Code      string `json:"code"`
	// }

	if err := c.BindJSON(&submissionRequest); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status":  http.StatusBadRequest,
			"Message": "Invalid request",
			"Error":   err.Error(),
		})
		return
	}

	grpcReq := &pb.UserSubmissionRequest{
		UserId:    id.(string),
		ProblemId: int32(submissionRequest.ProblemID),
		Language:  submissionRequest.Language,
		Code:      submissionRequest.Code,
	}

	response, err := client.SubmitCode(ctx, grpcReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Status":  http.StatusInternalServerError,
			"Message": "Failed to submit code",
			"Error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Status":          http.StatusOK,
		"Message":         response.Message,
		"PassedTestCases": response.Passed,
		"FailedTestCases": response.Failed,
	})
}
