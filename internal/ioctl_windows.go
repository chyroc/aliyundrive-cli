//go:build windows
// +build windows

package internal

type Termios struct{}

var term *Termios

func GetTermios() *Termios {
	return term
}

func IoctlGetTermios() *Termios {
	return nil
}

func IoctlSetTermios(termios *Termios) {
}
