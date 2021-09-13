package internal

import (
	"context"

	"github.com/chyroc/go-aliyundrive"
)

type CommandLs struct {
	cli *Cli
}

func (r *CommandLs) Run() error {
	if err := r.cli.setupDrive(); err != nil {
		return err
	}

	if err := r.cli.setupFiles(); err != nil {
		return err
	}

	r.cli.PrintFiles(r.cli.files)
	return nil
}

func (r *Cli) setupFiles() error {
	if len(r.files) == 0 {
		resp, err := r.ali.File.GetFileList(context.Background(), &aliyundrive.GetFileListReq{
			GetAll:       true,
			DriveID:      r.driveID,
			ParentFileID: r.currentFileID,
			Limit:        100,
		})
		if err != nil {
			return err
		}
		r.files = resp.Items
	}
	return nil
}

func (r *Cli) refreshFiles() error {
	r.files = nil
	return r.setupFiles()
}

func (r *Cli) findFileByName(name string) (*aliyundrive.File, error) {
	if err := r.setupFiles(); err != nil {
		return nil, err
	}
	for _, v := range r.files {
		if v.Name == name {
			return v, nil
		}
	}
	return nil, nil
}

func (r *Cli) removeByName(name string) (*aliyundrive.File, error) {
	if err := r.setupFiles(); err != nil {
		return nil, err
	}
	files := []*aliyundrive.File{}
	var removed *aliyundrive.File
	for _, v := range r.files {
		if v.Name == name {
			removed = v
		} else {
			files = append(files, v)
		}
	}
	r.files = files
	return removed, nil
}
