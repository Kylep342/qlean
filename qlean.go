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
	// TODO: Write a function to call .upper on any keyword and pad with \n & ' '
	// do this with ReplaceAllFunc
	return kwregex.ReplaceAllString(spacedsql, "\n$1 ")
}

func main() {
	_, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}

	/**
		 *
		 * handling meant to switch between piping and calling on files
		 * does not currenlty work
	   *
	*/
	// var reader io.Reader
	// if info.Mode()&os.ModeCharDevice != 0 {
	// 	reader = os.Stdin
	// } else {
	// 	fpath := "./" + os.Args[1]
	// 	reader, err = os.Open(fpath)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 		os.Exit(1)
	// 	}
	// }

	// consolereader := bufio.NewReader(reader)
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
