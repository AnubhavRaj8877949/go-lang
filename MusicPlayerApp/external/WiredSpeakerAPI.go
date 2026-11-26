package external

import "fmt"

type WiredSpeakerAPI struct{}

func (b *WiredSpeakerAPI) PlaySoundViaCable(data string) {
	fmt.Println("[WiredSpeakerAPI] Playing:", data)
}
