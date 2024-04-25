package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/Chara-X/priority"
	"github.com/Chara-X/search"
)

func main() {
	if len(os.Args) > 2 {
		var stopwatch = time.Now()
		var paths = []string{}
		filepath.Walk(os.Args[1], func(path string, info os.FileInfo, err error) error {
			var file, _ = os.ReadFile(path)
			if strings.Contains(string(file), os.Args[2]) {
				paths = append(paths, path)
			}
			return err
		})
		for _, v := range paths {
			fmt.Println(v)
		}
		fmt.Println(time.Since(stopwatch))
		return
	}
	var index = search.New[string, entry]()
	filepath.Walk(os.Args[1], func(path string, info os.FileInfo, err error) error {
		var file, _ = os.ReadFile(path)
		index.Store(search.Keys(string(file)), entry{search.Keys(string(file)), path})
		return err
	})
	for scanner := bufio.NewScanner(os.Stdin); scanner.Scan(); {
		var stopwatch = time.Now()
		var keys = search.Keys(scanner.Text())
		var pq = priority.New[float64, entry]()
		for _, v := range index.Load(keys...) {
			pq.Push(-search.Compare(keys, v.keys), v)
		}
		for i := 0; i < 5 && pq.Len() > 0; i++ {
			fmt.Println(pq.Pop().Value.value)
		}
		fmt.Println(time.Since(stopwatch))
	}
}

type entry struct {
	keys  []string
	value string
}
