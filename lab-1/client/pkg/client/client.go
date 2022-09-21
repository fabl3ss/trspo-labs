package client

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"trspo-lab1/client/pkg/config"
	"trspo-lab1/client/pkg/models"
)

func NewClient(settings *models.ConnectionSettings) {
	// Start the client and connect to the server.
	log.Printf("Connecting to %s server %s", settings.Protocol, settings.Ip+":"+settings.Port)
	host, err := net.Dial(settings.Protocol, settings.Ip+":"+settings.Port)
	if err != nil {
		fmt.Println("Error connecting:", err.Error())
		os.Exit(1)
	}

	messagesAmount, err := strconv.Atoi(os.Getenv(config.MessagesAmount))
	if err != nil {
		panic("unable get messages amount")
	}
	sendMessages(host, messagesAmount)
}

func sendMessages(host net.Conn, amount int) {
	for i := 0; i < amount; i++ {
		host.Write([]byte(strconv.Itoa(i) + "\n"))
		message, _ := bufio.NewReader(host).ReadString('\n')
		log.Print("Server relay: " + message)
	}
}
