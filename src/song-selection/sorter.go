package main

import "sort"

// Custom comparer for songs
type ByLikes Playlist

// Lenght of the songs
func (playlist ByLikes) Len() int {
	return len(playlist.songs)
}

//
func (playlist ByLikes) Swap(a, b int) {
	playlist.songs[a], playlist.songs[b] = playlist.songs[b], playlist.songs[a]
}

func (playlist ByLikes) Less(a, b int) bool {
	return playlist.songs[a].likes > playlist.songs[b].likes
}

func (playlist *Playlist) sort() {
	sort.Sort(ByLikes(*playlist))
}
