package main

import (
	"fmt"

	"github.com/c-bata/go-prompt"
	"github.com/chyroc/aliyundrive-cli/internal"
)

func main() {
	oldTermiosPtr := internal.IoctlGetTermios()
	defer internal.IoctlSetTermios(oldTermiosPtr)

	cli := internal.NewCli()
	fmt.Println("阿里云盘命令行客户端")

	p := prompt.New(cli.Executor, cli.Completer, prompt.OptionLivePrefix(cli.Prefix))
	p.Run()
}
