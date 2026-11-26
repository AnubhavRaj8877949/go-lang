package models

type Playlist struct {
	PlaylistName string
	SongList     []*Song
}

func NewPlaylist(PlaylistName string) *Playlist {
	return &Playlist{
		PlaylistName: PlaylistName,
		SongList:     make([]*Song, 0),
	}
}

func (p *Playlist) AddSongToPlaylist(s *Song) {
	p.SongList = append(p.SongList, s)
}

func (p *Playlist) GetSize() int {
	return len(p.SongList)
}

func (p *Playlist) GetSongs() []*Song {
	return p.SongList

}
