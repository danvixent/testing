package main

import (
	"sync"
)

var signup sync.Mutex
var out sync.Mutex
var see sync.Mutex

var days = 0
var allLibs []*library
var booksAndScores = make(map[int]int)
var numOfLibsToShipFrom = 0
var alpha []int
var seen = make(map[int]bool)
var taken = make(map[int]bool)
var wait sync.WaitGroup

// describes each library
type library struct {
	ID           int
	SignUpTime   int
	ScansPerDay  int
	ScannedBooks *[]int
	BookIDs      *[]int
	IsSignedUp   bool
	Quality      float64
}

func (l *library) calcQuality() {
	//Here we define different criteria for determining which libaries to process first
	//This is the first approach i took, it's quite verbose so read and understand
	// {
	// 	x := float64(len(*l.BookIDs) / l.ScansPerDay)
	// 	tmp := (x / float64(l.SignUpTime)) * l.avgBookScore()
	// 	l.Quality = tmp
	// }

	//for input b specifically, since all book have the same score and all libaries have the same number of books
	//, the concerns should be sign up itme and total score

	q := float64(l.SignUpTime)
	l.Quality = q
}

func (l *library) avgBookScore() float64 {
	scores := 0.0
	for _, id := range *l.BookIDs {
		scores += float64(booksAndScores[id])
	}
	return scores / float64(len(*l.BookIDs))
}

func (l *library) signUp() {
	l.IsSignedUp = true
	days = days - l.SignUpTime
}

func (l *library) scanBooks(shippingDays int) {
	// l.sortBooksByScore()
	maxBooksToShip := shippingDays * l.ScansPerDay
	if maxBooksToShip < len(*l.BookIDs) {
		l.ScannedBooks = shipBooks(l.BookIDs, maxBooksToShip)
		return
	}
	l.ScannedBooks = shipBooks(l.BookIDs, -1)
}

func (l *library) totalScore() int {
	score := 0
	for _, id := range *l.BookIDs {
		if !taken[id] { //keep track of books we've encontered
			taken[id] = true
			score += booksAndScores[id]
		}
	}
	return score
}
