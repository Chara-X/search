package search

import (
	"strings"

	"github.com/Chara-X/slices"
	"github.com/kljensen/snowball/english"
)

func Keys(s string) []string {
	return slices.Distinct(slices.Select(slices.Where(separator.Split(strings.ToLower(s), -1), func(key string) bool { return !english.IsStopWord(key) && len(key) > 2 }), func(key string) string { return english.Stem(key, false) }))
}
