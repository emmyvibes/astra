package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Starting subprocesses")

	go startYggdrasil()
	go startPouchdb()

	time.Sleep(5 * time.Second)
	fmt.Printf("my address: %v\n", getYggdrasilAddress())

	// block forever, since we're not listening rn
	// select {}
	time.Sleep(1000000000 * time.Second)
}
