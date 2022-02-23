package internal

import (
	"context"
	"fmt"

	"github.com/chyroc/go-aliyundrive"
)

type CommandRename struct {
	cli  *Cli
	from string
	to   string
}

func (r *CommandRename) Run() error {
	if err := r.cli.setupDrive(); err != nil {
		return err
	}

	if err := r.cli.setupFiles(); err != nil {
		return err
	}

	oldFile, err := r.cli.findFileByName(r.from)
	if err != nil {
		return err
	} else if oldFile == nil {
		return fmt.Errorf("文件: %q 不存在", r.from)
	}
	_, err = r.cli.ali.File.RenameFile(context.Background(), &aliyundrive.RenameFileReq{
		DriveID:       r.cli.driveID,
		FileID:        oldFile.FileID,
		CheckNameMode: "refuse",
		Name:          r.to,
	})
	if err != nil {
		return err
	}

	fmt.Println("重命名成功.")

	go r.cli.refreshFiles()

	return nil
}
