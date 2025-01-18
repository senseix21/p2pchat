package peer

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/peerstore"

	ma "github.com/multiformats/go-multiaddr"
)

func ConnectToPeer(h host.Host, targetAddr string) error {
	maAddr, err := ma.NewMultiaddr(targetAddr)
	if err != nil {
		fmt.Errorf("Invalid target address: %w\n", err)
	}

	peerInfo, err := peer.AddrInfoFromP2pAddr(maAddr)
	if err != nil {
		fmt.Errorf("Error parsing the addrs: %w\n", err)
	}

	//add peer to peerstore
	h.Peerstore().AddAddrs(peerInfo.ID, peerInfo.Addrs, peerstore.PermanentAddrTTL)
	log.Printf("Connecting to peer: %s\n", peerInfo.ID)

	//Connect to the peer
	err = h.Connect(context.Background(), *peerInfo)

	if err != nil {
		return fmt.Errorf("failed to connect to peer: %w", err)
	}

	fmt.Println("Successfully connected to the peer!")

	return nil
}

// StartChat initiates a chat session with a connected peer
func StartChat(h host.Host, targetAddr string) error {
	maddr, err := ma.NewMultiaddr(targetAddr)
	if err != nil {
		return fmt.Errorf("invalid multiaddress: %w", err)
	}

	peerInfo, err := peer.AddrInfoFromP2pAddr(maddr)
	if err != nil {
		return fmt.Errorf("failed to parse peer address: %w", err)
	}

	stream, err := h.NewStream(context.Background(), peerInfo.ID, "/chat/1.0.0")
	if err != nil {
		return fmt.Errorf("failed to open stream: %w", err)
	}
	defer stream.Close()

	log.Println("Chat session started. Type your messages below:")

	// Goroutine to listen for incoming messages
	go func() {
		buf := bufio.NewReader(stream)
		for {
			msg, err := buf.ReadString('\n')
			if err != nil {
				if err == io.EOF {
					log.Println("Chat ended by the peer.")
					return
				}
				log.Printf("Error reading message: %v\n", err)
				continue
			}
			log.Printf("Peer: %s", msg)
		}
	}()

	// Read input from the user and send it
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("You: ")
		msg, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("Error reading input: %v\n", err)
			continue
		}

		_, err = stream.Write([]byte(msg))
		if err != nil {
			log.Printf("Error sending message: %v\n", err)
			return err
		}
	}
}

// HandleIncomingStream handles incoming chat streams
func HandleIncomingStream(h host.Host) {
	h.SetStreamHandler("/chat/1.0.0", func(s network.Stream) {
		log.Println("Incoming chat request")

		defer s.Close()

		// Goroutine to listen for messages
		go func() {
			buf := bufio.NewReader(s)
			for {
				msg, err := buf.ReadString('\n')
				if err != nil {
					if err == io.EOF {
						log.Println("Chat ended by the peer.")
						return
					}
					log.Printf("Error reading message: %v\n", err)
					continue
				}
				log.Printf("Peer: %s", msg)
			}
		}()

		// Read input from the user and send it
		reader := bufio.NewReader(os.Stdin)
		for {
			fmt.Print("You: ")
			msg, err := reader.ReadString('\n')
			if err != nil {
				log.Printf("Error reading input: %v\n", err)
				continue
			}

			_, err = s.Write([]byte(msg))
			if err != nil {
				log.Printf("Error sending message: %v\n", err)
				return
			}
		}
	})
}
