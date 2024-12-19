package main

import "bufio"
import "os"
import "strings"
import "fmt"

var total = 0
var cache = make(map[string]int)
var towels []string

func possible(design string) bool {
	if len(design) == 0 {
		total++
		return true
	}
	if n, found := cache[design]; found {
		total += n
		return true
	}
	var result = false
	for _, t := range towels {
		var prev = total
		if len(t) <= len(design) && design[0:len(t)] == t && possible(design[len(t):]) {
			result = true
			cache[design[len(t):]] = total - prev
		}
	}
	return result
}

func main() {
	var scanner = bufio.NewScanner(os.Stdin)
	scanner.Scan()
	towels = strings.Split(scanner.Text(), ", ")
	var possibles = 0
	scanner.Scan()
	for scanner.Scan() {
		if possible(scanner.Text()) { possibles++ }
	}
	fmt.Println(possibles, total)
}
