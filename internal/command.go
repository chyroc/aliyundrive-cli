package internal

import (
	"fmt"
	"strings"
)

type Command interface {
	Run() error
}

func (r *Cli) ParseCommand(input string) (Command, error) {
	input = strings.TrimSpace(input)
	if input == "" {
		return nil, nil
	}
	if input == "ls" || strings.HasPrefix(input, "ls ") {
		return &CommandLs{cli: r}, nil
	}
	if strings.HasPrefix(input, "cd ") {
		return &CommandCd{cli: r, dir: strings.TrimSpace(input[len("cd "):])}, nil
	}
	if strings.HasPrefix(input, "mkdir ") {
		return &CommandMkdir{cli: r, dir: strings.TrimSpace(input[len("mkdir "):])}, nil
	}
	if strings.HasPrefix(input, "rm ") {
		return &CommandRm{cli: r, name: strings.TrimSpace(input[len("rm "):])}, nil
	}
	if strings.HasPrefix(input, "upload ") {
		return &CommandUpload{cli: r, file: strings.TrimSpace(input[len("upload "):])}, nil
	}
	if strings.HasPrefix(input, "download ") {
		return &CommandDownload{cli: r, name: strings.TrimSpace(input[len("download "):])}, nil
	}
	if strings.HasPrefix(input, "mv ") {
		l := splitSpace(strings.TrimSpace(input[len("mv "):]))
		if len(l) != 2 {
			return nil, fmt.Errorf("mv 命令不合法，需要两个以空格区分的参数，如: mv old new")
		}
		return &CommandMv{cli: r, from: strings.TrimSpace(l[0]), to: strings.TrimSpace(l[1])}, nil
	}
	return nil, fmt.Errorf("不支持的命令: %s", input)
}
