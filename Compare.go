package search

import "github.com/Chara-X/util/slices"

func Compare[T comparable](from, to []T) float64 {
	var set, count = slices.ToSet(to), 0.0
	for _, v := range from {
		if _, ok := set[v]; ok {
			count++
		}
	}
	return count / float64(len(to))
}
