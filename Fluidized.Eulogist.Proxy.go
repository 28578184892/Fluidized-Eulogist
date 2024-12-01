package main

import (
	authenticator "Eulogist/core/fluid_auth"
	Eulogist "Eulogist/eulogist"
	_ "embed"
	"fmt"
	"os"
	"strings"

	"github.com/pterm/pterm"
)

func main() {

	fmt.Println("Fluidized-Eulogist ProxyService...")

	var serverIp string
	var listenPort string

	for _, v := range os.Args {
		if strings.Contains(v, "serverIp") {
			serverIp = strings.Split(v, "=")[1]
		}

		if strings.Contains(v, "listenPort") {
			listenPort = strings.Split(v, "=")[1]
		}
	}

	if serverIp == "" || listenPort == "" {
		fmt.Println("No ServerIP or ListenPort Detected! Service will exit now...")
		os.Exit(-1)
	}

	authenticator.ServerIP = serverIp
	authenticator.ListenPort = listenPort

	err := Eulogist.Eulogist()
	if err != nil {
		pterm.Error.Println(err)
	}

	fmt.Println()
}
