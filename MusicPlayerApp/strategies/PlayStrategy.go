package strategies

import "food-delivery/models"

type PlayStrategy interface {
	SetPlaylist(playlist *models.Playlist)
	HasNext() bool
	Next() *models.Song
	Previous() *models.Song
	HasPrevious() bool
	AddToNext(song *models.Song)
}
