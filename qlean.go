package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
)

// used to capitalize SQL keywords
func formatKeyword(keyword string) string {
	return "\n" + strings.ToUpper(keyword)
}

func qleanSQL(query string) string {
	/**
	 *
	 *
	 *
	 *
	 */
	kwarray := []string{
		"WITH",
		"SELECT",
		"FROM",
		"(?:(?:LEFT |RIGHT |FULL )?(:OUTER |INNER )?JOIN)",
		"WHERE",
		"GROUP",
		"HAVING",
		"ORDER",
		"WINDOW"}

	kwstring := strings.Join(kwarray, "|")
	// Add case insensitive searching for keywords
	kwregex := regexp.MustCompile("(?i)(" + kwstring + ")\\s+")
	whtspregex := regexp.MustCompile("\\s+")
	//
	spacedsql := whtspregex.ReplaceAllString(query, " ")
	return strings.TrimSpace(kwregex.ReplaceAllStringFunc(spacedsql, formatKeyword))
}

func main() {
	info, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}

	var source io.Reader
	if info.Mode()&os.ModeCharDevice == 0 {
		source = os.Stdin
	} else {
		fpath := "./" + os.Args[1]
		source, err = os.Open(fpath)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
	}

	reader := bufio.NewReader(source)
	var output []string

	for {
		input, err := reader.ReadString(';')
		if err != nil && err == io.EOF {
			break
		}
		output = append(output, qleanSQL(input))
	}

	for j := 0; j < len(output); j++ {
		fmt.Printf("%s", output[j])
	}
}
