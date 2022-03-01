//go:build windows
// +build windows

package internal

import (
	"golang.org/x/sys/unix"
)

type Termios struct{}

var term *Termios

func GetTermios() *Termios {
	return term
}

func IoctlGetTermios() *unix.Termios {
	return nil
}

func IoctlSetTermios(termios *unix.Termios) {
}
