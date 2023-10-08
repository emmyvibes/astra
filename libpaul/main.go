package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	var privateKey = flag.String("private-key", "", "yggdrasil private key")
	var publicKey = flag.String("public-key", "", "yggdrasil public key")

	flag.Parse()

	if *privateKey == "" || *publicKey == "" {
		fmt.Println("yggdrasil --private-key and --public-key are required")
		os.Exit(1)
	}

	fmt.Println("Starting subprocesses")

	go startYggdrasil(*privateKey, *publicKey)
	go startPouchdb()

	// wait for yggdrasil
	waitForSocket(yggAdminListen)
	fmt.Printf("yggdrasil ready - my address: %v\n", getYggdrasilAddress())

	// wait for pouchdb
	waitForSocket(getPouchEndpoint())
	fmt.Println("pouchdb ready")

	// block forever, since we're not listening rn
	time.Sleep(1000000000 * time.Second)
}

func waitForSocket(endpoint string) {
	var conn net.Conn
	for {
		var err error
		conn, err = net.Dial("tcp", endpoint)
		if err == nil {
			break
		}
		time.Sleep(time.Second)
	}

	defer conn.Close()
}
