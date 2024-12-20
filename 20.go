package main

import "bufio"
import "os"
import "slices"
import "fmt"

type pos struct {row, col int}
type dir struct {row, col int}

func main() {
	var scanner = bufio.NewScanner(os.Stdin)
	var track [][]rune
	for scanner.Scan() {
		track = append(track, []rune(scanner.Text()))
	}
	var start, end pos
	for r, row := range track {
		for c, char := range row  {
			if char == 'S' { start = pos{r, c} }
			if char == 'E' { end = pos{r, c} }
		}
	}
	var path = path(track, start, end)
	var n2, n20 = 0, 0
	for i := 0; i < len(path) - 100; i++ {
		for j := i + 100; j < len(path); j++ {
			var cheatLen = abs(path[i].row - path[j].row) + abs(path[i].col - path[j].col)
			if cheatLen <= 20 && ((j - i) - cheatLen >= 100) { n20++ }
			if cheatLen <= 2 && ((j - i) - cheatLen >= 100) { n2++ }
		}
	}
	fmt.Println(n2, n20)
}

func abs(val int) int {
	if val < 0 { return -val }
	return val
}

func path(track [][]rune, start pos, end pos) []pos {
	type state struct {p pos; picosecnds int; path []pos}
	var queue []state
	var enqueue = func (p pos, picoseconds int, prevpath []pos) {
		var path = make([]pos, len(prevpath) + 1)
		copy(path, prevpath)
		path[len(path) - 1] = p
		for i := 0; i < len(queue); i++ {
			if queue[i].picosecnds >= picoseconds {
				queue = slices.Insert(queue, i, state{p, picoseconds, path})
				return
			}
		}
		queue = append(queue, state{p, picoseconds, path})
	}
	var dequeue = func () (p pos, steps int, path []pos) {
		p, steps, path = queue[0].p, queue[0].picosecnds, queue[0].path
		queue = queue[1:]
		return
	}
	enqueue(start, 0, []pos{})
	var processed = make(map[pos]bool)
	for len(queue) > 0 {
		var p, picoseconds, path = dequeue()
		if p == end { return path }
		if processed[p] { continue }
		processed[p] = true
		for _, d := range []dir{{1, 0}, {0, 1}, {-1, 0}, {0, -1}} {
			var next = pos{p.row + d.row, p.col + d.col}
			if next.row >= 0 && next.row < len(track) &&
			   next.col >= 0 && next.col < len(track[next.row]) &&
    		   track[next.row][next.col] != '#' {
                enqueue(next, picoseconds + 1, path)
            }
		}
	}
	return []pos{}
}
