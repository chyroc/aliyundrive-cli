package internal

import (
	"golang.org/x/sys/unix"
)

func IoctlGetTermios() *unix.Termios {
	return nil
}

func IoctlSetTermios(termios *unix.Termios) {
}
