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
	var sizes []string = strings.Split(strings.Split(txt, "\n")[0], " ")

	if !strings.Contains(txt, "size:") {
		return -1, -1, errors.New("not able to determine display size")
	}
	width, err := strconv.Atoi(strings.Split(strings.TrimSpace(sizes[2]), "x")[0])
	if err != nil {
		return -1, -1, err
	}
	height, err := strconv.Atoi(strings.Split(strings.TrimSpace(sizes[2]), "x")[1])
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
