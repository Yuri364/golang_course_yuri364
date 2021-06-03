package main

import (
	"bufio"
	"fmt"
	"github.com/Yuri364/golang_course_yuri364/sortedmap/pkg/sortedmap"
	"log"
	"os"
)

const (
	wordLengthLimit = 3
	wordCount     = 10
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	bufSize := 1024 * 1024 * 5 // 5 Mb
	buf := make([]byte, 0, bufSize)
	scanner.Buffer(buf, bufSize)

	sm := sortedmap.New()

	for scanner.Scan() {
		text := scanner.Text()

		sm.SplitTextIntoWords(text, wordLengthLimit)
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