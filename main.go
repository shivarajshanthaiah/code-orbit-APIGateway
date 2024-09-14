package main

import (
	"log"

	"github.com/shivaraj-shanthaiah/code_orbit_apigateway/pkg/admin"
	"github.com/shivaraj-shanthaiah/code_orbit_apigateway/pkg/config"
	"github.com/shivaraj-shanthaiah/code_orbit_apigateway/pkg/server"
	"github.com/shivaraj-shanthaiah/code_orbit_apigateway/pkg/user"
)

func main() {
	cnfg, err := config.LoadConfig()
	if err != nil {
		log.Printf("Error while loading config file")
	}

	server := server.NewServer()
	server.R.LoadHTMLGlob("templates/*")
	user.NewUserRoute(server.R, *cnfg)
	admin.NewAdminRoute(server.R, *cnfg)
	server.StartServer(cnfg.APIPORT)
}
