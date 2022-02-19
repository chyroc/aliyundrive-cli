package internal

import (
	"context"
	"fmt"
	"sort"

	"github.com/chyroc/aliyundrive-cli/internal/helper_ui"
	"github.com/chyroc/go2tv/devices"
	"github.com/chyroc/go2tv/sendtotv"
)

type CommandToTv struct {
	cli  *Cli
	name string
}

func (r *CommandToTv) Run() error {
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

	device, err := r.getDevice()
	if err != nil {
		return err
	}

	body, err := r.cli.ali.File.DownloadFileStream(context.Background(), r.cli.driveID, file.FileID)
	if err != nil {
		return err
	}
	defer body.Close()

	return sendtotv.SendReadCloser(&sendtotv.Media{Body: body, Name: file.Name}, nil, device.URL)
}

func (r *CommandToTv) getDevice() (*Device, error) {
	deviceList, err := r.loadDevice()
	if err != nil {
		return nil, err
	}

	items := []string{}
	for _, v := range deviceList {
		items = append(items, v.Name)
	}
	idx, err := helper_ui.Select("", items)
	if err != nil {
		return nil, err
	}
	device := deviceList[idx]

	return device, nil
}

type Device struct {
	Name string
	URL  string
}

func (r *CommandToTv) loadDevice() ([]*Device, error) {
	deviceList, err := devices.LoadSSDPservices(1)
	if err != nil {
		return nil, err
	}
	res := []*Device{}
	for k, v := range deviceList {
		res = append(res, &Device{
			Name: k,
			URL:  v,
		})
	}
	sort.SliceIsSorted(res, func(i, j int) bool {
		return res[i].Name < res[j].Name
	})
	return res, nil
}
