package sender

import (
	"context"
	"fmt"
	"net"
	"time"
)

const (
	defaultTimeout    = time.Second * 5
	defaultNumPackets = 2
	defaultPacketSize = 16
)

type transporter struct {
	packet []byte
	conn   net.Conn
}

func New(ctx context.Context, serverAddr string, localPort uint) (*transporter, error) {
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

	return &transporter{
		packet: make([]byte, defaultPacketSize),
		conn:   conn,
	}, nil
}

func (t *transporter) send(body []byte) error {
	if _, err := t.conn.Write(body); err != nil {
		return fmt.Errorf("failed to send data: %w", err)
	}

	return nil
}

func (t *transporter) Magic() error {
	defer t.conn.Close()
	for i := 0; i < defaultNumPackets; i++ {
		if err := t.send(t.packet); err != nil {
			return fmt.Errorf("failed to send data: %w", err)
		}
		fmt.Printf("packet %d of %d successfully sent\n", i+1, defaultNumPackets)
	}
	return nil
}
