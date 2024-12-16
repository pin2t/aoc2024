package main

import "bufio"
import "os"
import "fmt"

type pos struct {row, col int}
type dir struct {dr, dc int}

var boxes, walls, boxes2, walls2 = make(map[pos]bool), make(map[pos]bool), make(map[pos]bool), make(map[pos]bool)
var moves = map[rune]dir{'^': {-1, 0}, '>': {0, 1}, 'v': {1, 0}, '<': {0, -1}}

func canMove(box pos, d rune) bool {
	if !boxes2[box] { panic("invalid argument") }
	switch d {
	case '^':
		if walls2[pos{box.row - 1, box.col}] || walls2[pos{box.row - 1, box.col + 1}] { return false }
		var next = pos{box.row - 1, box.col - 1}
		if boxes2[next] && !canMove(next, d) { return false }
		next = pos{box.row - 1, box.col}
		if boxes2[next] && !canMove(next, d) { return false }
		next = pos{box.row - 1, box.col + 1}
		if boxes2[next] && !canMove(next, d) { return false }
	case '>':
		if walls2[pos{box.row, box.col + 2}] { return false }
		var next = pos{box.row, box.col + 2}
		if boxes2[next] && !canMove(next, d) { return false }
	case 'v':
		if walls2[pos{box.row + 1, box.col}] || walls2[pos{box.row + 1, box.col + 1}] { return false }
		var next = pos{box.row + 1, box.col - 1}
		if boxes2[next] && !canMove(next, d) { return false }
		next = pos{box.row + 1, box.col}
		if boxes2[next] && !canMove(next, d) { return false }
		next = pos{box.row + 1, box.col + 1}
		if boxes2[next] && !canMove(next, d) { return false }
	case '<':
		if walls2[pos{box.row, box.col - 1}] { return false }
		var next = pos{box.row, box.col - 2}
		if boxes2[next] && !canMove(next, d) { return false }
	}
	return true
}

func move(box pos, d rune) {
	if !boxes2[box] { panic("invalid argument") }
	switch d {
	case '^':
		var next = pos{box.row - 1, box.col - 1}
		if boxes2[next] { move(next, d) }
		next = pos{box.row - 1, box.col}
		if boxes2[next] { move(next, d) }
		next = pos{box.row - 1, box.col + 1}
		if boxes2[next] { move(next, d) }
		boxes2[pos{box.row - 1, box.col}] = true
	case '>':
		var next = pos{box.row, box.col + 2}
		if boxes2[next] { move(next, d) }
		boxes2[pos{box.row, box.col + 1}] = true
	case 'v':
		var next = pos{box.row + 1, box.col - 1}
		if boxes2[next] { move(next, d) }
		next = pos{box.row + 1, box.col}
		if boxes2[next] { move(next, d) }
		next = pos{box.row + 1, box.col + 1}
		if boxes2[next] { move(next, d) }
		boxes2[pos{box.row + 1, box.col}] = true
	case '<':
		var next = pos{box.row, box.col - 2}
		if boxes2[next] { move(next, d) }
		boxes2[pos{box.row, box.col - 1}] = true
	}
	delete(boxes2, box)
}

func main() {
	var row, cols int
	var scanner = bufio.NewScanner(os.Stdin)
	var robot, robot2 pos
	for scanner.Scan() {
		if len(scanner.Text()) == 0 { break }
		for col, c := range scanner.Text() {
			switch c {
			case '#':
				walls[pos{row, col}] = true
				walls2[pos{row, 2 * col}] = true
				walls2[pos{row, 2 * col + 1}] = true
			case 'O':
				boxes[pos{row, col}] = true
				boxes2[pos{row, 2 * col}] = true
			case '@':
				robot = pos{row, col}
				robot2 = pos{row, 2 * col}
			}
			if col > cols { cols = col }
		}
		row++
	}
	for scanner.Scan() {
		for _, c := range scanner.Text() {
			var m = moves[c]
			var p = pos{robot.row + m.dr, robot.col + m.dc}
			for {
				if !boxes[p] { break }
				p = pos{p.row + m.dr, p.col + m.dc}
			}
			if !walls[p] {
				var end = p
				robot = pos{robot.row + m.dr, robot.col + m.dc}
				p = robot
				var from = p
				for p != end {
					var next = pos{p.row + m.dr, p.col + m.dc}
					boxes[next] = true
					p = next
				}
				delete(boxes, from)
			}
			p = pos{robot2.row + m.dr, robot2.col + m.dc}
			switch c {
			case '^':
				if boxes2[p] && canMove(p, c) {
					move(p, c)
				} else if boxes2[pos{p.row, p.col - 1}] && canMove(pos{p.row, p.col - 1}, c) {
					move(pos{p.row, p.col - 1}, c)
				}
			case '>':
				if boxes2[p] && canMove(p, c) { move(p, c) }
			case 'v':
				if boxes2[p] && canMove(p, c) {
					move(p, c)
				} else if boxes2[pos{p.row, p.col - 1}] && canMove(pos{p.row, p.col - 1}, c) {
					move(pos{p.row, p.col - 1}, c)
				}
			case '<':
				if boxes2[pos{p.row, p.col - 1}] && canMove(pos{p.row, p.col - 1}, c) {
					move(pos{p.row, p.col - 1}, c)
				}
			}
			if !walls2[p] && !boxes2[p] && !boxes2[pos{p.row, p.col - 1}] {
				robot2 = p
			}
		}
	}
	var coords [2]int64
	for b, _ := range boxes { coords[0] += int64(100*b.row + b.col) }
	for b, _ := range boxes2 { coords[1] += int64(100*b.row + b.col) }
	fmt.Println(coords)
}
