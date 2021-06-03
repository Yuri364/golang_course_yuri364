package sortedmap

type SortedMap struct {
    Words []string
    wordCounter map[string]int
    stopList map[string]struct{}
}

func New() *SortedMap {
    return &SortedMap{
        wordCounter: make(map[string]int),
        stopList: make(map[string]struct{}),
    }
}

func (s *SortedMap) AddToStopLIst(word string) {
    if _, ok := s.stopList[word]; ! ok {
        s.stopList[word] = struct{}{}
    }
}

func (s *SortedMap) WordCounter(word string) {
    if _, ok := s.wordCounter[word]; ok {
        s.wordCounter[word]++
    } else {
        s.wordCounter[word] = 1
        s.Words = append(s.Words, word)
    }
}

func (s *SortedMap) GetWordStats(count int) map[string]int {
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
        if _, ok := s.stopList[v]; ! ok && max < i {
            max = i
            word = v
        }
    }

    return word, max
}