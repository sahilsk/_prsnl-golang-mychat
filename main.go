package main

import (
	"flag"
	"fmt"
	"mychat/lib"
	"os"
)

func main() {
	var isHost bool
	flag.BoolVar(&isHost, "listen", false, "listen to the specified ip address")
	flag.Parse()

	if isHost {
		//main.go -listen <ip>
		connIP := os.Args[2]
		fmt.Println("Is host")
		lib.RunHost(connIP)
	} else {
		// main.go <ip>
		connIP := os.Args[1]
		fmt.Println("is guest")
		lib.RunGuest(connIP)
	}

}
