package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"os"
	"time"
)

const (
	appName           = "wg-fake"
	version           = "undefined"
	defaultTimeout    = time.Second * 5
	defaultNumPackets = 2
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

	conn, err := makeConnection(context.TODO(), *serverAddr, *localPort)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
	defer conn.Close()

	if err := sendRandomData(conn); err != nil {
		fmt.Fprint(os.Stderr, err)
	}
	fmt.Println("done...")
}

func makeConnection(ctx context.Context, serverAddr string, localPort uint) (net.Conn, error) {
	dialer := &net.Dialer{
		LocalAddr: &net.UDPAddr{
			Port: int(localPort),
		},
		Timeout: defaultTimeout,
	}
	conn, err := dialer.DialContext(ctx, "udp", serverAddr)
	if err != nil {
		return nil, fmt.Errorf("make connection failed: %w", err)
	}
	return conn, nil
}

func sendRandomData(conn net.Conn) error {
	packet := make([]byte, 16)
	for i := 0; i < defaultNumPackets; i++ {
		if _, err := conn.Write(packet); err != nil {
			return fmt.Errorf("failed to send data: %w", err)
		}
		fmt.Printf("packet %d of %d successfully sent\n", i+1, defaultNumPackets)
	}
	return nil
}
