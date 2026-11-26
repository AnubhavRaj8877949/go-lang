package strategies

import "food-delivery/models"

type CustomQueueStrategy struct {
	playlist       models.Playlist
	currentSongIdx int
}

func NewCustomQueueStrategy() *CustomQueueStrategy {
	return &CustomQueueStrategy{
		currentSongIdx: 0,
	}
}

func (cu *CustomQueueStrategy) SetPlaylist(playlist *models.Playlist) {
	cu.playlist = *playlist
}

func (cu *CustomQueueStrategy) HasNext() bool {
	if cu.currentSongIdx < cu.playlist.GetSize() {
		return true
	}
	return false
}

func (cu *CustomQueueStrategy) Next() *models.Song {
	if !cu.HasNext() {
		return nil
	}
	cu.currentSongIdx++
	return cu.playlist.SongList[cu.currentSongIdx]
}

func (cu *CustomQueueStrategy) HasPrevious() bool {
	if cu.playlist.SongList != nil && cu.currentSongIdx > 0 && cu.playlist.GetSize() > 0 {
		return true
	}
	return false

}

func (cu *CustomQueueStrategy) Previous() *models.Song {
	if !cu.HasPrevious() {
		return nil
	}
	cu.currentSongIdx--
	return cu.playlist.SongList[cu.currentSongIdx]

}

func (cu *CustomQueueStrategy) AddToNext(song *models.Song) {
	insertAt := cu.currentSongIdx + 1
	cu.playlist.SongList = append(cu.playlist.SongList[:insertAt], append([]*models.Song{song}, cu.playlist.SongList[insertAt:]...)...)
}
