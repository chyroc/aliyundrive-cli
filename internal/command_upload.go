package internal

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/chyroc/go-aliyundrive"
)

type CommandUpload struct {
	cli  *Cli
	file string
}

func (r *CommandUpload) Run() error {
	if err := r.cli.setupDrive(); err != nil {
		return err
	}

	file := r.file
	if strings.HasPrefix(file, "~") {
		home, _ := os.UserHomeDir()
		file = home + file[1:]
	}
	file, err := filepath.Abs(file)
	if err != nil {
		return err
	}
	if _, err := os.Stat(file); err != nil {
		return err
	}

	_, err = r.cli.ali.File.UploadFile(context.Background(), &aliyundrive.UploadFileReq{
		DriveID:  r.cli.driveID,
		ParentID: r.cli.currentFileID,
		FilePath: file,
	})
	if err != nil {
		return err
	}
	fmt.Println("上传成功.")

	go r.cli.refreshFiles()

	return nil
}
