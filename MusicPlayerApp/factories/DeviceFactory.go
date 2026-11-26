package factories

import (
	"fmt"
	"food-delivery/device"
	"food-delivery/enums"
	"food-delivery/external"
)

type DeviceFactory interface {
	CreateDevice(dt enums.DeviceType) (device.IAudioOutputDevice, error)
}

type DefaultDeviceFactory struct{}

func (DefaultDeviceFactory) CreateDevice(dt enums.DeviceType) (device.IAudioOutputDevice, error) {
	switch dt {
	case enums.BLUETOOTH:
		api := &external.BluetoothSpeakerAPI{}
		return device.NewBluetoothSpeakerAdapter(api), nil
	case enums.WIRED:
		api := &external.WiredSpeakerAPI{}
		return device.NewWiredSpeakerAdapter(api), nil
	case enums.HEADPHONES:
		api := &external.HeadphonesAPI{}
		return device.NewHeadphonesAdapter(api), nil
	default:
		return nil, fmt.Errorf("unsupported device type: %q", dt)

	}

}
