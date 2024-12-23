package main

import "bufio"
import "os"
import "strings"
import "fmt"
import "slices"

func main() {
	var scanner = bufio.NewScanner(os.Stdin)
	type conn struct {c1, c2 string}
	var conns = make(map[conn]bool)
	var compSet = make(map[string]bool)
	for scanner.Scan() {
		var c = strings.Split(scanner.Text(), "-")
		conns[conn{c[0], c[1]}] = true
		conns[conn{c[1], c[0]}] = true
		compSet[c[0]] = true
		compSet[c[1]] = true
	}
	var comps []string
	for c, _ := range compSet { comps = append(comps, c) }
	var tconns = 0
	for i := 0; i < len(comps) - 2; i++ {
		for j := i + 1; j < len(comps) - 1; j++ {
			for k := j + 1; k < len(comps); k++ {
				var c1, c2, c3 = comps[i], comps[j], comps[k]
				if conns[conn{c1, c2}] && conns[conn{c2, c3}] && conns[conn{c1, c3}] &&
					(c1[0] == 't' || c2[0] == 't' || c3[0] == 't') {
					tconns++
				}
			}
		}
	}
	var largest = make(map[string]bool)
	for _, comp := range comps {
		for _, second := range comps {
			var lan = make(map[string]bool)
			if second != comp && conns[conn{comp, second}] {
				lan[comp] = true
				lan[second] = true
			}
			if len(lan) < 2 { continue }
			for _, next := range comps {
				if next != comp && conns[conn{comp, next}] && !lan[next] {
					var all = true
					for _, nc := range comps {
						if nc != next && conns[conn{next, nc}] {
							 for lc, _ := range lan {
								 if !conns[conn{next, lc}] { all = false }
							 }
						}
					}
					if all { lan[next] = true }
				}
			}
			if len(lan) > len(largest) { largest = lan }
		}
	}
	var list []string
	for c, _ := range largest { list = append(list, c) }
	slices.Sort(list)
	fmt.Println(tconns, strings.Join(list, ","))
}
