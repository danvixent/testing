package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	// create output directory
	err := os.Mkdir("outputs", os.ModePerm)
	if err != nil {
		fmt.Printf("Could not create outputs directory: %v\n", err)
		os.Exit(3)
	}
	// walk through the input directory
	str := time.Now()
	filepath.Walk("inputs", func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()
			data, _ := ioutil.ReadAll(file)
			splitted := bytes.Split(data, []byte("\n"))

			var lines []string
			cha := ""
			for i, b := range splitted {
				for _, c := range b {
					cha += string(c)
				}
				if i < len(splitted)-1 {
					lines = append(lines, cha)
					cha = ""
				}
			}

			alpha = *extract(strings.Split(lines[0], " "))
			days = alpha[2]                         //package level variable
			allLibs = make([]*library, 0, alpha[1]) //eliminate making calls to append() reallocate

			id := -1
			tmp := extract(strings.Split(lines[1], " "))
			for _, score := range *tmp {
				id++
				allBooks[id] = &book{
					ID:    id,
					Score: score,
				}
			}
			nxtID := -1
			for i := 2; i < len(lines); i = i + 2 {
				tmp := strings.Split(lines[i], " ")
				nxtID++
				struc := &library{}
				struc.ID = nxtID
				struc.SignUpTime, _ = strconv.Atoi(tmp[1])
				struc.ScansPerDay, _ = strconv.Atoi(tmp[2])

				noOfBooks, _ := strconv.Atoi(tmp[0])
				struc.Books = make([]*book, 0, noOfBooks)
				bks := extract(strings.Split(lines[i+1], " "))
				for _, id := range *bks {
					if !kept[id] {
						kept[id] = true
						struc.Books = append(struc.Books, allBooks[id])
					}
				}

				struc.NoOfBooks = len(struc.Books)
				struc.sortBooks()
				struc.ScoreLib()
				allLibs = append(allLibs, struc)
			}

			sort.SliceStable(allLibs, func(i, j int) bool {
				return allLibs[i].Score > allLibs[j].Score
			})
			for i := 0; i < len(allLibs); i++ {
				procLibs(allLibs[i])
			}
			printToFile(path)
		}
		return nil
	})
	stp := time.Since(str).Seconds()
	fmt.Println("Time:", stp)
}

func procLibs(lib *library) {
	if lib.Score > 0 {
		lib.signUp()
		lib.scanBooks(days)
		// fmt.Println("Libary", lib.ID, "finished the process")
	}
}

//Print needed output to file
func printToFile(path string) {
	output := ""
	noLib := 0
	for _, lib := range allLibs {
		if lib.ScannedBooks == nil || len(*lib.ScannedBooks) == 0 {
			continue
		}
		output += strconv.Itoa(lib.ID) + " " + strconv.Itoa(len(*lib.ScannedBooks)) + "\n"
		noLib++
		for _, id := range *lib.ScannedBooks {
			output += strconv.Itoa(id) + " "
		}
		output += "\n"
	}
	output = strconv.Itoa(noLib) + "\n" + output

	outFile := strings.Trim(path, "inputs")
	outFile = outFile[1:2] + ".out"
	outFile = "outputs/" + outFile
	f, err := os.Create(outFile)
	defer f.Close()
	if err != nil {
		fmt.Println("Cannot create output file: ", err)
		return // return since there is no file to write to
	}

	_, err = f.Write([]byte(output))
	if err != nil {
		fmt.Println("Cannot write output to file: ", err)
	}
	f.Sync()
}
