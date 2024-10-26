package user

import (
	"log"

	"github.com/shivaraj-shanthaiah/code_orbit_apigateway/pkg/config"
	pb "github.com/shivaraj-shanthaiah/code_orbit_apigateway/pkg/user/userpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ClientDial(cfg config.Config) (pb.UserServiceClient, error) {
	// Combine host and port to create the full address
	address := "user-service:" + cfg.USERPORT

	grpcConn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("Error dialing to grpc client: %s", err.Error())
		return nil, err
	}

	log.Printf("Successfully connected to user client at address : %s", address)
	return pb.NewUserServiceClient(grpcConn), nil
}
