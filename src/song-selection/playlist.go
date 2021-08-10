package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Song_list type
type Song_list []*Song

func (songs Song_list) random() *Song {
	if len(songs) == 0 {
		return nil
	}

	// To get random value every time
	rand.Seed(time.Now().UnixNano())
	random_pos := rand.Intn(len(songs))
	return songs[random_pos]
}

// Playlist type
type Playlist struct {
	songs       Song_list
	now_playing *Song
}

// servs as constuctor
func New_playlist(songs []Song) *Playlist {
	list := Playlist{}
	for i := range songs {
		list.songs = append(list.songs, &songs[i])
	}
	return &list
}

func (playlist *Playlist) print() {
	fmt.Println("\nCurrent Playlist is:")

	for _, song := range playlist.songs {
		song.print()
	}
}

func (playlist *Playlist) reset_likes() {
	for _, song := range playlist.songs {
		song.likes = 0
	}
}

func (playlist *Playlist) play(song *Song) {
	playlist.now_playing = song
	msg := fmt.Sprintf("\nNow playing: %s (%s)\npress 1/0 to like/dislike", song.title, song.category.to_string())
	fmt.Println(msg)
}

func (playlist *Playlist) get_next() *Song {
	var other_categ_songs Song_list

	if playlist.now_playing != nil {
		for _, song := range playlist.songs {
			if song.title != playlist.now_playing.title && song.likes < 2 {
				if song.category == playlist.now_playing.category {
					return song
				}
				other_categ_songs = append(other_categ_songs, song)
			}
		}
	}
	// if the current category was listened to twice, try another category
	if len(other_categ_songs) > 0 {
		return other_categ_songs.random()
	}

	playlist.reset_likes()

	return playlist.songs.random()
}
