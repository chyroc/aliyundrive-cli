//go:build linux
// +build linux

package internal

import (
	"os"

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
	termios, _ := unix.IoctlGetTermios(int(os.Stdin.Fd()), unix.TCGETS)
	term = &Termios{term: termios}
	return term
}

func IoctlSetTermios(termios *Termios) {
	if termios != nil && termios.term != nil {
		unix.IoctlSetTermios(int(os.Stdin.Fd()), unix.TCSETS, termios.term)
	}
}
