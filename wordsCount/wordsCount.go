package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
	"unicode"
	"unicode/utf8"
)

type Pair struct {
	Key   string
	Value int
}

type PairList []Pair

func (p PairList) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
func (p PairList) Len() int {
	return len(p)
}

func (p PairList) Less(i, j int) bool {
	return p[j].Value < p[i].Value
}

// Filt none Letter characters
func FiltNoneLetter(s string) []string {
	noneLetter := func(c rune) bool {
		return !unicode.IsLetter(c)
	}
	return strings.FieldsFunc(s, noneLetter)
}

type WordsCount map[string]int

// Print the result in reverse order
func (wordsCount WordsCount) SortResult() {
	p := make(PairList, len(wordsCount))
	i := 0
	for k, v := range wordsCount {
		p[i] = Pair{k, v}
		i++
	}
	sort.Sort(p) // Sort the pairlist
	// Print
	count := 0
	for _, pair := range p {
		if count == 10 {
			break
		}
		fmt.Printf("%d	%s\n", pair.Value, pair.Key)
		count++
	}
}

// Open the file to read words
func (WordsCount WordsCount) ReadFile(file *os.File) {
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		for _, word := range FiltNoneLetter(strings.TrimSpace(line)) {
			if len(word) > utf8.UTFMax || utf8.RuneCountInString(word) > 0 {
				WordsCount[strings.ToLower(word)] += 1
			}
		}
		if err != nil {
			if err != io.EOF {
				fmt.Println("Failed to finish reading the file:", err)
			}
			break
		}
	}
}

func main() {
	file, err := os.Open("./article.txt")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer file.Close()
	WordsCounter := make(WordsCount)
	WordsCounter.ReadFile(file)
	WordsCounter.SortResult()
	pause()
}

func pause() {
	fmt.Println("Press 1 to exit.")
	var num int
	for {
		fmt.Scanln(&num)
		if num == 1 {
			break
		}
	}
}
