package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	selection := []Song{
		{
			title:    "Metallica - One",
			likes:    0,
			category: rock,
		},
		{
			title:    "Lenny Kravitz - Fly Away",
			likes:    0,
			category: rock,
		},
		{
			title:    "Chillout Lounge",
			likes:    0,
			category: jazz,
		},
		{
			title:    "Miles Davis - The Doo-Bop Song",
			likes:    0,
			category: jazz,
		},
		{
			title:    "Michael Jackson - Billie Jean",
			likes:    0,
			category: pop,
		},
		{
			title:    "Madonna - Frozen",
			likes:    0,
			category: pop,
		},
	}

	playlist := New_playlist(selection)

	playlist.print()

	scanner := bufio.NewScanner(os.Stdin)

	ticker := time.NewTicker(50 * time.Millisecond)

	for range ticker.C {
		next := playlist.get_next()
		playlist.play(next)

		for scanner.Scan() {
			if scanner.Text() == "x" {
				fmt.Println("Closing player...")
				os.Exit(0)
			}

			choice, _ := strconv.ParseInt(scanner.Text(), 0, 64)
			next.set_likes(choice == 1) // true/false
			break
		}
		playlist.sort()
		playlist.print()
	}
}
