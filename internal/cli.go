package internal

import (
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/c-bata/go-prompt"
	"github.com/chyroc/go-aliyundrive"
	"github.com/golang-collections/collections/stack"
)

type Cli struct {
	// config
	downloadDir string

	// internal
	ali            *aliyundrive.AliyunDrive
	setupDriveOnce sync.Once
	driveID        string

	// runtime
	fileNames     []string
	fileIDs       []string
	currentFileID string
	fileStack     *stack.Stack
	files         []*aliyundrive.File
}

func NewCli(homedir string) *Cli {
	fileStack := stack.New()
	fileStack.Push("root")
	_ = os.MkdirAll(homedir, 0o777)

	return &Cli{
		ali:           aliyundrive.New(),
		fileNames:     []string{""},
		fileIDs:       []string{"root"},
		currentFileID: "root",
		fileStack:     fileStack,
		downloadDir:   homedir,
	}
}

func (r *Cli) Executor(input string) {
	cmd, err := r.ParseCommand(input)
	if err != nil {
		fmt.Println(err)
		return
	} else if cmd == nil {
		return
	}
	err = cmd.Run()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (r *Cli) Completer(doc prompt.Document) (res []prompt.Suggest) {
	cmd, args := parsePrefixCommand(doc.Text, []string{"cd", "download", "ls", "rm"})
	args = strings.ToLower(args)
	if cmd != "" {
		for _, v := range r.files {
			if args == "" {
				res = append(res, prompt.Suggest{Text: v.Name})
			} else if inText(args, strings.ToLower(v.Name)) {
				res = append(res, prompt.Suggest{Text: v.Name})
			}
		}
	}
	if strings.HasPrefix(doc.Text, "mv ") {
		args := doc.Text[len("mv "):]
		if strings.HasSuffix(args, " ") && !strings.HasSuffix(args, "\\ ") {
			// 以空格结尾，则提示所有文件
			for _, v := range r.files {
				res = append(res, prompt.Suggest{Text: v.Name})
			}
		} else {
			// 否则，取最后的单词，匹配单词
			argsList := splitSpace(args)
			if len(argsList) == 0 {
				for _, v := range r.files {
					res = append(res, prompt.Suggest{Text: v.Name})
				}
			} else {
				args := argsList[len(argsList)-1]
				for _, v := range r.files {
					if args == "" {
						res = append(res, prompt.Suggest{Text: v.Name})
					} else if inText(args, strings.ToLower(v.Name)) {
						res = append(res, prompt.Suggest{Text: v.Name})
					}
				}
			}
		}
	}

	return res
}

func parsePrefixCommand(s string, prefixs []string) (string, string) {
	for _, v := range prefixs {
		if strings.HasPrefix(s, v+" ") {
			return v, strings.TrimSpace(s[len(v+" "):])
		}
	}
	return "", ""
}

func (r *Cli) Prefix() (string, bool) {
	if len(r.fileNames) == 1 {
		return "/> ", true
	}
	return strings.Join(r.fileNames, "/") + "> ", true
}

func hasPrefix(s string, prefix []string) bool {
	for _, v := range prefix {
		if strings.HasPrefix(s, v) {
			return true
		}
	}
	return false
}
