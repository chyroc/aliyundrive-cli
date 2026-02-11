package internal

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"path"
	"syscall"

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

	// 创建支持 Ctrl+C 取消的 context
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	fmt.Printf("开始下载 %s 到 %s ...\n", file.Name, r.cli.downloadDir)
	if err := r.download(ctx, r.cli.downloadDir, file); err != nil {
		if err == context.Canceled {
			return fmt.Errorf("下载已取消")
		}
		return fmt.Errorf("下载失败: %v", err)
	}
	fmt.Printf("下载完成: %s\n", file.Name)
	return nil
}

func (r *CommandDownload) download(ctx context.Context, dir string, file *aliyundrive.File) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.MkdirAll(dir, os.ModePerm)
	}
	// 下载普通文件
	if file.Type != "folder" {
		err := r.cli.ali.File.DownloadFile(ctx, &aliyundrive.DownloadFileReq{
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
		res, err := r.cli.ali.File.GetFileList(ctx, &aliyundrive.GetFileListReq{
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
			r.download(ctx, path.Join(dir, file.Name), f)
		}
	}
	return nil
}
