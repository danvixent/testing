package main

import (
	"fmt"
	"sort"
	"sync"
)

var signup sync.Mutex
var out sync.Mutex
var see sync.Mutex

var days = 0
var allLibs []*library
var allBooks = make(map[int]*book)
var alpha []int
var seen = make(map[int]bool)
var taken = make(map[int]bool)
var kept = make(map[int]bool)

var wait sync.WaitGroup

// describes each library
type library struct {
	ID           int
	SignUpTime   int
	ScansPerDay  int
	ScannedBooks *[]int
	Score        int //
	Books        []*book
	NoOfBooks    int
	IsSignedUp   bool
	Quality      float64
}

type book struct {
	ID    int
	Score int
}

func (l *library) sortBooks() {
	sort.SliceStable(l.Books, func(i, j int) bool {
		return l.Books[i].Score > l.Books[j].Score
	})
}

func (l *library) calcQuality() {
	// q := float64(l.ScoreLib()) / float64(l.SignUpTime)
	// l.Quality = q
}

func (l *library) signUp() {
	l.IsSignedUp = true
	days = days - l.SignUpTime
}

func (l *library) scanBooks(shippingDays int) {
	max := shippingDays * l.ScansPerDay
	l.ScannedBooks = shipBooks(&l.Books, max)
}

func (l *library) ScoreLib() {
	score := 0
	for _, book := range l.Books {
		// if !taken[book.ID] { //keep track of books we've encountered
		// 	taken[book.ID] = true
		score += book.Score
		// }
	}
	fmt.Println("Libary", l.ID, "'s score =", score)
	l.Score = score
}
