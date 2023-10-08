package main

import (
	"context"
	"fmt"
	"os/exec"
	"time"

	_ "github.com/go-kivik/couchdb/v3"

	"github.com/go-kivik/kivik"
	"github.com/go-kivik/kivik/v3"
)

// The CouchDB driver



func main() {
	fmt.Println("SUBPROCESS です")

	//run couch
	exec.Command("node", "couch.js")

	//Don't @ me, there's probably some way to wait properly
	time.Sleep(2 * time.Second)

	dsn := "http://localhost:5984"
	client, err := kivik.New(context.TODO(), "couch", dsn)

	if err != nil {
		panic(err)
	}
	err.
}