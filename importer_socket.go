package transporter

import (
	"bufio"
	"fmt"
	"net"
)

// SocketSource struct for handling socket-based imports
type SocketSource struct {
	address  string
	listener net.Listener
	conn     net.Conn
	reader   *bufio.Reader
}

// NewSocketSource initializes and returns a new SocketSource
func NewSocketSource(address string) *SocketSource {
	return &SocketSource{
		address: address,
	}
}

// startListener starts listening for connections on the specified address.
func (s *SocketSource) startListener() error {
	if s.listener == nil {
		network, addressForNetwork := parseNetworkAddress(s.address)
		listener, err := net.Listen(network, addressForNetwork)
		if err != nil {
			return err
		}
		s.listener = listener
	}
	return nil
}

// acceptConnection accepts a new connection if one is not already established.
func (s *SocketSource) acceptConnection() error {
	if s.conn == nil {
		err := s.startListener()
		if err != nil {
			return fmt.Errorf("failed to start listener: %w", err)
		}

		conn, err := s.listener.Accept()
		if err != nil {
			return err
		}
		s.conn = conn
		s.reader = bufio.NewReader(conn)
	}
	return nil
}

// ReadData reads a line of data from the socket.
func (s *SocketSource) ReadData() (string, error) {
	err := s.acceptConnection()
	if err != nil {
		return "", fmt.Errorf("failed to accept connection: %w", err)
	}

	line, err := s.reader.ReadString('\n')
	if err != nil {
		return "", err // Connection closed or read error
	}

	return line, nil
}

// Close closes the socket connection and listener.
func (s *SocketSource) Close() error {
	if s.conn != nil {
		s.conn.Close()
		s.conn = nil
	}
	if s.listener != nil {
		s.listener.Close()
		s.listener = nil
	}
	s.reader = nil
	return nil
}
