package internal

import (
	"context"
	"fmt"

	"github.com/chyroc/go-aliyundrive"
)

type CommandDownload struct {
	cli  *Cli
	name string
}

func (r *CommandDownload) Run() error {
	if err := r.cli.setupDrive(); err != nil {
		return err
	}

	file, err := r.cli.findFileByName(r.name)
	if err != nil {
		return err
	}
	if file == nil {
		return fmt.Errorf("不存在文件 %q", r.name)
	}

	err = r.cli.ali.File.DownloadFile(context.Background(), &aliyundrive.DownloadFileReq{
		DriveID:      r.cli.driveID,
		FileID:       file.FileID,
		DistDir:      r.cli.downloadDir,
		ConflictType: aliyundrive.DownloadFileConflictTypeAutoRename,
	})
	if err != nil {
		return err
	}

	fmt.Println("下载成功.")

	return nil
}
