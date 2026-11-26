package device

import (
	"food-delivery/external"
	"food-delivery/models"
)

type BluetoothSpeakerAdapter struct {
	bluetoothApi *external.BluetoothSpeakerAPI
}

func NewBluetoothSpeakerAdapter(b *external.BluetoothSpeakerAPI) *BluetoothSpeakerAdapter {
	return &BluetoothSpeakerAdapter{
		bluetoothApi: b,
	}
}

func (b *BluetoothSpeakerAdapter) PlayAudio(song models.Song) {
	b.bluetoothApi.PlaySoundViaBluetooth(song.FilePath)
}
