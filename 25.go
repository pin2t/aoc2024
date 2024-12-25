package main

import "bufio"
import "os"
import "fmt"

func main() {
	var scanner = bufio.NewScanner(os.Stdin)
	type lock [5]int
	type key [5]int
	var locks []lock
	var keys []key
	var lines []string
	var parse = func () {
		if lines[0][0] == '#' {
			var l lock
			for c := 0; c < 5; c++ {
				var r = 1
				for r < len(lines) && lines[r][c] == '#' { r++ }
				l[c] = r - 1
			}
			locks = append(locks, l)
		} else {
			var k key
			for c := 0; c < 5; c++ {
				var r = len(lines) - 1
				for r >= 0 && lines[r][c] == '#' { r-- }
				k[c] = len(lines) - r - 2
			}
			keys = append(keys, k)
		}
		lines = make([]string, 0)
	}
	for scanner.Scan() {
		var l = scanner.Text()
		if len(l) > 0 {
			lines = append(lines, l)
		} else { parse() }
	}
	parse()
	type combKey struct { l lock; k key }
	var combinations = make(map[combKey]bool)
	for _, l := range locks {
		for _, k := range keys {
			var fit = true
			for c := 0; c < 5 && fit; c++ { fit = l[c] + k[c] <= 5 }
			if fit { combinations[combKey{l, k}] = true }
		}
	}
	fmt.Println(len(combinations))
}
