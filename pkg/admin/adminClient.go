package admin

import (
	"log"

	pb "github.com/shivaraj-shanthaiah/code_orbit_apigateway/pkg/admin/adminpb"
	"github.com/shivaraj-shanthaiah/code_orbit_apigateway/pkg/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ClientDial(cfg config.Config) (pb.AdminServiceClient, error) {
	grpc, err := grpc.Dial(":"+cfg.ADMINPORT, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("Error dialing to grpc to client : %s", err.Error())
		return nil, err
	}
	log.Printf("Successfully connected to admin client at port : %s", cfg.ADMINPORT)
	return pb.NewAdminServiceClient(grpc), nil
}