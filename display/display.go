package display

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/supaplextor/adbGoAutomation/device"
)

type Display struct {
	dev device.Device
}

func NewDisplay(dev device.Device) Display {
	return Display{dev: dev}
}

func (disp Display) GetDisplaySize() (int, int, error) {
	txt, err := disp.dev.Shell("wm", "size")
	if err != nil {
		return -1, -1, err
	}
	if !strings.Contains(txt, "Physical size:") {
		return -1, -1, errors.New("not able to determine display size")
	}
	size := strings.Split(strings.TrimSpace(txt), "Physical size:")[1]
	//	size := strings.Split(strings.TrimSpace(size), "Physical size:")[1]
	width, err := strconv.Atoi(strings.Split(strings.TrimSpace(size), "x")[0])
	if err != nil {
		return -1, -1, err
	}
	height, err := strconv.Atoi(strings.Split(strings.TrimSpace(size), "x")[1])
	if height == 0 {
		height = 3200 // FIXME BUG TODO
	}
	if err != nil {
		return -1, -1, err
	}
	return width, height, nil
}

func (disp Display) SetDisplaySize(width int, height int) error {
	size := fmt.Sprintf("%dx%d", width, height)
	_, err := disp.dev.Shell("wm", "size", size)
	return err
}
