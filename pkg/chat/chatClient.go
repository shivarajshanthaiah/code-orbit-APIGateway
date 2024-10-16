package chat

import (
	"log"

	pb "github.com/shivaraj-shanthaiah/code_orbit_apigateway/pkg/chat/chatpb"
	"github.com/shivaraj-shanthaiah/code_orbit_apigateway/pkg/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ClientDial(cfg config.Config) (pb.ChatServiceClient, error) {
	grpc, err := grpc.Dial(":"+cfg.CHATPORT, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("Error dialing to grpc to client : %s", err.Error())
		return nil, err
	}
	log.Printf("Successfully connected to chat client at port : %s", cfg.CHATPORT)
	return pb.NewChatServiceClient(grpc), nil
}
