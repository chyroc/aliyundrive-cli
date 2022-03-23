package internal

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
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
	defer r.cli.refreshFiles()

	file := r.file
	files, err := filepath.Glob(file)
	if err != nil {
		return err
	}
	if len(files) == 0 {
		return fmt.Errorf("找不到文件 %q", file)
	}
	go func() {
		for _, file := range files {
			err = r.upload(file, r.cli.driveID, r.cli.currentFileID)
			if err != nil {
				fmt.Printf("upload %s filed: %s\n", file, err.Error())
				break
			}
		}
		go r.cli.refreshFiles()
	}()
	return nil
}

func (r *CommandUpload) upload(file string, driveID string, fileID string) error {
	if strings.HasPrefix(file, "~") {
		home, _ := os.UserHomeDir()
		file = home + file[1:]
	}
	file, err := filepath.Abs(file)
	if err != nil {
		return err
	}
	fileInfo, err := os.Stat(file)
	if err != nil {
		return err
	}
	if !fileInfo.IsDir() {
		_, err = r.cli.ali.File.UploadFile(context.Background(), &aliyundrive.UploadFileReq{
			DriveID:         driveID,
			ParentID:        fileID,
			FilePath:        file,
			ShowProgressBar: true,
		})
		if err != nil {
			return err
		}
		fmt.Printf("%s 上传成功.\n", file)
	} else {
		// mkdir
		files, err := ioutil.ReadDir(file)
		if err != nil {
			log.Fatal(err)
		}
		response, err := r.cli.ali.File.CreateFolder(context.Background(), &aliyundrive.CreateFolderReq{
			DriveID:      r.cli.driveID,
			ParentFileID: r.cli.currentFileID,
			Name:         fileInfo.Name(),
		})
		if err != nil {
			return err
		}
		if err := r.cli.checkoutDir(response.FileID, fileInfo.Name()); err != nil {
			return err
		}
		for _, subFile := range files {
			err := r.upload(filepath.Join(file, subFile.Name()), response.DriveID, response.FileID)
			if err != nil {
				return err
			}
			fmt.Printf("%s 上传成功.\n", filepath.Join(file, subFile.Name()))
		}
	}
	return nil
}
