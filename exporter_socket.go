package transporter

import (
	"fmt"
	"net"
)

// SocketDestination struct for handling socket-based exports
type SocketDestination struct {
	address string
	conn    net.Conn
}

// NewSocketDestination initializes and returns a new SocketDestination
// The connection is not established immediately.
func NewSocketDestination(address string) *SocketDestination {
	return &SocketDestination{
		address: address,
	}
}

// connect establishes the socket connection if it's not already connected.
func (s *SocketDestination) connect() error {
	if s.conn == nil {
		network, addressForNetwork := parseNetworkAddress(s.address)
		conn, err := net.Dial(network, addressForNetwork)
		if err != nil {
			return err
		}
		s.conn = conn
	}
	return nil
}

// WriteData writes the buffered data to the socket.
// It attempts to connect if it's not already connected and skips sending if an error occurs.
func (s *SocketDestination) WriteData(data string) error {
	err := s.connect()
	if err != nil {
		fmt.Println("Failed to connect to socket:", err)
		return nil // Skip sending if unable to connect
	}

	_, err = s.conn.Write([]byte(data))
	if err != nil {
		fmt.Println("Failed to write to socket:", err)
		s.conn.Close() // Close broken connection
		s.conn = nil   // Reset connection for next attempt
	}

	return nil
}

// Close closes the socket connection.
func (s *SocketDestination) Close() error {
	if s.conn != nil {
		err := s.conn.Close()
		s.conn = nil
		return err
	}
	return nil
}
