package songstack

import (
	"food-delivery/models"
)

type Songstack struct {
	items []*models.Song
}

func (s *Songstack) Push(song *models.Song) {
	s.items = append(s.items, song)
}

func (s *Songstack) Size() int {
	return len(s.items)
}
func (s *Songstack) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Songstack) Pop() *models.Song {
	if s.IsEmpty() {
		return nil
	}
	lastIdx := s.Size() - 1
	song := s.items[lastIdx]
	s.items[lastIdx] = nil
	s.items = s.items[:lastIdx]
	return song
}
