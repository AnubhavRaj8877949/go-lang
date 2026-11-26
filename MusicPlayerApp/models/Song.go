package models

type Song struct {
	Id       string
	Name     string
	Artist   string
	FilePath string
}

// constructor
func NewSong(id, name, artist, filepath string) *Song {
	return &Song{
		Id:       id,
		Name:     name,
		Artist:   artist,
		FilePath: filepath,
	}

}
