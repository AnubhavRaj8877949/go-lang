package external

import "fmt"

type HeadphonesAPI struct{}

func (b *HeadphonesAPI) PlaySoundViaJack(data string) {
	fmt.Println("[HeadphonesAPI] Playing:", data)
}
