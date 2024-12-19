package main

import "bufio"
import "os"
import "strings"
import "fmt"

var npossibles = make(map[string]int)
var towels []string

func possible(design string) int {
	if len(design) == 0 { return 1 }
	if n, found := npossibles[design]; found { return n }
	var result = 0
	for _, t := range towels {
		if len(t) <= len(design) && design[0:len(t)] == t {
			var p = possible(design[len(t):])
			result += p
			npossibles[design[len(t):]] = p
		}
	}
	return result
}

func main() {
	var scanner = bufio.NewScanner(os.Stdin)
	scanner.Scan()
	towels = strings.Split(scanner.Text(), ", ")
	var possibles, total = 0, 0
	scanner.Scan()
	for scanner.Scan() {
		var p = possible(scanner.Text())
		if p > 0 { possibles++ }
		total += p
	}
	fmt.Println(possibles, total)
}
