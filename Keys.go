package search

import (
	"slices"
	"strings"

	slice "github.com/Chara-X/slices"
	"github.com/kljensen/snowball/english"
)

func Keys(s string) []string {
	return slice.Distinct(slice.Select(slices.DeleteFunc(separator.Split(strings.ToLower(s), -1), func(key string) bool { return english.IsStopWord(key) || len(key) < 2 }), func(key string) string { return english.Stem(key, false) }))
}
