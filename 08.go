package main

import "bufio"
import "os"
import "fmt"

type pos struct { row, col int }

func abs(a, b int) int {
	if a < b { return b - a } else { return a - b }
}

var rows, cols int
var antennas = make(map[rune][]pos)
var antinodes = make(map[pos]bool)

func putInbound(p pos) {
	if p.row >= 0 && p.row < rows && p.col >= 0 && p.col < cols {
		antinodes[p] = true
	}
}

func putInline(a1, a2 pos, distance int) {
	var dr, dc = abs(a1.row, a2.row), abs(a1.col, a2.col)
	if a1.row > a2.row {
		if a1.col > a2.col {
			putInbound(pos{a2.row - distance * dr, a2.col - distance * dc})
			putInbound(pos{a1.row + distance * dr, a1.col + distance * dc})
		}  else {
			putInbound(pos{a2.row - distance * dr, a2.col + distance * dc})
			putInbound(pos{a1.row + distance * dr, a1.col - distance * dc})
		}
	} else {
		if a1.col > a2.col {
			putInbound(pos{a2.row + distance * dr, a2.col - distance * dc})
			putInbound(pos{a1.row - distance * dr, a1.col + distance * dc})
		}  else {
			putInbound(pos{a2.row + distance * dr, a2.col + distance * dc})
			putInbound(pos{a1.row - distance * dr, a1.col - distance * dc})
		}
	}
}

func putAll(distance int) int {
	var prev = len(antinodes)
	for _, as := range antennas {
		for i := 0; i < len(as) - 1; i++ {
			for j := i + 1; j < len(as); j++ {
				putInline(as[i], as[j], distance)
			}
		}
	}
	return len(antinodes) - prev
}

func main() {
	var scanner = bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var row = scanner.Text()
		cols = len(row)
		for col, c := range row {
			if c == '.' { continue }
			antennas[c] = append(antennas[c], pos{rows, col})
		}
		rows++
	}
	putAll(1)
	fmt.Print(len(antinodes))
	antinodes = make(map[pos]bool)
	for n := 0; putAll(n) > 0; n++ {}
	fmt.Println("", len(antinodes))
}
