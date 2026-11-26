package device

import (
	"food-delivery/external"
	"food-delivery/models"
)

type WiredSpeakerAdapter struct {
	wiredSpeakerAPI *external.WiredSpeakerAPI
}

func NewWiredSpeakerAdapter(h *external.WiredSpeakerAPI) *WiredSpeakerAdapter {
	return &WiredSpeakerAdapter{
		wiredSpeakerAPI: h,
	}
}

func (b *WiredSpeakerAdapter) PlayAudio(song models.Song) {
	b.wiredSpeakerAPI.PlaySoundViaCable(song.FilePath)
}
