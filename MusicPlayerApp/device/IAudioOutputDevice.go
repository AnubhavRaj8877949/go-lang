package device

import "food-delivery/models"

type IAudioOutputDevice interface {
	PlayAudio(song models.Song)
}
