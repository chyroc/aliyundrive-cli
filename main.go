package main

import (
	"fmt"
	"os"

	"github.com/c-bata/go-prompt"
	"github.com/chyroc/aliyundrive-cli/internal"
	"golang.org/x/sys/unix"
)

func main() {
	oldTermiosPtr, _ := unix.IoctlGetTermios(int(os.Stdin.Fd()), unix.TCGETS)
	defer func() {
		unix.IoctlSetTermios(int(os.Stdin.Fd()), unix.TCSETS, oldTermiosPtr)
	}()
	cli := internal.NewCli()
	fmt.Println("阿里云盘命令行客户端")

	p := prompt.New(cli.Executor, cli.Completer, prompt.OptionLivePrefix(cli.Prefix))
	p.Run()
}
