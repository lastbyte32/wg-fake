package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"random-udp-sender/internal/app"
)

const (
	appName = "wg-fake"
	version = "undefined"
)

func usage() string {
	return fmt.Sprintf("%s -s <WG SERVER ADDRESS:PORT> -p <LOCAL WG PORT>\n", appName)
}

func main() {
	localPort := flag.Uint("p", 0, "local wireguard port")
	serverAddr := flag.String("s", "", "wireguard server address")
	flag.Parse()

	if *serverAddr == "" {
		fmt.Fprintf(os.Stderr, "error: server address is required\n\n%s\n", usage())
		os.Exit(1)
	}
	if *localPort == 0 || *localPort > 65535 {
		fmt.Fprintf(os.Stderr, "error: local port must be between 1 and 65535\n\n%s\n", usage())
		os.Exit(1)
	}

	fmt.Printf("%s: %s\nprobing %s from local port %d\n", appName, version, *serverAddr, *localPort)

	if err := app.Run(context.Background(), *serverAddr, *localPort); err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println("done...")
}
