package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net"
	"os"
	"os/exec"

	yggadmin "github.com/yggdrasil-network/yggdrasil-go/src/admin"
	yggconfig "github.com/yggdrasil-network/yggdrasil-go/src/config"
	yggdefaults "github.com/yggdrasil-network/yggdrasil-go/src/defaults"
)

const yggAdminListen = "[::]:9090"

func startYggdrasil(privateKey string, publicKey string) {
	initialConfig := yggdefaults.GenerateConfig()

	// default static peers
	initialConfig.Peers = []string{"tls://supergay.network:9001", "tls://102.223.180.74:993", "tls://ygg.mnpnk.com:443"}
	// yggdrasil management port
	initialConfig.AdminListen = yggAdminListen
	// FOR DEVELOPMENT: don't automatically connect to LAN ygg peers
	initialConfig.MulticastInterfaces = []yggconfig.MulticastInterfaceConfig{}

	// TODO: persist your pub/privkeys somehow
	// initialConfig.NewKeys()

	initialConfig.PrivateKey = privateKey
	initialConfig.PublicKey = publicKey

	// fmt.Println(initialConfig.PrivateKey, initialConfig.PublicKey)

	cmd := exec.Command("sudo", "yggdrasil", "-useconf")
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
	conn, err := net.Dial("tcp", yggAdminListen)
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
