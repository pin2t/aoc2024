package main

import "bufio"
import "os"
import "fmt"
import "slices"

type pos struct { row, col int }
type dir struct { dr, dc int }
type posdir struct { p pos; d dir }

var maze [][]rune
var start, end pos

func part1() (score int) {
	type state struct { p pos; d dir; score int }
	var queue []state
	var processed = make(map[posdir]bool)
	var enqueue = func (p pos, d dir, score int) {
		for i := 0; i < len(queue); i++ {
			if queue[i].score >= score {
				queue = slices.Insert(queue, i, state{p, d, score})
				return
			}
		}
		queue = append(queue, state{p, d, score})
	}
	enqueue(start, dir{0, 1}, 0)
	for len(queue) > 0 {
		var st = queue[0]
		queue = queue[1:]
		var pd = posdir{st.p, st.d}
		if processed[pd] { continue }
		processed[pd] = true
		if st.p == end {
			score = st.score
			return
		}
		var next = pos{st.p.row + st.d.dr, st.p.col + st.d.dc}
		if maze[next.row][next.col] != '#' {
			enqueue(next, st.d, st.score + 1)
		}
		enqueue(st.p, dir{st.d.dc, -st.d.dr}, st.score + 1000)
		enqueue(st.p, dir{-st.d.dc, st.d.dr}, st.score + 1000)
	}
	return
}

func part2() int {
	type state struct { p pos; d dir; score int; path map[pos]bool }
	var queue []state
	var processed = make(map[posdir]int)
	var enqueue = func (p pos, d dir, score int, path map[pos]bool) {
		var extended = make(map[pos]bool)
		for p, _ := range path { extended[p] = true }
		extended[p] = true
		for i := 0; i < len(queue); i++ {
			if queue[i].score >= score {
				queue = slices.Insert(queue, i, state{p, d, score, extended})
				return
			}
		}
		queue = append(queue, state{p, d, score, extended})
	}
	var singleton = make(map[pos]bool)
	singleton[start] = true
	var seats = make(map[pos]bool)
	var minscore = 1000000000
	enqueue(start, dir{0, 1}, 0, singleton)
	for len(queue) > 0 {
		var st = queue[0]
		queue = queue[1:]
		if st.score > minscore { continue }
		var pd = posdir{st.p, st.d}
		if v, ok := processed[pd]; ok && st.score > v  { continue }
		processed[pd] = st.score
		if st.p == end {
			if st.score < minscore { minscore = st.score }
			for p, _ := range st.path { seats[p] = true }
		}
		var next = pos{st.p.row + st.d.dr, st.p.col + st.d.dc}
		if maze[next.row][next.col] != '#' {
			enqueue(next, st.d, st.score + 1, st.path)
		}
		enqueue(st.p, dir{st.d.dc, -st.d.dr}, st.score + 1000, st.path)
		enqueue(st.p, dir{-st.d.dc, st.d.dr}, st.score + 1000, st.path)
	}
	return len(seats)
}

func main() {
	var scanner = bufio.NewScanner(os.Stdin)
	var row = 0
	for scanner.Scan() {
		var line = scanner.Text()
		for col, cell := range line {
			if cell == 'S' { start = pos{row, col} }
			if cell == 'E' { end = pos{row, col} }
		}
		maze = append(maze, []rune(line))
		row++
	}
	fmt.Print(part1())
	fmt.Println(" ", part2())
}
