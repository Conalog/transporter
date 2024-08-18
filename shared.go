package transporter

import "strings"

func parseNetworkAddress(input string) (network, address string) {
	// Check for known prefixes and split accordingly
	switch {
	case strings.HasPrefix(input, "tcp://"):
		network = "tcp"
		address = strings.TrimPrefix(input, "tcp://")
	case strings.HasPrefix(input, "udp://"):
		network = "udp"
		address = strings.TrimPrefix(input, "udp://")
	case strings.HasPrefix(input, "unix://"):
		network = "unix"
		address = strings.TrimPrefix(input, "unix://")
	default:
		// Default to Unix socket if no known prefix is found
		network = "unix"
		address = input
	}
	return
}
