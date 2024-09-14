package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Authorization(key string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")

		if tokenString == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"Status":  "failed",
				"Message": "Token not found in header",
				"Data":    "",
				"Error":   "null token",
			})
			ctx.Abort()
			return
		}

		tokenString = strings.TrimSpace(strings.Replace(tokenString, "Bearer ", "", 1))

		// Decode and parse the token
		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			// Validate the algorithm
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
			}
			return []byte(key), nil
		})

		// Handle parsing errors
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"Status":  "failed",
				"Message": "Token not valid",
				"Data":    "",
				"Error":   err.Error(),
			})
			ctx.Abort()
			return
		}

		// Validate the token claims
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"Status":  "failed",
				"Message": "Invalid token claims",
				"Data":    "",
				"Error":   "invalid claims",
			})
			ctx.Abort()
			return
		}

		// Extract "Email" from claims
		email, ok := claims["Email"].(string)
		if !ok || email == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"Status":  "failed",
				"Message": "Email not found in claims",
				"Data":    "",
				"Error":   "invalid email",
			})
			ctx.Abort()
			return
		}

		// Extract "UserID" from claims
		userId, ok := claims["UserID"].(string)
		if !ok || userId == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"Status":  "failed",
				"Message": "UserID not found in claims",
				"Data":    "",
				"Error":   "invalid userID",
			})
			ctx.Abort()
			return
		}

		// Set user details in context
		ctx.Set("email", email)
		ctx.Set("user_id", userId)
		ctx.Next()
	}
}

func AdminAuthorization(key, role string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")

		if tokenString == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"Status": "Failed",
				"Message": "Token not found in header",
				"Data":    "",
				"Error":   "null token"})
			ctx.Abort()
			return
		}

		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			return []byte(key), nil
		})

		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"Status": "Failed",
				"Message": "Token not valid",
				"Data":    "",
				"Error":   err.Error()})
			ctx.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			ctx.JSON(http.StatusUnauthorized, gin.H{"Status": "Failed",
				"Message": "Invalid token claims",
				"Data":    "",
				"Error":   ok})
			ctx.Abort()
			return
		}
		email, ok := claims["Email"].(string)
		if !ok {
			ctx.JSON(http.StatusUnauthorized, gin.H{"Status": "Failed",
				"Message": "Email not found in claims",
				"Data":    "",
				"Error":   ok})
			ctx.Abort()
			return
		}

		claimRole, ok := claims["Role"].(string)
		if !ok {
			ctx.JSON(http.StatusUnauthorized, gin.H{"Status": "Failed",
				"Message": "role not found in token",
				"Data":    role,
				"Error":   ok})
			ctx.Abort()
			return
		}

		if role != claimRole {
			ctx.JSON(http.StatusUnauthorized, gin.H{"Status": "Failed",
				"Message": "role not matching",
				"Data":    role,
				"Error":   ok})
			ctx.Abort()
			return
		}
		ctx.Set("email", email)
		ctx.Set("role", role)
		ctx.Next()
	}
}
