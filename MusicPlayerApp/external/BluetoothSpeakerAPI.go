package external

import "fmt"

type BluetoothSpeakerAPI struct{}

func (b *BluetoothSpeakerAPI) PlaySoundViaBluetooth(data string) {
	fmt.Println("[BluetoothSpeaker] Playing:", data)
}
