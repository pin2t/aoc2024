package main

import "bufio"
import "os"
import "fmt"

func main() {
	var world [][]rune
	var scanner = bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		world = append(world, []rune(scanner.Text()))
	}
	var char = func(x, y int) rune {
		if y < 0 || y >= len(world) || x < 0 || x >= len(world[y]) { return ' ' }
		return world[y][x]
	}
	var xmas = func(x, y, dx, dy int) bool {
		return char(x, y) == 'X' && char(x + dx, y + dy) == 'M' && char(x + 2 * dx, y + 2 * dy) == 'A' && char(x + 3 * dx, y + 3 * dy) == 'S'
	}
	var crossMAS = func(x, y int) bool {
		return char(x, y) == 'A' &&
			(char(x - 1, y - 1) == 'M' && char(x + 1, y - 1) == 'M' && char(x - 1, y + 1) == 'S' && char(x + 1, y + 1) == 'S' || // M's up
			 char(x - 1, y + 1) == 'M' && char(x + 1, y + 1) == 'M' && char(x - 1, y - 1) == 'S' && char(x + 1, y - 1) == 'S' || // M's down
			 char(x - 1, y - 1) == 'M' && char(x - 1, y + 1) == 'M' && char(x + 1, y - 1) == 'S' && char(x + 1, y + 1) == 'S' || // M's left
			 char(x + 1, y - 1) == 'M' && char(x + 1, y + 1) == 'M' && char(x - 1, y - 1) == 'S' && char(x - 1, y + 1) == 'S')   // M's right
	}
	var xmases, crosses int
	for y, row := range world {
		for x, _ := range row {
			for _, direction := range [][]int{{1, 0}, {1, 1}, {0, 1}, {-1, 1}, {-1, 0}, {-1, -1}, {0, -1}, {1, -1}} {
				if xmas(x, y, direction[0], direction[1]) { xmases++ }
			}
			if crossMAS(x, y) { crosses++ }
		}
	}
	fmt.Println(xmases, crosses)
}
