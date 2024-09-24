package handler

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	pb "github.com/shivaraj-shanthaiah/code_orbit_apigateway/pkg/admin/adminpb"
	"github.com/shivaraj-shanthaiah/code_orbit_apigateway/pkg/model"
)

func InsertProblemHanlder(c *gin.Context, client pb.AdminServiceClient) {
	timeout := time.Second * 100
	ctx, cancel := context.WithTimeout(c, timeout)
	defer cancel()

	var problem model.Problem
	if err := c.BindJSON(&problem); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status":  http.StatusBadRequest,
			"Message": "error while binding JSON",
			"Error":   err.Error(),
		})
		return
	}

	response, err := client.InsertProblem(ctx, &pb.Problem{
		Title:       problem.Title,
		Discription: problem.Description,
		Difficulty:  problem.Difficulty,
		Tags:        problem.Tags,
		IsPremium:   problem.IsPremium,
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
		"Message": "problem added successfully",
		"Data":    response,
	})
}

func AdminGetAllProblemsHandler(c *gin.Context, client pb.AdminServiceClient) {
	timeout := time.Second * 100
	ctx, cancel := context.WithTimeout(c, timeout)
	defer cancel()

	response, err := client.AdminGetAllProblems(ctx, &pb.AdNoParam{})
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
		"Message": "All problems fetched successfully",
		"Data":    response,
	})
}

// func EditProblemHandler(c *gin.Context, client pb.AdminServiceClient) {
// 	timeout := time.Second * 100
// 	ctx, cancel := context.WithTimeout(c, timeout)
// 	defer cancel()

// 	id, ok := c.Get("id")
// 	if !ok {
// 		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
// 			"Status":  http.StatusBadRequest,
// 			"Message": "error while fetching problem id from context",
// 			"Error":   ""})
// 		return
// 	}

// 	problemID, ok := id.(uint)
// 	if !ok {
// 		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Status": http.StatusBadRequest,
// 			"Message": "error while user id converting",
// 			"Error":   ""})
// 		return
// 	}

// 	var problem model.Problem
// 	if err := c.BindJSON(&problem); err != nil {
// 		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
// 			"Status":  http.StatusBadRequest,
// 			"Message": "error while binding json",
// 			"Error":   err.Error()})
// 		return
// 	}

// 	response, err := client.AdminEditProblem(ctx, &pb.Problem{
// 		ID:          uint32(problemID),
// 		Title:       problem.Title,
// 		Discription: problem.Description,
// 		Difficulty:  problem.Difficulty,
// 		Tags:        problem.Tags,
// 		IsPremium:   problem.IsPremium,
// 	})

// 	if err != nil {
// 		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Status": http.StatusBadRequest,
// 			"Message": "error in client response",
// 			"Data":    response,
// 			"Error":   err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusAccepted, gin.H{
// 		"Status":  http.StatusAccepted,
// 		"Message": "problem edited successfully",
// 		"Data":    response,
// 	})
// }

func EditProblemHandler(c *gin.Context, client pb.AdminServiceClient) {
	timeout := time.Second * 100
	ctx, cancel := context.WithTimeout(c, timeout)
	defer cancel()

	// Retrieve the problem ID from the URL parameter
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

	// Bind JSON body to the problem struct
	var problem model.Problem
	if err := c.BindJSON(&problem); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status":  http.StatusBadRequest,
			"Message": "Error while binding JSON",
			"Error":   err.Error(),
		})
		return
	}

	// Call the gRPC client to edit the problem
	response, err := client.AdminEditProblem(ctx, &pb.Problem{
		ID:          uint32(problemID),
		Title:       problem.Title,
		Discription: problem.Description,
		Difficulty:  problem.Difficulty,
		Tags:        problem.Tags,
		IsPremium:   problem.IsPremium,
	})

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status":  http.StatusBadRequest,
			"Message": "Error in client response",
			"Error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"Status":  http.StatusAccepted,
		"Message": "Problem edited successfully",
		"Data":    response,
	})
}

func AdminUpgradeProblemHandler(c *gin.Context, client pb.AdminServiceClient) {
	timeout := time.Second * 100
	ctx, cancel := context.WithTimeout(c, timeout)
	defer cancel()

	problemIdStr := c.Param("id")
	problemId, err := strconv.Atoi(problemIdStr)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Status": http.StatusBadRequest,
			"Message": "error while converting userID to int",
			"Error":   err.Error()})
		return
	}

	response, err := client.AdminUpgradeProbem(ctx, &pb.AdProblemId{
		ID: uint32(problemId),
	})
	
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status":  http.StatusBadRequest,
			"Message": "error in client response",
			"Error":   err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"Status":  http.StatusAccepted,
		"Message": "problem upgraded successfully",
		"Data":    response,
	})
}
