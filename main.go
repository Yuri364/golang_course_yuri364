package main

import (
	"bufio"
	"fmt"
	"github.com/Yuri364/golang_course_yuri364/sortedmap/pkg/sortedmap"
	"log"
	"os"
	"regexp"
	"strings"
	"unicode/utf8"
)

const (
	charWordLimit = 3
	wordCount     = 10
)

var reg = regexp.MustCompile(`[^a-zA-Z0-9\s]+`)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	bufSize := 1024 * 1024 * 5 // 5 Mb
	buf := make([]byte, 0, bufSize)
	scanner.Buffer(buf, bufSize)

	sm := sortedmap.New()

	for scanner.Scan() {
		text := scanner.Text()

		for _, line := range strings.Split(text, ". ") {
			if line == "" {
				continue
			}
			processedString := reg.ReplaceAllString(strings.ToLower(line), "")
			words := strings.Split(processedString, " ")

			for i, word := range words {
				word = strings.ReplaceAll(word, "\t", "")

				if utf8.RuneCountInString(word) <= charWordLimit {
					continue
				}

				if i == 0 || i == len(words)-1 {
					sm.AddToStopLIst(word)

					continue
				}

				sm.WordCounter(word)
			}

		}
	}

	wordStats := sm.GetWordStats(wordCount)

	for _, v := range sm.Words {
		if _, ok := wordStats[v]; ok {
			fmt.Printf("Word %q was repeated %d times \n", v, wordStats[v])
		}
	}

	if err := scanner.Err(); err != nil {
		log.Printf("%v", err)
	}
}