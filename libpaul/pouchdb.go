package main

import (
	"fmt"
	"os"
	"os/exec"
)

func startPouchdb() {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	storageDir := fmt.Sprintf("%v/pouchdb-data", wd)

	cmd := exec.Command("sudo", "-u", "nik", "pouchdb-server", "--dir", storageDir)
	// pipe output
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		panic(err)
	}
}
