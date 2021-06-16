package main

import (
    "bufio"
    "encoding/json"
    "fmt"
    "log"
    "os"
    "sortedmap/pkg/sortedmap"
)

const (
	wordLengthLimit = 3
	wordCount     = 10
)

type jsonStruct struct {
    Number int
    Text string
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	bufSize := 1024 * 1024 * 5 // 5 Mb
	buf := make([]byte, 0, bufSize)
	scanner.Buffer(buf, bufSize)

	sm := sortedmap.New()

	for scanner.Scan() {
        jsonString := scanner.Text()

        var target []jsonStruct

        err := json.Unmarshal([]byte(jsonString), &target)
        if err != nil {
            fmt.Println(err)
            return
        }

        var text string
        for _, t := range target {
            text += t.Text
        }

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