package cache

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Party struct {
	originalName string
	encodedName  string
	songs        []string
	open         bool
}

func NewParty(name, encodedName string) *Party {
	return &Party{
		name,
		encodedName,
		[]string{},
		false,
	}
}

func (p *Party) PopSong() (string, error) {
	if len(p.songs) > 0 {
		// fmt.Println("before: " + strings.Join(p.songs, " "))
		song := p.songs[0]
		p.songs = p.songs[1:]
		return song, nil
		// fmt.Println("after: " + strings.Join(p.songs, " "))
	} else {
		fmt.Println("nothing to pop")
	}
	return "", nil
}

func (p *Party) RemoveSongByID(songID string) error {
	//TODO make this wayyyy more efficient
	for i, song := range p.songs {
		if song == songID {
			//delete the song
			p.songs = append(p.songs[:i], p.songs[i+1:]...)
			return nil
		}
	}
	return NoSongError{songID}
}

func (p *Party) ClearSongs() error {
	p.songs = []string{}
	return nil
}

func (p *Party) AppendSong(sid string) error {
	p.songs = append(p.songs, sid)
	return nil
}

// func (p *Party) PrependSong(songName, requester string) error {
// p.songs = append([]*song{newSong(songName, requester)}, p.songs...)
// return nil
// }

func (p *Party) GetSongList() []string {
	return p.songs
}

func (p *Party) GetSongsJson() ([]byte, error) {
	return json.Marshal(p.songs)
}

func (p *Party) OriginalName() string {
	return p.originalName
}

func (p *Party) EncodedName() string {
	return p.encodedName
}

func (p *Party) IsOpen() bool {
	return p.open
}
func (p *Party) Open() {
	p.open = true
}
func (p *Party) Close() {
	p.open = false
}

//helpers

func getInitials(fullName string) string {
	names := strings.Split(fullName, " ")
	if len(names) <= 0 {
		return "??"
	}

	var first string
	if len(names[0]) <= 0 {
		first = "?"
	} else {
		first = string(names[0][0])
	}
	var last string
	if len(names[len(names)-1]) <= 0 {
		last = "?"
	} else {
		last = string(names[len(names)-1][0])
	}
	return first + last
}
