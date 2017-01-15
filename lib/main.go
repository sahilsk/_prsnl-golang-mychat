package lib

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

const port = "8080"

//RunHost is a function that takes listener ip
func RunHost(ip string) {
	ipAndPort := ip + ":" + port
	listener, listenErr := net.Listen("tcp", ipAndPort)
	if listenErr != nil {
		log.Fatal("Error", listenErr)
	}
	fmt.Println("listening on :" + ipAndPort)
	conn, acceptErr := listener.Accept()
	if acceptErr != nil {
		log.Fatal("Error:", acceptErr)
	}
	fmt.Println("New connection accepted")

	for {
		handleHost(conn)
	}

}

func handleHost(conn net.Conn) {
	reader := bufio.NewReader(conn)
	message, readErr := reader.ReadString('\n')
	if readErr != nil {
		log.Fatal("Error:", readErr)
	}
	fmt.Println("message received: ", message)
	fmt.Print("send message:")
	replyReader := bufio.NewReader(os.Stdin)
	replyMessage, replyErr := replyReader.ReadString('\n')
	if replyErr != nil {
		log.Fatal("Error", replyErr)
	}
	fmt.Fprintln(conn, replyMessage)
}

func RunGuest(ip string) {
	ipAndPort := ip + ":" + port
	conn, dialErr := net.Dial("tcp", ipAndPort)
	if dialErr != nil {
		log.Fatal("Error:", dialErr)
	}

	for {
		handleGuest(conn)
	}

}

func handleGuest(conn net.Conn) {
	fmt.Print("Send message: ")
	reader := bufio.NewReader(os.Stdin)
	message, readErr := reader.ReadString('\n')
	if readErr != nil {
		log.Fatal("Error", readErr)
	}
	_, err := fmt.Fprintln(conn, message)
	if err != nil {
		log.Fatal("Error: ", err)
	}
	replyReader := bufio.NewReader(conn)
	replyMessage, replyErr := replyReader.ReadString('\n')
	if replyErr != nil {
		log.Fatal("Error", replyErr)
	}
	fmt.Println("message received: " + replyMessage)
}
