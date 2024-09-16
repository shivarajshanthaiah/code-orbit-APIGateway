package handler

import (
	"context"
	"net/http"
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
