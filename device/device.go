// TODO : Documentation

package device

import (
	"errors"
	"fmt"
	"github.com/kunaldawn/goandroid/adbutility"
	"strings"
)

type Device struct {
	Serial  string // Device serial number
	Timeout int    // Timeout in seconds for all adb and shell operations
}

func NewDevice(serial string, timeout int) Device {
	return Device{Serial: serial, Timeout: timeout}
}

func (dev Device) IsAvailable() (bool, error) {
	devices, err := adbutility.GetAttachedDevices(dev.Timeout)
	if err != nil {
		return false, err
	}
	for index := range devices {
		if dev.Serial == devices[index] {
			return true, nil
		}
	}
	return false, nil
}

func (dev Device) Adb(command string, args ...string) (string, error) {
	return adbutility.Adb(dev.Timeout, append([]string{"-s", dev.Serial, command}, args...)...)
}

func (dev Device) Shell(command string, args ...string) (string, error) {
	return dev.Adb("shell", append([]string{command}, args...)...)
}

func (dev Device) GetProperty(key string) (string, error) {
	prop, err := dev.GetAllProperties()
	if err != nil {
		return "", err
	}
	val, ok := prop[key]
	if !ok {
		return "", errors.New(fmt.Sprintf("Key [%s] is not found in device properties", key))
	}
	return val, nil
}

func (dev Device) GetAllProperties() (map[string]string, error) {
	prop_map := make(map[string]string)
	prop, err := dev.Shell("getprop")
	if err != nil {
		return prop_map, err
	}
	lines := strings.Split(prop, "\n")
	for index := range lines {
		parts := strings.Split(lines[index], ":")
		if len(parts) == 2 {
			key := strings.TrimSpace(strings.Replace(strings.Replace(parts[0], "[", "", -1), "]", "", -1))
			value := strings.TrimSpace(strings.Replace(strings.Replace(parts[1], "[", "", -1), "]", "", -1))
			prop_map[key] = value
		}
	}
	return prop_map, nil
}

func (dev Device) Pull(src string, dst string) (string, error) {
	return dev.Adb("pull", src, dst)
}

func (dev Device) Push(src string, dst string) (string, error) {
	return dev.Adb("push", src, dst)
}

func (dev Device) WaitForAvailability() (string, error) {
	return dev.Adb("wait-for-device")
}

func (dev Device) Root() (string, error) {
	return dev.Adb("root")
}