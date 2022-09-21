package main

import (
	"os"
	"strconv"
	"sync"
	"trspo-lab1/client/pkg/client"
	"trspo-lab1/client/pkg/config"
	"trspo-lab1/client/pkg/models"
)

func main() {
	clientsAmount, err := strconv.Atoi(os.Getenv(config.ClientsAmount))
	if err != nil {
		panic("unable get clients amount")
	}

	var wg sync.WaitGroup
	connectionSettings := &models.ConnectionSettings{
		Ip:       os.Getenv(config.HostIp),
		Port:     os.Getenv(config.HostPort),
		Protocol: os.Getenv(config.HostProtocol),
	}

	for i := 0; i < clientsAmount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			client.NewClient(connectionSettings)
		}()
	}
	wg.Wait()
}
