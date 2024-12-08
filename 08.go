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

func put(p pos) {
	if p.row >= 0 && p.row < rows && p.col >= 0 && p.col < cols {
		antinodes[p] = true
	}
}

func putAll(distance int) int {
	var prev = len(antinodes)
	for _, as := range antennas {
		for i := 0; i < len(as) - 1; i++ {
			for j := i + 1; j < len(as); j++ {
				put(pos{as[i].row - distance * (as[j].row - as[i].row), as[i].col - distance * (as[j].col - as[i].col)})
				put(pos{as[j].row - distance * (as[i].row - as[j].row), as[j].col - distance * (as[i].col - as[j].col)})
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
