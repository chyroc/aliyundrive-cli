package main

import (
	"flag"
	"fmt"
	"os"

	"runtime/debug"

	"github.com/c-bata/go-prompt"
	"github.com/chyroc/aliyundrive-cli/internal"
)

var version bool

func init() {
	flag.BoolVar(&version, "version", false, "Print program version")
	if !flag.Parsed() {
		flag.Parse()
	}
	if version {
		info, ok := debug.ReadBuildInfo()
		if ok {
			println(info.Main.Version)
		}
		os.Exit(0)
	}
}

func main() {
	oldTermiosPtr := internal.IoctlGetTermios()
	defer internal.IoctlSetTermios(oldTermiosPtr)

	cli := internal.NewCli()
	fmt.Println("阿里云盘命令行客户端")

	p := prompt.New(cli.Executor, cli.Completer, prompt.OptionLivePrefix(cli.Prefix))
	p.Run()
}
