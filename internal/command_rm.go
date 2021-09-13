package internal

import (
	"context"
	"fmt"

	"github.com/chyroc/go-aliyundrive"
)

type CommandRm struct {
	cli  *Cli
	name string
}

func (r *CommandRm) Run() error {
	if err := r.cli.setupDrive(); err != nil {
		return err
	}

	if err := r.cli.setupFiles(); err != nil {
		return err
	}

	file, err := r.cli.findFileByName(r.name)
	if err != nil {
		return err
	}
	if file == nil {
		return fmt.Errorf("文件 %q 不存在", r.name)
	}

	_, err = r.cli.ali.File.DeleteFile(context.Background(), &aliyundrive.DeleteFileReq{
		DriveID: r.cli.driveID,
		FileID:  file.FileID,
	})
	if err != nil {
		return err
	}
	fmt.Println("删除成功.")
	_, _ = r.cli.removeByName(r.name)

	return nil
}
