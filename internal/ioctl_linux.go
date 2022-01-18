package internal

import "golang.org/x/sys/unix"

func IoctlGetTermios() *unix.Termios {
	termios, _ := unix.IoctlGetTermios(int(os.Stdin.Fd()), unix.TCGETS)
	return termios
}

func IoctlSetTermios(termios *unix.Termios) {
	unix.IoctlSetTermios(int(os.Stdin.Fd()), unix.TCSETS, termios)
}
