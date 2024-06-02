package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"time"
	"unicode/utf8"

	"github.com/Chara-X/priority"
	"github.com/Chara-X/search"
)

func main() {
	var index = search.New[string, entry]()
	filepath.Walk(os.Args[1], func(path string, info os.FileInfo, err error) error {
		var file, _ = os.ReadFile(path)
		if utf8.Valid(file) {
			var keys = search.Keys(string(file))
			index.Store(keys, entry{keys, path})
		}
		return err
	})
	fmt.Println()
	for scanner := bufio.NewScanner(os.Stdin); scanner.Scan(); {
		var stopwatch = time.Now()
		var keys = search.Keys(scanner.Text())
		var pq = &priority.Queue[float64, string]{}
		for _, v := range index.Load(keys...) {
			pq.Push(-search.Compare(keys, v.keys), v.value)
		}
		for i := 0; i < 5 && pq.Len() > 0; i++ {
			fmt.Println(pq.Pop().Value)
		}
		fmt.Println(time.Since(stopwatch))
	}
}

type entry struct {
	keys  []string
	value string
}
