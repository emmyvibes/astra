package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net"
	"os"
	"os/exec"

	yggadmin "github.com/yggdrasil-network/yggdrasil-go/src/admin"
	yggdefaults "github.com/yggdrasil-network/yggdrasil-go/src/defaults"
)

const adminListen = "127.0.0.1:9090"

func startYggdrasil() {
	initialConfig := yggdefaults.GenerateConfig()

	// default static peers
	initialConfig.Peers = []string{"tls://supergay.network:9001", "tls://102.223.180.74:993", "tls://ygg.mnpnk.com:443"}
	// yggdrasil management port
	initialConfig.AdminListen = adminListen

	// TODO: persist your pub/privkeys somehow
	initialConfig.NewKeys()

	// fmt.Println(initialConfig)

	cmd := exec.Command("yggdrasil", "-useconf")
	// pipe output
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	stdin, err := cmd.StdinPipe()
	if err != nil {
		panic(err)
	}

	if err := cmd.Start(); err != nil {
		panic(err)
	}

	// pass config
	configData, _ := json.Marshal(initialConfig)
	// fmt.Println(string(configData))
	io.Copy(stdin, bytes.NewReader(configData))
	// we need to close stdin when done using
	defer stdin.Close()
}

func getYggdrasilAddress() string {
	conn, err := net.Dial("tcp", adminListen)
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	decoder := json.NewDecoder(conn)
	encoder := json.NewEncoder(conn)
	send := &yggadmin.AdminSocketRequest{
		Name: "getSelf",
	}
	recv := &yggadmin.AdminSocketResponse{}

	if err := encoder.Encode(&send); err != nil {
		panic(err)
	}
	if err := decoder.Decode(&recv); err != nil {
		panic(err)
	}

	var resp yggadmin.GetSelfResponse
	if err := json.Unmarshal(recv.Response, &resp); err != nil {
		panic(err)
	}

	return resp.IPAddress
}
