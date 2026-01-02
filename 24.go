package main
import "bufio"
import "os"
import "strings"
import "fmt"
import "slices"

type gate struct {op string; in1, in2 string}
var inputs = make(map[string]bool)
var gates = make(map[string]gate)

func main() {
	var scanner = bufio.NewScanner(os.Stdin)
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
				switch (val.op) {
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
	var swaps = make([]string, 0)
	for len(swaps) < 8 {
		var base = progress()
		out:
		for aname, a := range gates {
			for bname, b := range gates {
				if a == b { continue }
				gates[aname] = b
				gates[bname] = a
				if progress() > base {
					swaps = append(append(swaps, bname), aname)
					break out
				}
				gates[aname] = a
				gates[bname] = b
			}
		}
	}
	slices.Sort(swaps)
	fmt.Println(val, strings.Join(swaps, ","))
}

func check(nm string, bit int, checkgate func (g gate, bit int) bool) bool {
	g, ok := gates[nm]
	if !ok { return false }
	return checkgate(g, bit)
}

func add(nm string, bit int) bool {
	return check(nm, bit, func (g gate, bit int) bool {
		if bit == 0 { return g.op == "XOR" && (g.in1 == "x00" && g.in2 == "y00" || g.in1 == "y00" && g.in2 == "x00") }
		return g.op == "XOR" && (inter(g.in1, bit) && carry(g.in2, bit) || inter(g.in2, bit) && carry(g.in1, bit))
	})
}

func inter(nm string, bit int) bool {
	return check(nm, bit, func (g gate, bit int) bool {
		return g.op == "XOR" && (g.in1 == name("x", bit) && g.in2 == name("y", bit) || g.in1 == name("y", bit) && g.in2 == name("x", bit))
	})
}

func carry(nm string, bit int) bool {
	return check(nm, bit, func (g gate, bit int) bool {
		if bit == 1 {
			return g.op == "AND" && (g.in1 == "x00" && g.in2 == "y00" || g.in1 == "y00" && g.in2 == "x00")
		}
		return g.op == "OR" && (direct(g.in1, bit - 1) && recarry(g.in2, bit - 1) || direct(g.in2, bit - 1) && recarry(g.in1, bit - 1))
	})
}

func direct(nm string, bit int) bool {
	return check(nm, bit, func (g gate, bit int) bool {
		return g.op == "AND" && (g.in1 == name("x", bit) && g.in2 == name("y", bit) || g.in1 == name("y", bit) && g.in2 == name("x", bit))
	})
}

func recarry(nm string, bit int) bool {
	return check(nm, bit, func (g gate, bit int) bool {
		return g.op == "AND" && (inter(g.in1, bit) && carry(g.in2, bit) || inter(g.in2, bit) && carry(g.in1, bit))
	})
}

func progress() (bit int) {
	for ; add(name("z", bit), bit); bit++ {}
	return
}

func name(pre string, bit int) string {
	return pre + fmt.Sprintf("%02d", bit)
}