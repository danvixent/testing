package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
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
				booksAndScores[id] = score
			}

			nxtID := -1
			for i := 2; i < len(lines); i = i + 2 {
				tmp := strings.Split(lines[i], " ")
				nxtID++
				struc := &library{}
				struc.ID = nxtID
				struc.SignUpTime, _ = strconv.Atoi(tmp[1])
				struc.ScansPerDay, _ = strconv.Atoi(tmp[2])
				struc.BookIDs = extract(strings.Split(lines[i+1], " "))
				struc.calcQuality()
				allLibs = append(allLibs, struc)
			}
			sortLibs()
			for _, lib := range allLibs {
				procLibs(lib)
			}
			printToFile(path)
			clearDataStructures()
		}
		return nil
	})
	stp := time.Since(str).Seconds()
	fmt.Println("Time:", stp)
}

func procLibs(lib *library) {
	// time.Sleep(10 * time.Millisecond) //do not remove used to curtail speed of goroutines
	tmp := 0

	lib.signUp()
	tmp = days

	lib.scanBooks(tmp)
	fmt.Println("Libary", lib.ID, "finished the process")

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
