package strategies

import "food-delivery/models"

type SequentialPlayStrategy struct {
	playlist       models.Playlist
	currentSongIdx int
}

func NewSequentialPlayStrategy() *SequentialPlayStrategy {
	return &SequentialPlayStrategy{
		currentSongIdx: 0,
	}
}

func (sq *SequentialPlayStrategy) SetPlaylist(playlist *models.Playlist) {
	sq.playlist = *playlist
}

func (sq *SequentialPlayStrategy) HasNext() bool {
	if sq.currentSongIdx < sq.playlist.GetSize() {
		return true
	}
	return false
}

func (sq *SequentialPlayStrategy) Next() *models.Song {
	if !sq.HasNext() {
		return nil
	}
	sq.currentSongIdx++
	return sq.playlist.SongList[sq.currentSongIdx]
}

func (sq *SequentialPlayStrategy) HasPrevious() bool {
	if sq.playlist.SongList != nil && sq.currentSongIdx > 0 && sq.playlist.GetSize() > 0 {
		return true
	}
	return false

}

func (sq *SequentialPlayStrategy) Previous() *models.Song {
	if !sq.HasPrevious() {
		return nil
	}
	sq.currentSongIdx--
	return sq.playlist.SongList[sq.currentSongIdx]

}
