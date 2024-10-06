package handler

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	pb "github.com/shivaraj-shanthaiah/code_orbit_apigateway/pkg/user/userpb"
	"google.golang.org/grpc/metadata"
)

func GenerateInvoiceHandler(c *gin.Context, client pb.UserServiceClient) {
	timeout := time.Second * 100
	ctx, cancel := context.WithTimeout(c, timeout)
	defer cancel()

	userID := c.GetString("user_id")
	userEmail := c.GetString("email")

	// Add user_id and email to the gRPC metadata
	md := metadata.New(map[string]string{
		"user_id": userID,
		"email":   userEmail,
	})
	ctx = metadata.NewOutgoingContext(ctx, md)

	log.Println("User ID:", userID, "User Email:", userEmail)

	var req struct {
		SubscriptionID uint32 `json:"subscription_id"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := client.GenerateInvoice(ctx, &pb.InvoiceRequest{
		UserId:        userID,
		UserEmail:     userEmail,
		SubsriptionId: req.SubscriptionID,
	})
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
		"Message": "Invoice generated successfully.. Please proceed to payment",
		"Data":    response,
	})
}
