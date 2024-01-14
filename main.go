package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	var (
		c = flag.Bool("c", false, "print the byte counts")
		l = flag.Bool("l", false, "print the newline counts")
		w = flag.Bool("w", false, "print the word counts")
		m = flag.Bool("m", false, "print the character counts")
	)
	flag.Parse()
	filePath := flag.Arg(0)
	var (
		data []byte
		err  error
	)
	var newlines int

	if filePath != "" {
		data, err = os.ReadFile(filePath)
	} else {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			data = append(data, scanner.Bytes()...)
			data = append(data, []byte{'\r', '\n'}...)
			newlines += 1
		}
		err = scanner.Err()
	}
	if err != nil {
		panic(err)
	}

	var byteCounts int
	if *c {
		byteCounts = getBytes(&data)
	}

	if *l && filePath != "" {
		newlines = getNewlines(&data)
	}

	var words int
	if *w {
		words = getWords(&data)
	}

	var chars int
	if *m {
		chars = getChars(&data)
	}

	if !*w && !*m && !*c && !*l {
		if filePath != "" {
			newlines = getNewlines(&data)
		}
		words = getWords(&data)
		byteCounts = getBytes(&data)
		chars = getChars(&data)
	}

	printValues(newlines, words, byteCounts, chars, *l, *w, *c, *m, filePath)
}

func getNewlines(data *[]byte) int {
	lineSep := []byte{'\n'}
	var newlines int
	for _, b := range *data {
		if lineSep[0] == b {
			newlines += 1
		}
	}
	return newlines
}

func getWords(data *[]byte) int {
	r := strings.NewReader(string(*data))
	s := bufio.NewScanner(r)
	var words int
	s.Split(bufio.ScanWords)
	for s.Scan() {
		words += 1
	}
	return words
}

func getChars(data *[]byte) int {
	r := strings.NewReader(string(*data))
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanRunes)
	var chars int
	for s.Scan() {
		chars += 1
	}
	return chars
}

func getBytes(data *[]byte) int {
	return len(*data)
}

func printValues(newlines, words, byteCounts, chars int, l, w, c, m bool, filePath string) {
	head := ""
	out := ""
	if l {
		head += fmt.Sprintf("%15s", "newlines")
		out = fmt.Sprintf("%15d ", newlines)
	}
	if w {
		head += fmt.Sprintf("%15s", "words")
		out += fmt.Sprintf("%15d ", words)
	}
	if c {
		head += fmt.Sprintf("%15s", "bytes")
		out += fmt.Sprintf("%15d ", byteCounts)
	}
	if m {
		head += fmt.Sprintf("%15s", "chars")
		out += fmt.Sprintf("%15d ", chars)
	}
	if !w && !m && !c && !l {
		head += fmt.Sprintf("%15s %15s %15s %15s", "newlines", "words", "bytes", "chars")
		out += fmt.Sprintf("%15d %15d %15d %15d ", newlines, words, byteCounts, chars)
	}
	out += fmt.Sprintf("%10s", filePath)
	fmt.Println(head)
	fmt.Println(out)
}
