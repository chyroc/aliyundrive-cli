//go:build !linux
// +build !linux

package internal

import (
	"golang.org/x/sys/unix"
)

var term *unix.Termios

func IoctlGetTermios() *unix.Termios {
	return nil
}

func IoctlSetTermios(termios *unix.Termios) {
}
