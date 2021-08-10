package main

import "fmt"

// Category enum
type Category int

const (
	pop  Category = iota // 0
	jazz                 // 1
	rock                 // 2
)

// Song type
type Song struct {
	title    string
	category Category //string
	likes    int
}

func (c Category) to_string() string {
	switch c {
	case jazz:
		return "jazz"
	case pop:
		return "pop"
	case rock:
		return "rock"
	default:
		return ""
	}
}

func (song *Song) print() {
	msg := fmt.Sprintf("%s (%s), likes: %d", song.title, song.category.to_string(), song.likes)
	fmt.Println(msg)
}

func (song *Song) set_likes(liked bool) {
	if liked {
		song.likes++
	} else {
		if song.likes > 0 {
			song.likes--
		}
	}
}
