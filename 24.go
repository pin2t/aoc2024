package main

import "bufio"
import "os"
import "strings"
import "fmt"

func main() {
	type gate struct {t string; in1, in2 string}
	var scanner = bufio.NewScanner(os.Stdin)
	var inputs = make(map[string]bool)
	var gates = make(map[string]gate)
	for scanner.Scan() {
		var l = scanner.Text()
		if l == "" { break }
		var s = strings.Split(l, ": ")
		inputs[s[0]] = s[1] == "1"
	}
	var zd = make(map[string]bool)
	for scanner.Scan() {
		var s = strings.Split(scanner.Text(), " ")
		gates[s[4]] = gate{s[1], s[0], s[2]}
		if s[4][0] == 'z' { zd[s[4]] = true }
	}
	var zv = make(map[string]bool)
	for len(zv) < len(zd) {
		for g, val := range gates {
			var in1, found1 = inputs[val.in1]
			var in2, found2 = inputs[val.in2]
			if found1 && found2 {
				switch (val.t) {
				case "AND": inputs[g] = in1 && in2
				case "OR": inputs[g] = in1 || in2
				case "XOR": inputs[g] = (in1 && !in2) || (!in1 && in2)
				}
				if g[0] == 'z' { zv[g] = true }
			}
		}
	}
	var val int64
	for i := 64; i >= 0; i-- {
		if inputs["z" + fmt.Sprintf("%02d", i)] {
			val = (val * 2) | 1
 		} else {
			val *= 2
 		}
	}
	fmt.Println(val)
}
