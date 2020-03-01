package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
)

// extract returns the integer equivalents of numbers in the slice parameter...translated into a slice of ints
func extract(slice []string) *[]int {
	var tmp []int
	for ix := range slice {
		ref, err := strconv.Atoi(slice[ix])
		if err != nil {
			fmt.Println("Conversion failed")
			os.Exit(3) // we make this stringent because this error should never occur
		}
		tmp = append(tmp, ref)
	}
	return &tmp
}

func shipBooks(books *[]*book, max int) *[]int {
	if max == -1 {
		tmp := make([]int, 0)
		for _, book := range *books {
			if !seen[book.ID] {
				seen[book.ID] = true
				tmp = append(tmp, book.ID)
			}
		}
		return &tmp
	}

	tmp := make([]int, 0, len(*books)/2) //eliminate slice re-allocations to my best ability
	for _, book := range *books {
		if max > 0 { //ship possibly to the maximum number of books allowed
			//i placed this in this loop because it number of books might be less than the maximum given
			if !seen[book.ID] {
				seen[book.ID] = true
				tmp = append(tmp, book.ID)
				max--
			}
			continue
		}
		break
	}
	return &tmp
}

func sortLibs() {
	sort.SliceStable(allLibs, func(i, j int) bool {
		return allLibs[i].SignUpTime < allLibs[j].SignUpTime
	})
}
