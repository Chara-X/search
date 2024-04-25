package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"

	"github.com/Chara-X/priority"
	"github.com/Chara-X/search"
)

func main() {
	var index = search.New[string, entry]()
	filepath.Walk(os.Args[1], func(path string, info os.FileInfo, err error) error {
		var file, _ = os.ReadFile(path)
		index.Store(search.Keys(string(file)), entry{search.Keys(string(file)), path})
		return err
	})
	for scanner := bufio.NewScanner(os.Stdin); scanner.Scan(); {
		var keys = search.Keys(scanner.Text())
		var pq = priority.New[float64, entry]()
		for _, v := range index.Load(keys...) {
			pq.Push(-search.Compare(keys, v.keys), v)
		}
		for i := 0; i < 5 && pq.Len() > 0; i++ {
			fmt.Println(pq.Pop().Value.value)
		}
	}
}

type entry struct {
	keys  []string
	value string
}
