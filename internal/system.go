package internal

import "os"

var fs []func()

// 取消
func Cancel() {
	// 没有可以取消的任务，直接推出程序
	if len(fs) == 0 {
		Exit(0)
	} else {
		fs[0]()
		fs = fs[1:]
	}
}
func Exit(code int) {
	IoctlSetTermios(term)
	os.Exit(code)
}
