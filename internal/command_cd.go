package internal

import (
	"fmt"
)

type CommandCd struct {
	cli *Cli
	dir string
}

func (r *CommandCd) Run() error {
	if err := r.cli.setupDrive(); err != nil {
		return err
	}
	if r.dir == ".." {
		return r.cli.checkoutToParentDir()
	}

	if err := r.cli.setupFiles(); err != nil {
		return err
	}

	for _, v := range r.cli.files {
		if v.Name == r.dir && v.Type == "folder" {
			return r.cli.checkoutDir(v.FileID, v.Name)
		}
	}
	return fmt.Errorf("文件夹 %q 不存在", r.dir)
}

func (r *Cli) checkoutToParentDir() error {
	if len(r.fileNames) == 1 {
		return fmt.Errorf("当前在根目录，无法前往父目录")
	}
	r.fileStack.Pop()

	r.currentFileID = r.fileStack.Peek().(string)
	r.fileNames = r.fileNames[:len(r.fileNames)-1]
	r.files = nil
	return nil
}

func (r *Cli) checkoutDir(fileID, name string) error {
	r.currentFileID = fileID
	r.fileNames = append(r.fileNames, name)
	r.files = nil
	r.fileStack.Push(fileID)
	return nil
}
