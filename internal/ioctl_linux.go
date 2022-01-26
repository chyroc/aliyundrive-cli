//go:build linux
// +build linux

package internal

import (
	"os"

	"golang.org/x/sys/unix"
)

var term *unix.Termios

func IoctlGetTermios() *unix.Termios {
	termios, _ := unix.IoctlGetTermios(int(os.Stdin.Fd()), unix.TCGETS)
	term = termios
	return termios
}

func IoctlSetTermios(termios *unix.Termios) {
	unix.IoctlSetTermios(int(os.Stdin.Fd()), unix.TCSETS, termios)
}
