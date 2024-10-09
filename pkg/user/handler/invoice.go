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

func MakePaymentHandler(c *gin.Context, client pb.UserServiceClient) {
	timeout := time.Second * 100
	ctx, cancel := context.WithTimeout(c, timeout)
	defer cancel()

	invoice_id := c.Query("invoice_id")
	if invoice_id == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status": http.StatusBadRequest,
			"Error":  "Invoice ID is empty",
		})
		return
	}

	InvoiceID, err := strconv.Atoi(invoice_id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status":  http.StatusBadRequest,
			"Message": "Error while converting invoice id to integer",
			"Error":   err.Error(),
		})
		return
	}

	// Call the gRPC service to create a payment in Razorpay
	response, err := client.MakePayment(ctx, &pb.PaymentRequest{
		InvoiceId: uint32(InvoiceID),
		// UserId:   userID,
	})

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"Status": http.StatusInternalServerError,
			"Error":  err.Error(),
		})
		return
	}

	// Render the payment page with the required details
	c.HTML(http.StatusCreated, "app.html", gin.H{
		"userID":    response.User_ID,
		"invoiceID": response.Invoice_ID,
		"plan":      response.Plan,
		"total":     response.Amount,
		"orderID":   response.Order_ID,
	})
}

func PaymentSuccessHandler(c *gin.Context, client pb.UserServiceClient) {
	timeout := time.Second * 1000
	ctx, cancel := context.WithTimeout(c, timeout)
	defer cancel()

	paymentAmount := c.Query("total")
	signature := c.Query("signature")
	invoiceID := c.Query("invoice_id")
	orderID := c.Query("order_id")
	paymentID := c.DefaultQuery("payment_id", "")

	if invoiceID == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status": http.StatusBadRequest,
			"Error":  "invoice reference is empty",
		})
		return
	}

	response, err := client.PaymentSuccess(ctx, &pb.ConfirmRequest{
		Payment_ID: paymentID,
		Invoice_ID: invoiceID,
		Order_ID:   orderID,
		Signature:  signature,
		Amount:     paymentAmount,
	})

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"Status": http.StatusInternalServerError,
			"Error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"status":  http.StatusAccepted,
		"message": "Payment Confirmed",
		"data":    response,
	})
}

func PaymentSuccessPage(c *gin.Context, client pb.UserServiceClient) {
	c.HTML(http.StatusOK, "success.html", gin.H{
		"paymentID": c.Query("booking_reference"),
	})
}
