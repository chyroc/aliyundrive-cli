package internal

import (
	"context"
	"fmt"

	"github.com/chyroc/go-aliyundrive"
)

type CommandMv struct {
	cli  *Cli
	from string
	to   string
}

func (r *CommandMv) Run() error {
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

	newFile, err := r.cli.findFileByName(r.to)
	if err != nil {
		return err
	} else if oldFile == nil {
		return fmt.Errorf("文件夹: %q 不存在", r.to)
	}

	_, err = r.cli.ali.File.MoveFile(context.Background(), &aliyundrive.MoveFileReq{
		DriveID:        r.cli.driveID,
		FileID:         oldFile.FileID,
		ToDriveID:      r.cli.driveID,
		ToParentFileID: newFile.FileID,
	})
	if err != nil {
		return err
	}

	fmt.Println("移动成功.")

	go r.cli.refreshFiles()

	return nil
}
