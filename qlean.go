package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

func qleanSQL(query string) string {
	/**
	 *
	 *
	 *
	 *
	 */
	kwarray := []string{
		"SELECT",
		"FROM",
		"(?:(?:LEFT |RIGHT |FULL )?(:OUTER |INNER )?JOIN)",
		"WHERE",
		"GROUP",
		"HAVING",
		"ORDER",
		"WINDOW"}

	kwstring := strings.Join(kwarray, "|")
	//
	kwregex := regexp.MustCompile("(?i)(" + kwstring + ")\\s+")
	whtspregex := regexp.MustCompile("\\s+")
	//
	spacedsql := whtspregex.ReplaceAllString(query, " ")
	return kwregex.ReplaceAllString(spacedsql, "\n$1")
}

func main() {
	_, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}

	consolereader := bufio.NewReader(os.Stdin)
	var output []string

	for {
		input, err := consolereader.ReadString(';')
		if err != nil && err == io.EOF {
			break
		}
		output = append(output, qleanSQL(input))
	}

	for j := 0; j < len(output); j++ {
		fmt.Printf("%s", output[j])
	}
}
