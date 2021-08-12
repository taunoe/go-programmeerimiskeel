package main

import "fmt"

type Move struct {
	data string
	next *Move // Pointer to same type
}

func NewMove(data string, next *Move) *Move {
	return &Move{data: data, next: next}
}

// Linked list storing chess moves
type ChessMatch struct {
	head *Move
}

//
func NewChessMatch() *ChessMatch {
	c := ChessMatch{}
	return &c
}

func (c *ChessMatch) getAt(index int) *Move {
	currentElm := c.head
	pos := 0

	for pos < index && currentElm != nil {
		currentElm = currentElm.next
		pos++
	}

	return currentElm
}

func (c *ChessMatch) getLast() *Move {
	if c.head == nil {
		return nil
	}

	current := c.head

	for current != nil && current.next != nil {
		current = current.next
	}
	return current
}

func (c *ChessMatch) insertAt(index int, data string) {
	if c.head == nil {
		c.head = NewMove(data, nil)
		return
	}

	if index == 0 {
		h := c.head
		c.head = NewMove(data, h)
		return
	}

	prev := c.getAt(index - 1)
	if prev == nil {
		prev = c.getLast()
	}
	prev.next = NewMove(data, prev.next)
}

func (c *ChessMatch) removeAt(index int) {
	if c.head == nil {
		return
	}

	if index == 0 {
		c.head = c.head.next
		return
	}

	prev := c.getAt(index - 1)
	if prev != nil && prev.next != nil {
		prev.next = prev.next.next
	}
}

func (c *ChessMatch) forEach(predicateFn func(*Move)) {
	if c.head == nil {
		return
	}

	current := c.head
	for current != nil {
		predicateFn(current)
		current = current.next
	}
}

func (c *ChessMatch) replayMatch() {
	printMove := func(t *Move) { fmt.Println(t.data) }
	c.forEach(printMove)
}
