package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shivaraj-shanthaiah/code_orbit_apigateway/pkg/model"
	pb "github.com/shivaraj-shanthaiah/code_orbit_apigateway/pkg/user/userpb"
)

func UserSignupHandler(c *gin.Context, client pb.UserServiceClient) {
	timeout := time.Second * 100
	ctx, cancel := context.WithTimeout(c, timeout)
	defer cancel()

	var user model.User
	if err := c.BindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status":  http.StatusBadRequest,
			"Message": "error while binding json",
			"Error":   err.Error()})
		return
	}

	response, err := client.UserSignup(ctx, &pb.Signup{
		User_Name: user.UserName,
		Email:     user.Email,
		Password:  user.Password,
		Phone:     user.Phone,
	})

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status":  http.StatusBadRequest,
			"Message": "error in client response",
			"Data":    response,
			"Error":   err.Error()})
		return
	}

	if response.Status == pb.Response_ERROR {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status":  http.StatusBadRequest,
			"Message": response.Message,
			"Data":    response.Payload,
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"Status":  http.StatusAccepted,
		"Message": "Go to verification page and enter otp",
		"Data":    response,
	})
}

func VerificationHandler(c *gin.Context, client pb.UserServiceClient) {
	timeout := time.Second * 100
	ctx, cancel := context.WithTimeout(c, timeout)
	defer cancel()

	var verficationdetails model.OTP

	if err := c.BindJSON(&verficationdetails); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status":  http.StatusBadRequest,
			"Message": "error while binding jason",
			"Error":   err.Error()})
		return
	}

	response, err := client.VerifyUser(ctx, &pb.OTP{
		Email: verficationdetails.Email,
		OTP:   verficationdetails.OTP,
	})

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status":  http.StatusBadRequest,
			"Message": "error in client response",
			"Data":    response,
			"Error":   err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"Status":  http.StatusAccepted,
		"Message": "Verification succesfull",
		"Data":    response,
	})
}

func UserLoginHandler(c *gin.Context, client pb.UserServiceClient) {
	timeout := time.Second * 100
	ctx, cancel := context.WithTimeout(c, timeout)
	defer cancel()

	var user model.Login
	if err := c.BindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status":  http.StatusBadRequest,
			"Message": "error while binding json",
			"Error":   err.Error()})
		return
	}

	response, err := client.UserLogin(ctx, &pb.Login{
		Email:    user.Email,
		Password: user.Password,
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status":  http.StatusBadRequest,
			"Message": "error in client request",
			"Error":   err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"Status":  http.StatusAccepted,
		"Message": "Login Successful",
		"Data":    response,
	})
}
