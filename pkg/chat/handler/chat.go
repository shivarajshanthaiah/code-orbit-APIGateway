package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	pb "github.com/shivaraj-shanthaiah/code_orbit_apigateway/pkg/chat/chatpb"
	"github.com/shivaraj-shanthaiah/code_orbit_apigateway/pkg/model"
	userpb "github.com/shivaraj-shanthaiah/code_orbit_apigateway/pkg/user/userpb"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func HandleWebSocketConnection(c *gin.Context, client pb.ChatServiceClient, userClient userpb.UserServiceClient) {
	ctx := c.Request.Context()

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("Error upgrading to WebSocket:", err)
		return
	}
	defer conn.Close()

	for {
		select {
		case <-ctx.Done():
			log.Println("WebSocket connection closed")
			return
		default:
			_, msg, err := conn.ReadMessage()
			if err != nil {
				log.Println("Error reading message:", err)
				break
			}

			var message model.Message
			if err := json.Unmarshal(msg, &message); err != nil {
				log.Println("Invalid message format:", err)
				continue
			}

			// Check if user and receiver exist
			if _, err := userClient.ViewProfile(ctx, &userpb.ID{ID: message.SenderID}); err != nil {
				log.Println("Invalid UserID:", err)
				continue
			}
			if _, err := userClient.ViewProfile(ctx, &userpb.ID{ID: message.RecipientID}); err != nil {
				log.Println("Invalid ReceiverID:", err)
				continue
			}

			// Send the message to the gRPC chat service
			stream, err := client.Connect(ctx)
			if err != nil {
				log.Println("Error connecting to chat service:", err)
				continue
			}

			ch := &clientHandle{
				stream:     stream,
				userID:     message.SenderID,
				receiverID: message.RecipientID,
			}

			// Write received message back to WebSocket client
			if err := conn.WriteMessage(websocket.TextMessage, msg); err != nil {
				log.Println("Error writing message:", err)
				break
			}

			// Start goroutines to send and receive messages
			go ch.sentMessage(message.Content)
			go ch.receiveMessage(conn, message.SenderID, message.RecipientID)
		}
	}
}

type clientHandle struct {
	userID     string
	receiverID string
	stream     pb.ChatService_ConnectClient
}

func ChatScreen(c *gin.Context, client pb.ChatServiceClient) {
	ctx, cancel := context.WithTimeout(c, 100*time.Second)
	defer cancel()

	userID := c.Query("id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status":  http.StatusBadRequest,
			"Message": "User ID is required",
		})
		return
	}

	receiverID := c.Query("receiverId")
	if receiverID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status":  http.StatusBadRequest,
			"Message": "Receiver ID is required",
		})
		return
	}

	response, err := client.FetchHistory(ctx, &pb.ChatID{
		User_ID:     userID,
		Receiver_ID: receiverID,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Status":  http.StatusInternalServerError,
			"Message": "Failed to fetch chat history",
			"Error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Status":     http.StatusOK,
		"UserID":     userID,
		"ReceiverID": receiverID,
		"Chats":      response.Chats,
	})
}

func (c *clientHandle) sentMessage(msg string) {
	message := &pb.Message{
		User_ID:     string(c.userID),
		Receiver_ID: string(c.receiverID),
		Content:     msg,
	}

	err := c.stream.Send(message)
	if err != nil {
		log.Printf("Error while sending message to server :: %v", err)
	}

}

func (ch *clientHandle) receiveMessage(c *websocket.Conn, userID, receiverID string) {
	for {
		mssg, err := ch.stream.Recv()
		if err != nil {
			if websocket.IsCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway) {
				log.Println("Connection closed. Stopping message reception.")
				return
			}
			log.Printf("Error in receiving message from server: %v", err)
			return
		}

		if string(userID) == mssg.Receiver_ID && receiverID == string(mssg.User_ID) {
			model := &model.Message{
				SenderID:    string(mssg.User_ID),
				RecipientID: string(mssg.Receiver_ID),
				Content:     mssg.Content,
			}
			msg, err := json.Marshal(model)
			if err != nil {
				log.Println("Error decoding JSON:", err)
				continue
			}
			err = c.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				log.Println("Error writing message:", err)
				break
			}
		}
		fmt.Printf("%s : %s to %s\n", mssg.Content, mssg.User_ID, mssg.Receiver_ID)
	}
}
