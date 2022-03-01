//go:build !linux
// +build !linux

package internal

import (
	"golang.org/x/sys/unix"
)

type Termios struct {
	term *unix.Termios
}

var term *Termios

func GetTermios() *Termios {
	return term
}

func IoctlGetTermios() *Termios {
	return nil
}

func IoctlSetTermios(termios *Termios) {
}
