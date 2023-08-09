package net

import (
	"errors"
	"net"
)

// FreePort returns a free port on localhost
func FreePort() (int, error) {
	var a *net.TCPAddr
	a, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		return 0, err
	}
	l, err := net.ListenTCP("tcp", a)
	if err != nil {
		return 0, err
	}
	defer l.Close()
	t, ok := l.Addr().(*net.TCPAddr)
	if !ok {
		return 0, errors.New("failed to cast TCPAddr")
	}
	return t.Port, nil
}
