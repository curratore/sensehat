package sensehat

import (
	"errors"
	"io/ioutil"
	"path"
	"path/filepath"
	"strings"

	"github.com/mephux/common"
)

// Device data
type Device struct {
	NamePath string
	FDPath   string
	Map      map[int][][]int
	Rotation int
}

// GetDevice returns the sense hat
// device file descriptor
func GetDevice() (*Device, error) {
	device := &Device{}
	device.Map = PixelMap
	device.Rotation = 0

	paths, _ := filepath.Glob("/sys/class/graphics/fb*")

	for _, fd := range paths {
		namePath := path.Join(fd, "name")

		if common.IsExist(namePath) && common.IsFile(namePath) {
			device.NamePath = namePath

			if dat, err := ioutil.ReadFile(namePath); err == nil {
				name := string(dat)

				if strings.Trim(name, "\n") == SenseHatFbName {
					fdPath := strings.Replace(fd, filepath.Dir(fd), "/dev", -1)

					if common.IsExist(fdPath) {
						device.FDPath = fdPath
						return device, nil
					}
				}
			}
		}
	}

	return device, errors.New("device not found")
}
