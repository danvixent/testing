package main

import (
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
var wait sync.WaitGroup

// describes each library
type library struct {
	ID           int
	SignUpTime   int
	ScansPerDay  int
	ScannedBooks *[]int
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
	//for input b specifically, since all book have the same score and all libaries have the same number of books
	//, the concerns should be sign up itme and total score

	q := (float64(l.totalScore()) / float64(l.SignUpTime)) * float64(l.ScansPerDay)
	l.Quality = q
}

// func (l *library) avgBookScore() float64 {
// 	scores := 0.0
// 	for _, id := range *l.BookIDs {
// 		scores += float64(allBooks[id].Score)
// 	}
// 	return scores / float64(len(*l.BookIDs))
// }

func (l *library) signUp() {
	l.IsSignedUp = true
	days = days - l.SignUpTime
}

func (l *library) scanBooks(shippingDays int) {
	// l.sortBooksByScore()
	l.ScannedBooks = shipBooks(&l.Books, -1)
}

func (l *library) totalScore() int {
	score := 0
	for _, book := range l.Books {
		if !taken[book.ID] { //keep track of books we've encountered
			taken[book.ID] = true
			score += book.Score
		}
	}
	return score
}
