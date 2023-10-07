import (
	"os/exec"
)

func main() {
	cmd := exec.Command("node", "index.js")
	cd
}

/*
package main

import (
	"fmt"
	"os/exec"

	couchdb "github.com/leesper/couchdb-golang"
)

func main() {
	cmd := exec.Command("node", "path/to/your/node/script.js")

	// cmd.Stdout = os.Stdout
	// cmd.Stderr = os.Stderr

	err := cmd.Run()

	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("started server")
	server, err := couchdb.NewServer("")
	fmt.Print(server)
	fmt.Print(err)

	server.Login()
}*/