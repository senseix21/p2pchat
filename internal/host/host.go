package host

import (
	"crypto/rand"
	"fmt"
	"log"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/host"
)

func CreateHost(port int) host.Host {

	//Generate a secret key
	priv, _, err := crypto.GenerateKeyPairWithReader(crypto.RSA, 2048, rand.Reader)
	if err != nil {
		log.Fatalf("Error creating key: %v\n", err)
	}

	//Multiaddress for host
	addr := fmt.Sprintf("/ip4/0.0.0.0/tcp/%d", port)

	//Create the host
	h, err := libp2p.New(
		libp2p.Identity(priv),
		libp2p.ListenAddrStrings(addr))

	if err != nil {
		log.Fatalf("Error creating the host: %v\n", err)
	}

	return h

}

func DisplayHostInfo(h host.Host) {
	fmt.Printf("Host created with: %v\n", h.ID())

	for _, addr := range h.Addrs() {
		fmt.Printf("Listening on: %s/p2p/%s\n", addr.String(), h.ID())
	}
}
