package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

/*
Usage: go run changeDT.go <filename>:
Example:go run changeDT.go logEntries.txt
*/
func main() {
	arguments := os.Args
	//if len(arguments) == 1 {
	//	fmt.Println("Please provide one text file to process!")
	//	os.Exit(1)
	//}
	if len(arguments) == 1 {
		fmt.Println("Please provide one text file to process!Now using default file logEntries.txt")
		arguments = append(arguments, "/ch04/logEntries.txt")
	}

	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	filename := arguments[1]
	f, err := os.Open(filepath.Join(pwd, filename))
	if err != nil {
		fmt.Printf("error opening file %s", err)
		os.Exit(1)
	}
	defer f.Close()

	notAMatch := 0
	r := bufio.NewReader(f)
	r1 := regexp.MustCompile(`.*\[(\d\d\/\w+/\d\d\d\d:\d\d:\d\d:\d\d.*)\] .*`)
	r2 := regexp.MustCompile(`.*\[(\w+\-\d\d-\d\d:\d\d:\d\d:\d\d.*)\] .*`)
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("error reading file %s", err)
		}

		if r1.MatchString(line) {
			match := r1.FindStringSubmatch(line)
			d1, err := time.Parse("02/Jan/2006:15:04:05 -0700", match[1])
			if err == nil {
				newFormat := d1.Format(time.Stamp)
				fmt.Print(strings.Replace(line, match[1], newFormat, 1))
			} else {
				notAMatch++
			}
			continue
		}

		if r2.MatchString(line) {
			match := r2.FindStringSubmatch(line)
			d1, err := time.Parse("Jan-02-06:15:04:05 -0700", match[1])
			if err == nil {
				newFormat := d1.Format(time.Stamp)
				fmt.Print(strings.Replace(line, match[1], newFormat, 1))
			} else {
				notAMatch++
			}
			continue
		}
	}
	fmt.Println(notAMatch, "lines did not match!")
}
