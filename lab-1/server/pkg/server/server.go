package server

import (
	"bufio"
	"log"
	"net"
	"os"
	"strconv"
	"time"
	"trspo-lab1/server/pkg/config"
	"trspo-lab1/server/pkg/models"
)

func StartAndListen(settings *models.ConnectionSettings) error {
	log.Println("Starting " + settings.Protocol + " server on " + settings.Ip + ":" + settings.Port)
	listener, err := net.Listen(settings.Protocol, settings.Ip+":"+settings.Port)
	if err != nil {
		return err
	}
	defer listener.Close()

	acceptAndHandle(listener)

	return nil
}

func acceptAndHandle(listener net.Listener) error {
	delaySecs, err := strconv.Atoi(os.Getenv(config.HostDelay))
	if err != nil {
		return err
	}
	delayTime := time.Second * time.Duration(delaySecs)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error connecting:", err.Error())
			return err
		}
		log.Println("Client " + conn.RemoteAddr().String() + " connected.")

		go handleConnection(
			conn,
			time.Duration(delayTime),
		)
	}
}

func handleConnection(conn net.Conn, delay time.Duration) {
	buffer, err := bufio.NewReader(conn).ReadBytes('\n')
	if err != nil {
		log.Printf("Client %s left.", conn.RemoteAddr().String())
		conn.Close()
		return
	}

	time.Sleep(delay)
	log.Printf("Client %s message: %s", conn.RemoteAddr().String(), string(buffer[:len(buffer)-1]))
	conn.Write(buffer)

	handleConnection(conn, delay)
}
