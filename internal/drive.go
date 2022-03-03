package internal

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/chyroc/go-aliyundrive"
	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
)

func (r *Cli) setupDrive() (finalErr error) {
	r.setupDriveOnce.Do(func() {
		user, err := r.ali.Auth.LoginByQrcode(context.TODO(),
			&aliyundrive.LoginByQrcodeReq{
				SmallQrCode: true})
		if err != nil {
			finalErr = err
			return
		}
		r.driveID = user.DefaultDriveID
	})
	return finalErr
}

func (r *Cli) PrintFiles(files []*aliyundrive.File) {
	if len(files) == 0 {
		fmt.Println("没有文件")
		return
	}

	header := []string{
		"名称", "类型", "大小", "修改时间",
	}
	data := [][]string{}
	for _, f := range files {
		data = append(data, []string{
			formatFileType(f.Name, f.Type), f.Type, formatFileSize(f.Size), formatTime(f.UpdatedAt),
		})
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(header)
	// table.SetBorder(false)
	table.SetAutoWrapText(false)
	table.AppendBulk(data)
	table.Render()
}

func formatFileType(name, t string) string {
	if t == "folder" {
		return colorDirFormat.Sprint(name)
	}
	return name
}

var colorDirFormat = color.New(color.FgCyan, color.Bold)

func formatTime(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

func formatFileSize(size int64) string {
	s := float64(size)
	var unit string
	for _, unit = range []string{"B", "K", "M", "G", "T", "P"} {
		if s < 1024 || unit == "P" {
			break
		}
		s = s / 1024
	}
	if fmt.Sprintf("%.2f", s) == fmt.Sprintf("%d.00", int64(s)) {
		return fmt.Sprintf("%d%s", int64(s), unit)
	}
	return fmt.Sprintf("%.2f%s", s, unit)
}
