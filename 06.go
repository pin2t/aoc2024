package main

import "bufio"
import "os"
import "strings"
import "fmt"

type pos struct {x, y int}

var _map [][]rune
var directions = []pos{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}

func loop(guard pos, di int) (bool, int) {
	var visited = make(map[pos]bool, 10000)
	var onPath int                                            // number of steps on already existing path
	visited[guard] = true
	for {
		var next = pos{guard.x + directions[di % 4].x, guard.y + directions[di % 4].y}
		if next.x < 0 || next.y < 0 || next.x >= len(_map[guard.y]) || next.y >= len(_map) {
			return false, len(visited)
		}
		if onPath > 100 { return true, len(visited) }         // assume a loop when stepping on path for some time
		if _map[next.y][next.x] == '#' {
			di++
		} else {
			if visited[next] { onPath++ } else { onPath = 0 }
			guard = next
		}
		visited[guard] = true
	}
}

func main() {
	var scanner = bufio.NewScanner(os.Stdin)
	var guard pos
	var y int
	for scanner.Scan() {
		_map = append(_map, []rune(scanner.Text()))
		var p = strings.IndexRune(scanner.Text(), '^')
		if p > 0 { guard = pos{p, y} }
		y++
	}
	var _, length = loop(guard, 0)
	fmt.Print(length)
	var nloops int
	for y, row := range _map {
		for x, c := range row {
			if c != '^' && c != '#' {
				_map[y][x] = '#'
				var isLoop, _ = loop(guard, 0)
				if isLoop { nloops++ }
				_map[y][x] = '.'
			}
		}
	}
	fmt.Println("", nloops)
}