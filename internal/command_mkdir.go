package internal

import (
	"context"
	"fmt"

	"github.com/chyroc/go-aliyundrive"
)

type CommandMkdir struct {
	cli *Cli
	dir string
}

func (r *CommandMkdir) Run() error {
	if err := r.cli.setupDrive(); err != nil {
		return err
	}
	if !isValidDirName(r.dir) {
		return fmt.Errorf("%q 不是合法的文件夹名称", r.dir)
	}
	_, err := r.cli.ali.File.CreateFolder(context.Background(), &aliyundrive.CreateFolderReq{
		DriveID:      r.cli.driveID,
		ParentFileID: r.cli.currentFileID,
		Name:         r.dir,
	})
	if err != nil {
		return err
	}
	fmt.Println("创建成功.")
	go r.cli.refreshFiles()

	return nil
}
