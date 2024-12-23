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
	fmt.Print(tconns)
	var foreach = func (comp string, process func (second string)) {
		for _, second := range comps {
			if second != comp && conns[conn{comp, second}] {
				process(second)
			}
		}
	}
	var largest = make(map[string]bool)
	for _, comp := range comps {
		foreach(comp, func (second string) {
			var lan = make(map[string]bool)
			lan[comp] = true
			lan[second] = true
			foreach(comp, func (second string) {
				if lan[second] { return }
				var all = true
				foreach(second, func (_ string) {
					 for lc, _ := range lan {
						 if !conns[conn{second, lc}] { all = false }
					 }
				})
				if all { lan[second] = true }
			})
			if len(lan) > len(largest) { largest = lan }
		})
	}
	var list []string
	for c, _ := range largest { list = append(list, c) }
	slices.Sort(list)
	fmt.Println("", strings.Join(list, ","))
}
