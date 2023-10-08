package main

import (
	"fmt"
	"os"
	"os/exec"
)

const pouchHost = "::"
const pouchPort = "5985"

func getPouchEndpoint() string {
	return fmt.Sprintf("[%v]:%v", pouchHost, pouchPort)
}

func startPouchdb() {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	storageDir := fmt.Sprintf("%v/pouchdb-data", wd)

	cmd := exec.Command("pouchdb-server", "--dir", storageDir, "-o", "::", "-p", "5985")
	// pipe output
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		panic(err)
	}
}

func configureReplication(targetIp string) {

}
