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

func AddSubscriptionHandler(c *gin.Context, client pb.AdminServiceClient) {
	timeout := time.Second * 100
	ctx, cancel := context.WithTimeout(c, timeout)
	defer cancel()

	var plan model.Subscription
	if err := c.BindJSON(&plan); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status":  http.StatusBadRequest,
			"Message": "error while binding JSON",
			"Error":   err.Error(),
		})
		return
	}

	response, err := client.AddSubscriptionPlan(ctx, &pb.AdSubscription{
		Plan:       plan.Plan,
		Duration:   plan.Duration,
		Price:      plan.Price,
		Gst:        plan.GST,
		TotalPrice: plan.TotalPrice,
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
		"Message": "plan added successfully",
		"Data":    response,
	})

}

func UpdatePlanHandler(c *gin.Context, client pb.AdminServiceClient) {
	timeout := time.Second * 100
	ctx, cancel := context.WithTimeout(c, timeout)
	defer cancel()

	planIDParam := c.Param("id")
	planID, err := strconv.Atoi(planIDParam)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status":  http.StatusBadRequest,
			"Message": "Invalid plan ID",
			"Error":   err.Error(),
		})
		return
	}

	var plan model.Subscription
	if err := c.BindJSON(&plan); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status":  http.StatusBadRequest,
			"Message": "Error while binding JSON",
			"Error":   err.Error(),
		})
		return
	}

	response, err := client.AdminUpdatePlan(ctx, &pb.AdSubscription{
		ID:         uint32(planID),
		Plan:       plan.Plan,
		Duration:   plan.Duration,
		Price:      plan.Price,
		Gst:        plan.GST,
		TotalPrice: plan.TotalPrice,
	})

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"Status":  http.StatusInternalServerError,
			"Message": "Error in client response",
			"Error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"Status":  http.StatusAccepted,
		"Message": "Plan updated successfully",
		"Data":    response,
	})
}

func AdminGetAllPlansHandler(c *gin.Context, client pb.AdminServiceClient) {
	timeout := time.Second * 100
	ctx, cancel := context.WithTimeout(c, timeout)
	defer cancel()

	response, err := client.GetAllPlans(ctx, &pb.AdNoParam{})
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
		"Message": "All plans fetched successfully",
		"Data":    response,
	})
}
