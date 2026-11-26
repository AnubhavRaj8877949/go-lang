package device

import (
	"food-delivery/external"
	"food-delivery/models"
)

type HeadphonesAdapter struct {
	headphonesAPI *external.HeadphonesAPI
}

func NewHeadphonesAdapter(h *external.HeadphonesAPI) *HeadphonesAdapter {
	return &HeadphonesAdapter{
		headphonesAPI: h,
	}
}

func (b *HeadphonesAdapter) PlayAudio(song models.Song) {
	b.headphonesAPI.PlaySoundViaJack(song.FilePath)
}
