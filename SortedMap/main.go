package main

import (
    "bufio"
    "encoding/json"
    "fmt"
    "log"
    "os"
    "sort"
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
        jsonString := scanner.Bytes()

        var target []jsonStruct

        err := json.Unmarshal(jsonString, &target)
        if err != nil {
            fmt.Println(err)
            return
        }

        m := make(map[int]string)
        for _, t := range target {
            m[t.Number] = t.Text
        }
        text := sortedMapIntString(m)

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

func sortedMapIntString(m map[int]string) string {
    var text string
    var keys []int
    for k, _ := range m {
        keys = append(keys, k)
    }
    sort.Ints(keys)

    for _, k := range keys {
        text += m[k]
    }

    return text
}