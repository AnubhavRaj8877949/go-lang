package strategies

import (
	"food-delivery/datastructures/songstack"
	"food-delivery/models"
	"math/rand"
)

type RandomPlayStrategy struct {
	playlist       models.Playlist
	remainingSongs []*models.Song
	playedSong     songstack.Songstack
	currentSong    *models.Song
}

func NewRandomPlayStrategy() *RandomPlayStrategy {
	return &RandomPlayStrategy{
		remainingSongs: []*models.Song{},
	}
}

func (rm *RandomPlayStrategy) SetPlaylist(playlist *models.Playlist) {
	if playlist == nil {
		return
	}

	rm.playlist = *playlist
	rm.playedSong = songstack.Songstack{}
	rm.remainingSongs = make([]*models.Song, len(rm.playlist.SongList))
	rm.currentSong = nil
	copy(rm.remainingSongs, rm.playlist.SongList) //playlist ke sare song  remainingSongs me copy

}

func (rm *RandomPlayStrategy) HasNext() bool {
	if len(rm.remainingSongs) > 0 {
		return true
	}
	return false
}

func (rm *RandomPlayStrategy) Next() *models.Song {
	if !rm.HasNext() {
		return nil
	}
	// 1. Current song ko history stack mein save karein
	rm.playedSong.Push(rm.currentSong)
	//genrate a rendon numaber
	randomIdx := rand.Intn(len(rm.remainingSongs))
	rm.currentSong = rm.playlist.SongList[randomIdx]
	// remove  that  song from  remainingSongs    to prevent  duplicate played O(n) time
	// rm.remainingSongs = append(rm.remainingSongs[randomIdx:], rm.remainingSongs[randomIdx+1:]...)
	// remove  that  song from  remainingSongs    to prevent  duplicate played O(1)
	lastIdx := len(rm.remainingSongs) - 1
	rm.remainingSongs[randomIdx] = rm.remainingSongs[lastIdx]
	rm.remainingSongs = rm.remainingSongs[:lastIdx]

	return rm.currentSong
}

func (rm *RandomPlayStrategy) HasPrevious() bool {
	if !rm.playedSong.IsEmpty() {
		return true
	}
	return false
}

func (rm *RandomPlayStrategy) Previous() *models.Song {
	if !rm.HasPrevious() {
		return nil
	}
	rm.remainingSongs = append(rm.remainingSongs, rm.currentSong)
	rm.currentSong = rm.playedSong.Pop()
	return rm.currentSong

}
