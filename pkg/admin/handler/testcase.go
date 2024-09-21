package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	pb "github.com/shivaraj-shanthaiah/code_orbit_apigateway/pkg/admin/adminpb"
	"github.com/shivaraj-shanthaiah/code_orbit_apigateway/pkg/model"
)

func InsertTestCaseHandler(c *gin.Context, client pb.AdminServiceClient) {
	timeout := time.Second * 100
	ctx, cancel := context.WithTimeout(c, timeout)
	defer cancel()

	var requestBody struct {
		ProblemID int              `json:"problem_id"`
		TestCases []model.TestCase `json:"test_cases"`
	}
	if err := c.BindJSON(&requestBody); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status":  http.StatusBadRequest,
			"Message": "error while binding JSON",
			"Error":   err.Error(),
		})
		return
	}

	var grpcTestCases []*pb.AdTestCase
	for _, tc := range requestBody.TestCases {
		grpcTestCases = append(grpcTestCases, &pb.AdTestCase{
			Input:          tc.Input,
			ExpectedOutput: tc.ExpectedOutput,
		})
	}

	response, err := client.InsertTestCases(ctx, &pb.AdTestCaseRequest{
		ProblemId: int32(requestBody.ProblemID),
		TestCases: grpcTestCases,
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
		"Message": "testcase added successfully",
		"Data":    response,
	})
}

func UpdateTestCaseHandler(c *gin.Context, client pb.AdminServiceClient) {
	timeout := time.Second * 100
	ctx, cancel := context.WithTimeout(c, timeout)
	defer cancel()

	var requestBody struct {
		ProblemID  int              `json:"problem_id"`
		TestCases  []model.TestCase `json:"test_cases"`
		TestCaseID string           `json:"test_case_id"` // ID of the test case to update
	}
	if err := c.BindJSON(&requestBody); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status":  http.StatusBadRequest,
			"Message": "error while binding JSON",
			"Error":   err.Error(),
		})
		return
	}

	var grpcTestCases []*pb.AdTestCase
	for _, tc := range requestBody.TestCases {
		grpcTestCases = append(grpcTestCases, &pb.AdTestCase{
			Input:          tc.Input,
			ExpectedOutput: tc.ExpectedOutput,
		})
	}

	response, err := client.UpdateTestCases(ctx, &pb.AdUpdateTestCaseRequest{
		TestCaseId: requestBody.TestCaseID,
		ProblemId:  int32(requestBody.ProblemID),
		TestCases:  grpcTestCases,
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
		"Message": "testcase updated successfully",
		"Data":    response,
	})
}
