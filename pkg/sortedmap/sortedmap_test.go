package sortedmap

import (
	"reflect"
	"testing"
)

func TestSortedMap_GetWordStats(t *testing.T) {
	sm := New()

	words := getTestWords()

	for _,v := range words {
		sm.WordCounter(v)
	}

	expected := map[string]int{
		"ambitious":2, "finger":2, "food":4, "healthy":2, "letters":2, "pastoral":3, "right":3, "slimy":2, "tank":2, "writing":2,
	}

	r := sm.GetWordStats(10)

	if !reflect.DeepEqual(r, expected) {
		t.Fatalf("failed to confirm that the maps are equal (expected: %v, result: %v)", expected, r)
	}
}


func getTestWords() []string {
	return []string{
		"impolite", "ready", "buzz", "healthy", "holistic","ants","decision","mix",
		"witty", "point", "weigh", "writing", "huge", "slow", "library", "damage",
		"highfalutin", "slimy", "food", "ambitious", "pastoral", "mug", "books", "rabbits", "warm", "tank",
		"embarrassed", "resolute", "stamp", "slimy", "right", "hurry", "pastoral", "motionless",
		"fact", "shelter", "food", "letters", "who", "food", "healthy", "whole", "shaggy", "pastoral", "writing", "help", "finger",
		"wicked", "tank", "story", "tendency", "right", "somber", "scream", "geese", "clap", "playground", "fool", "valuable",
		"crook", "right", "grade", "ambitious", "letters", "lie", "food", "finger", "easy", "terrific", "bit", "disastrous", "defiant", "nosy", "madly",
	}
}
