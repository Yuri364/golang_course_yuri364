package sortedmap

import (
    "bufio"
    "math/rand"
    "os"
    "reflect"
    "testing"
    "unsafe"
)

type testData struct {
    wordCount int
    words []string
    expected map[string]int
}

func TestSortedMap_GetWordStats(t *testing.T) {

    testData := []testData{
        {
            0,
            []string{"foo","bar"},
            map[string]int {},
        },
        {
            1,
            []string{"foo","bar","foo"},
            map[string]int {"foo":2},
        },
        {
            2,
            []string{"","","string"},
            map[string]int {"":2,"string":1},
        },
        {
            10,
            []string{
                "impolite", "ready", "buzz", "healthy", "holistic","ants","decision","mix",
                "witty", "point", "weigh", "writing", "huge", "slow", "library", "damage",
                "highfalutin", "slimy", "food", "ambitious", "pastoral", "mug", "books", "rabbits", "warm", "tank",
                "embarrassed", "resolute", "stamp", "slimy", "right", "hurry", "pastoral", "motionless",
                "fact", "shelter", "food", "letters", "who", "food", "healthy", "whole", "shaggy", "pastoral", "writing", "help", "finger",
                "wicked", "tank", "story", "tendency", "right", "somber", "scream", "geese", "clap", "playground", "fool", "valuable",
                "crook", "right", "grade", "ambitious", "letters", "lie", "food", "finger", "easy", "terrific", "bit", "disastrous", "defiant", "nosy", "madly",
            },
            map[string]int {
                	"ambitious":2, "finger":2, "food":4, "healthy":2, "letters":2, "pastoral":3, "right":3, "slimy":2, "tank":2, "writing":2,
                },
        },
    }

    for _, td := range testData {
        sm := New()

        for _,v := range td.words {
            sm.WordCounter(v)
        }

        r := sm.GetWordStats(td.wordCount)

        if !reflect.DeepEqual(r, td.expected) {
        	t.Fatalf("failed to confirm that the maps are equal (expected: %v, result: %v)", td.expected, r)
        }
    }
}


func TestSortedMap_GetWordStatsFromFile(t *testing.T) {
    expected := map[string]int {
        "like":35, "told":29, "looked":39, "marry":21, "went":62, "love":32, "want":29, "into":29, "took":20, "cant":18,
    }

    f, _ := os.Open("testdata/text.txt")
    defer f.Close()

    sm := New()
    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
        text := scanner.Text()

        sm.SplitTextIntoWords(text, 3)
    }
    wordStats := sm.GetWordStats(10)

    if !reflect.DeepEqual(wordStats, expected) {
        t.Fatalf("failed to confirm that the maps are equal (expected: %v, result: %v)", wordStats, expected)
    }
}

func BenchmarkSortedMap_GetWordStats_from_file(b *testing.B) { benchmarkSortedMapGetWordStatsFromFIle(b, 10) }
func BenchmarkSortedMap_GetWordStats_with_1000_words(b *testing.B) { benchmarkSortedMapGetWordStats(b, 1000) }
func BenchmarkSortedMap_GetWordStats_with_10000_words(b *testing.B) { benchmarkSortedMapGetWordStats(b, 10000) }
func BenchmarkSortedMap_GetWordStats_with_100000_words(b *testing.B) { benchmarkSortedMapGetWordStats(b, 100000) }


var Result map[string]int
func benchmarkSortedMapGetWordStatsFromFIle(b *testing.B, n int) {

    f, _ := os.Open("testdata/text.txt")
    defer f.Close()

    sm := New()
    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
        text := scanner.Text()

        sm.SplitTextIntoWords(text, 3)
    }
    b.ResetTimer()

    var r map[string]int
    for i := 0; i < b.N; i++ {
        r = sm.GetWordStats(10)
    }

    Result = r
}

func benchmarkSortedMapGetWordStats(b *testing.B, n int) {
    words := make([]string, n)
    wordCounter := make(map[string]int)

    for i := 0; i < n; i++ {
       word := generate(rand.Intn(10))
       words = append(words, word)
       wordCounter[word] = rand.Intn(200)
    }
    b.ResetTimer()

    sm := &SortedMap{Words: words, wordCounter: wordCounter}

    var r map[string]int
    for i := 0; i < b.N; i++ {
       r = sm.GetWordStats(10)
    }

    Result = r
}

func generate(size int) string {
    var chars = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

    b := make([]byte, size)
    rand.Read(b)

    for i := 0; i < size; i++ {
        b[i] = chars[b[i] / 5]
    }
    return *(*string)(unsafe.Pointer(&b))
}