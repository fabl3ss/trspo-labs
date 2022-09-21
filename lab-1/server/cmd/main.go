package main

import (
	"os"
	"trspo-lab1/server/pkg/config"
	"trspo-lab1/server/pkg/models"
	"trspo-lab1/server/pkg/server"
)

func main() {
	connSettings := &models.ConnectionSettings{
		Ip:       os.Getenv(config.HostIp),
		Port:     os.Getenv(config.HostPort),
		Protocol: os.Getenv(config.HostProtocol),
	}

	err := server.StartAndListen(connSettings)
	if err != nil {
		panic("unable to start and listen")
	}
}
