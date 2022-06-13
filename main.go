package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"runtime/debug"

	"github.com/c-bata/go-prompt"
	"github.com/chyroc/aliyundrive-cli/internal"
)

var (
	printVersion bool
	dir          string
)

func init() {
	home, _ := os.UserHomeDir()
	downloadDir := path.Join(home, "/Downloads/aliyundrive-cli")
	flag.BoolVar(&printVersion, "version", false, "Print program version")
	flag.StringVar(&dir, "dir", downloadDir, "File download directory")
	if !flag.Parsed() {
		flag.Parse()
	}
	if printVersion {
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
	os.Stdout.Sync()
	cli := internal.NewCli(dir)
	fmt.Println("阿里云盘命令行客户端")

	p := prompt.New(cli.Executor, cli.Completer, prompt.OptionLivePrefix(cli.Prefix), prompt.OptionAddKeyBind(prompt.KeyBind{
		Key: prompt.ControlC,
		Fn: func(b *prompt.Buffer) {
			internal.Cancel()
		},
	}))

	p.Run()
}
