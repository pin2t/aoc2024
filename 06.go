package main

import "bufio"
import "os"
import "strings"
import "fmt"

type pos struct {x, y int}

var _map [][]rune
var directions = []pos{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}              // up, right, down, left (in order of 90 degree turns)

// return true if guard stuck in a loop. Second item in returned pair is number of visited points
func loop(guard pos, di int) (bool, int) {
	var visited = make(map[pos]bool, 10000)
	type vdt struct {p pos; di int}
	var visitedDir = make(map[vdt]bool, 10000)
	visited[guard] = true
	visitedDir[vdt{guard, di}] = true
	for {
		var next = pos{guard.x + directions[di].x, guard.y + directions[di].y}
		if next.x < 0 || next.y < 0 || next.x >= len(_map[guard.y]) || next.y >= len(_map) {
			return false, len(visited)
		}
		if visitedDir[vdt{next, di}] { return true, len(visited) }
		if _map[next.y][next.x] == '#' {
			di = (di + 1) % 4
		} else {
			guard = next
		}
		visited[guard] = true
		visitedDir[vdt{guard, di}] = true
	}
}

func main() {
	var scanner = bufio.NewScanner(os.Stdin)
	var guard pos
	var y int
	for scanner.Scan() {
		_map = append(_map, []rune(scanner.Text()))
		if p := strings.IndexRune(scanner.Text(), '^'); p > 0 { guard = pos{p, y} }
		y++
	}
	var _, nvisited = loop(guard, 0)
	fmt.Print(nvisited)
	var nloops int
	for y, row := range _map {
		for x, c := range row {
			if c != '^' && c != '#' {
				_map[y][x] = '#'
				if isLoop, _ := loop(guard, 0); isLoop { nloops++ }
				_map[y][x] = '.'
			}
		}
	}
	fmt.Println("", nloops)
}
