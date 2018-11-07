/**
*
* qlean - a SQL pretty printer tool
* Copyright 2018 by Kyle Pekosh
*
*
 */

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
	fmtKeyword := strings.ToUpper(strings.TrimSpace(keyword))
	var prefix string
	switch fmtKeyword {
	case "AS", "ASC", "DESC", "IN", "LIKE", "ON", "USING":
		prefix = " "
	default:
		prefix = "\n"
	}
	return prefix + fmtKeyword + " "
}

func qleanSQL(query string) string {
	/**
	 *
	 */
	kwarray := []string{
		"WITH ",
		"SELECT",
		"FROM",
		"AS",
		"(?:(?:LEFT |RIGHT |FULL )?(OUTER |INNER )?JOIN)",
		"ON",
		"USING",
		"WHERE",
		"AND",
		"OR",
		"IN",
		"LIKE",
		"GROUP BY",
		"HAVING",
		"ORDER BY",
		"WINDOW",
		"LIMIT",
		"OFFSET",
		"ASC",
		"DESC"}

	kwstring := strings.Join(kwarray, "|")
	// Add case insensitive searching for keywords
	kwregex := regexp.MustCompile("(?i)[\r\n\t\f\v ()](" + kwstring + ")[\r\n\t\f\v ()]")
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
	var output string

	for {
		input, err := reader.ReadString(';')
		if err != nil && err == io.EOF {
			break
		}
		output += qleanSQL(input)
	}

	fmt.Printf("%s", output)
}
