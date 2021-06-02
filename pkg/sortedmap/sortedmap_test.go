package sortedmap

import (
    "math/rand"
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

func BenchmarkSortedMap_GetWordStats_with_1000_words(b *testing.B) {

    words := make([]string, 1000)
    wordCounter := make(map[string]int)

    for i := 0; i < 1000; i++ {
       word := generate(rand.Intn(10))
       words = append(words, word)
       wordCounter[word] = rand.Intn(200)
    }

    sm := &SortedMap{Words: words, wordCounter: wordCounter}

    for i := 0; i < b.N; i++ {
       _ = sm.GetWordStats(10)
    }
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