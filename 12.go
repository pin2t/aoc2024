package main

import "bufio"
import "os"
import"fmt"
import "slices"

type pos struct {row, col int}

var processed = make(map[pos]bool)
var _map [][]rune

func inside(p pos) bool {
	return p.row >= 0 && p.row < len(_map) && p.col >= 0 && p.col < len(_map[p.row]);
}

func expand(t rune, p pos, outsides *int, reg map[pos]bool) {
	if !inside(p) || _map[p.row][p.col] != t {
		(*outsides)++
		return
	}
	if processed[p] {
		return
	}
	processed[p] = true
	reg[p] = true
	for _, d := range []pos{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
		expand(t, pos{p.row + d.row, p.col + d.col}, outsides, reg);
	}
}

func region(p pos) (int, int) {
	var reg = make(map[pos]bool)
	var outsides, sides int
	var _type = _map[p.row][p.col]
	expand(_type, p, &outsides, reg)
	var outDirections = make(map[pos][]pos)
	for row, _ := range _map {
		for col, _ := range _map[0] {
			if !reg[pos{row, col}] { continue }
			for _, d := range []pos{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
				var np = pos{row + d.row, col + d.col}
				if inside(np) && _map[np.row][np.col] == _type { continue }
				var ldir, rdir = pos{-d.col, d.row}, pos{d.col, -d.row}
				var left = pos{row + ldir.row, col + ldir.col}
				var right = pos{row + rdir.row, col + rdir.col}
				if !slices.Contains(outDirections[left], d) && !slices.Contains(outDirections[right], d) {
					sides++
				}
				outDirections[pos{row, col}] = append(outDirections[pos{row, col}], d)
			}
		}
	}
	return outsides, sides
}

func main() {
	var scanner = bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		_map = append(_map, []rune(scanner.Text()))
	}
	var sums [2]int64
	var p pos
	for p.row, _ = range _map {
		for p.col, _ = range _map[p.row] {
			if !processed[p] {
				var prev = len(processed)
				var outsides, sides = region(p)
				var area = len(processed) - prev
				sums[0] += int64(area) * int64(outsides)
				sums[1] += int64(area) * int64(sides)
			}
		}
	}
	fmt.Println(sums)
}
