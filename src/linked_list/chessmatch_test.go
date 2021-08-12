package main

import "testing"

func TestGetAtRandomPosition(t *testing.T) {
	c := NewChessMatch()

	mData := "e4"
	m2Data := "Nf6"

	c.insertAt(0, mData)
	c.insertAt(1, m2Data)

	actual := c.getAt(1)

	if actual == nil || actual.data != m2Data {
		t.Errorf("expected element %s at position 1.", m2Data)
	}
}

func TestGetAtHeadPosition(t *testing.T) {
	c := NewChessMatch()

	mData := "e4"
	m2Data := "Nf6"

	c.insertAt(0, mData)
	c.insertAt(1, m2Data)

	actual := c.getAt(0)

	if actual == nil || actual.data != mData {
		t.Errorf("expected element %s at position 0.", mData)
	}
}

func TestGetAtOutOfBounds(t *testing.T) {
	c := NewChessMatch()

	mData := "e4"

	c.insertAt(0, mData)

	actual := c.getAt(10)

	if actual != nil {
		t.Errorf("expected position 10 return nil")
	}
}

func TestInsertRandomPosition(t *testing.T) {
	c := NewChessMatch()

	mData := "e4"
	m2Data := "Nf6"
	m3Data := "e5"
	m4Data := "Nd5"

	c.insertAt(0, mData)
	c.insertAt(1, m2Data)
	c.insertAt(2, m3Data)
	c.insertAt(3, m4Data)

	actual := c.getAt(2)

	if actual == nil || actual.data != m3Data {
		t.Errorf("expected element %s at position 3.", m3Data)
	}
}

func TestInsertAtHead(t *testing.T) {
	c := NewChessMatch()

	mData := "e4"
	m2Data := "e5"
	c.insertAt(0, mData)
	c.insertAt(0, m2Data)

	if c.head == nil || c.head.data != m2Data {
		t.Errorf("expected element %s at position 2.", m2Data)
	}
}

func TestInsertAtOutOfBounds(t *testing.T) {
	c := NewChessMatch()

	mData := "e4"
	m2Data := "Nf6"
	c.insertAt(0, mData)
	c.insertAt(10, m2Data)

	if c.getAt(1) == nil || c.getAt(1).data != m2Data {
		t.Errorf("expected element %s at position 1.", m2Data)
	}
}

func TestRemoveAtHead(t *testing.T) {
	c := NewChessMatch()

	mData := "e4"
	m2Data := "Nf6"
	c.insertAt(0, mData)
	c.insertAt(1, m2Data)

	c.removeAt(0)

	actual := c.getAt(0)

	if actual == nil || actual.data != m2Data {
		t.Errorf("expected element %s at position 0", m2Data)
	}
}

func TestRemoveAtRandomPosition(t *testing.T) {
	c := NewChessMatch()

	mData := "e4"
	m2Data := "Nf6"
	m3Data := "e5"
	m4Data := "Nd5"

	c.insertAt(0, mData)
	c.insertAt(1, m2Data)
	c.insertAt(2, m3Data)
	c.insertAt(3, m4Data)

	c.removeAt(2)
	actual := c.getAt(2)

	if actual == nil || actual.data != m4Data {
		t.Errorf("expected element %s at position 2", m4Data)
	}
}

func TestRemoveLastPosition(t *testing.T) {
	c := NewChessMatch()

	mData := "e4"
	m2Data := "Nf6"
	m3Data := "e5"
	m4Data := "Nd5"

	c.insertAt(0, mData)
	c.insertAt(1, m2Data)
	c.insertAt(2, m3Data)
	c.insertAt(3, m4Data)

	c.removeAt(3)
	actualLast := c.getLast()

	if actualLast == nil || actualLast.data != m3Data {
		t.Errorf("expected element %s at last position 2", m3Data)
	}
}

func TestRemoveAtOutOfBounds(t *testing.T) {
	c := NewChessMatch()

	mData := "e4"

	c.insertAt(0, mData)

	c.removeAt(20)
}

func TestForeach(t *testing.T) {
	c := NewChessMatch()

	mData := "e4"
	m2Data := "Nf6"
	m3Data := "e5"

	c.insertAt(0, mData)
	c.insertAt(1, m2Data)
	c.insertAt(2, m3Data)

	newData := "modifiedData"
	modifiedDataFn := func(element *Move) { element.data = newData }
	c.forEach(modifiedDataFn)

	actual1 := c.getAt(0)
	actual2 := c.getAt(1)
	actual3 := c.getAt(2)

	if actual1.data != newData || actual2.data != newData || actual3.data != newData {
		t.Errorf("expected all items to contain data %s", newData)
	}
}
