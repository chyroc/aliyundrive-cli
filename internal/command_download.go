package internal

import (
	"context"
	"fmt"
	"os"
	"path"

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
	go r.download(r.cli.downloadDir, file)
	return nil
}

func (r *CommandDownload) download(dir string, file *aliyundrive.File) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.MkdirAll(dir, os.ModePerm)
	}
	// 下载普通文件
	if file.Type != "folder" {
		err := r.cli.ali.File.DownloadFile(context.Background(), &aliyundrive.DownloadFileReq{
			DriveID:         r.cli.driveID,
			FileID:          file.FileID,
			DistDir:         dir,
			ConflictType:    aliyundrive.DownloadFileConflictTypeAutoRename,
			ShowProgressBar: true,
		})
		if err != nil {
			return err
		}
		return nil
	} else {
		// 递归下载文件夹
		// 获取文件夹下面的所有文件
		res, err := r.cli.ali.File.GetFileList(context.Background(), &aliyundrive.GetFileListReq{
			GetAll:       true,
			DriveID:      r.cli.driveID,
			ParentFileID: file.FileID,
			Limit:        0,
		})
		if err != nil {
			return err
		}
		for _, f := range res.Items {
			// 忽略错误
			r.download(path.Join(dir, file.Name), f)
		}
	}
	return nil

}
