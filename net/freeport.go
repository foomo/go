package net

import (
	"net"
)

// FreePort returns a free port on localhost
func FreePort() (string, error) {
	l, err := net.Listen("tcp", "127.0.0.1:0") //nolint:noctx
	if err != nil {
		return "", err
	}

	addr := l.Addr().String()
	if err := l.Close(); err != nil {
		return "", err
	}

	return addr, nil
}
