package main

import (
	"flag"
	"fmt"
	"log"
	"p2pchat/internal/host"
	"p2pchat/peer"
)

func main() {
	//ctx = context.Background()

	//Parse port form command line arguments
	port := flag.Int("port", 9000, "Port to listen on")
	target := flag.String("target", "", "Target peer multi addrs")
	flag.Parse()

	//Create libp2p host
	h := host.CreateHost(*port)

	//Display host info
	host.DisplayHostInfo(h)

	//Handele the incoming streams
	peer.HandleIncomingStream(h)

	//if target is specified, connect and send message
	if *target != "" {
		err := peer.ConnectToPeer(h, *target)

		if err != nil {
			fmt.Errorf("Failed to connect: %v\n", err)
		}

		err = peer.StartChat(h, *target)
		if err != nil {
			log.Fatalf("Chat error: %v", err)
		}
	}
	select {}

}
