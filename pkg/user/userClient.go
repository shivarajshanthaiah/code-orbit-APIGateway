package user

import (
	"fmt"
	"log"

	"github.com/shivaraj-shanthaiah/code_orbit_apigateway/pkg/config"
	pb "github.com/shivaraj-shanthaiah/code_orbit_apigateway/pkg/user/userpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ClientDial(cfg config.Config) (pb.UserServiceClient, error) {
	// Combine host and port to create the full address
	grpcAddr := fmt.Sprintf("localhost:%s", cfg.USERPORT)

	grpcConn, err := grpc.Dial(grpcAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("Error dialing to grpc client: %s", err.Error())
		return nil, err
	}

	log.Printf("Successfully connected to user client at address : %s", grpcAddr)
	return pb.NewUserServiceClient(grpcConn), nil
}
