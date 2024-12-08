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

func putDistance(d int) int {
	var prev = len(antinodes)
	for _, as := range antennas {
		for i := 0; i < len(as) - 1; i++ {
			for j := i + 1; j < len(as); j++ {
				var dr, dc = abs(as[i].row, as[j].row), abs(as[i].col, as[j].col)
				if as[i].row > as[j].row {
					if as[i].col > as[j].col {
						putInbound(pos{as[j].row - d * dr, as[j].col - d * dc})
						putInbound(pos{as[i].row + d * dr, as[i].col + d * dc})
					}  else {
						putInbound(pos{as[j].row - d * dr, as[j].col + d * dc})
						putInbound(pos{as[i].row + d * dr, as[i].col - d * dc})
					}
				} else {
					if as[i].col > as[j].col {
						putInbound(pos{as[j].row + d * dr, as[j].col - d * dc})
						putInbound(pos{as[i].row - d * dr, as[i].col + d * dc})
					}  else {
						putInbound(pos{as[j].row + d * dr, as[j].col + d * dc})
						putInbound(pos{as[i].row - d * dr, as[i].col - d * dc})
					}
				}
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
	putDistance(1)
	fmt.Print(len(antinodes))
	antinodes = make(map[pos]bool)
	for n := 0; putDistance(n) > 0; n++ {}
	fmt.Println("", len(antinodes))
}
