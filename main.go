package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "regexp"
    "strings"
    "unicode/utf8"
)

type SortedMap struct {
    words []string
    wordCounter map[string]int
}

func (s *SortedMap) WordCounter(word string)  {
    if _, ok := s.wordCounter[word]; ok {
        s.wordCounter[word]++
    } else {
        s.wordCounter[word] = 1
        s.words = append(s.words, word)
    }
}

func (s *SortedMap) GetWordStats(count int) map[string]int  {
    stats := make(map[string]int)

    for i:=0; i < count; i++ {
        word, max := s.GetMaxWordCount()

        if _, ok := stats[word]; ! ok {
            stats[word] = max
            delete(s.wordCounter, word)
        }
    }

    return stats
}

func (s *SortedMap) GetMaxWordCount() (string, int) {
    var word string
    var max int

    for v, i := range s.wordCounter {
        if max < i {
            max = i
            word = v
        }
    }

    return word, max
}

const (
    charWordLimit = 3
    wordCount = 10
)

func main()  {
	scanner := bufio.NewScanner(os.Stdin)

	bufSize := 1024 * 1024 * 5 // 5 Mb
	buf := make([]byte, 0, bufSize)
	scanner.Buffer(buf, bufSize)

    reg := regexp.MustCompile("[^a-zA-Z0-9\\s]+")

    sm := SortedMap{
      wordCounter: make(map[string]int),
    }

	for scanner.Scan() {
		text := scanner.Text()

		for _, line := range strings.Split(text, ". ") {
			if line == "" {
				continue
			}

          processedString := reg.ReplaceAllString(strings.ToLower(line), "")
          words := strings.Split(processedString, " ")

          for i, word := range words {
              if utf8.RuneCountInString(word) <= charWordLimit || i == 0 || i == len(words)-1 {
                  continue
              }

              sm.WordCounter(word)
          }

		}
	}

    wordStats := sm.GetWordStats(wordCount)

    for _, v := range sm.words {
      if _, ok := wordStats[v]; ok {
          fmt.Printf("Word '%s' was repeated %d times \n", v, wordStats[v])
      }
    }

	if err := scanner.Err(); err != nil {
		log.Printf("%v", err)
	}
}