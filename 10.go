package main

import "bufio"
import "os"
import "fmt"

func main() {
	type pos struct { row, col int }
	
	var scanner = bufio.NewScanner(os.Stdin);
	var _map [][]int
	for scanner.Scan() {
		var row []int
		for _, char := range scanner.Text() {
			row = append(row, int(char) - int('0'))
		}
		_map = append(_map, row)
	}
	var scores, ratings int
	for r, row := range _map {
		for c, _ := range row {
			if row[c] != 0 { continue }
			var visited = make(map[pos]bool)
			var queue []pos
			queue = append(queue, pos{r, c})
			for len(queue) > 0 {
				var p = queue[0]
				queue = queue[1:]
				if _map[p.row][p.col] == 9 {
					if !visited[p] { scores++ }
					ratings++
					visited[p] = true
				}
				for _, move := range []pos{{1, 0}, {0, 1}, {-1, 0}, {0, -1}} {
					if p.row + move.row >= 0 && p.row + move.row < len(_map) &&
					   p.col + move.col >= 0 && p.col + move.col < len(_map[p.row])	&&
					   (_map[p.row + move.row][p.col + move.col] - _map[p.row][p.col]) == 1 {
						queue = append(queue, pos{p.row + move.row, p.col + move.col})
					}
				}
			}
		}
	}
	fmt.Println(scores, ratings)
}
