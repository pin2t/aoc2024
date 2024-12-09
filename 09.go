package main

import "bufio"
import "os"
import "fmt"

func main() {
	type block struct { pos, len int }
	const EMPTY = -1

	var scanner = bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var line = scanner.Text()
	var disk = make([]int, 0)
	var id = 0
	var space = make([]block, 0)
	var files = make(map[int]block)
	for i, b := range line {
		var _id, _len = EMPTY, int(b - '0')
		var pos = len(disk)
		if i % 2 == 0 {
			files[id] = block{pos, _len}
			_id = id
			id++
		} else if b != '0' {
			space = append(space, block{pos, _len})
		}
		for j := 0; j < _len; j++ {
			disk = append(disk, _id)
		}
	}
	var last = len(disk) - 1
	for ;disk[last] == EMPTY; last-- {}
	for i := 0; i < last; i++ {
		if disk[i] != EMPTY { continue }
		disk[i] = disk[last]
		disk[last] = -1
		for ;disk[last] == -1; last-- {}
	}
	var sum int64
	for i, v := range disk {
		if v != EMPTY { sum += int64(i) * int64(v) }
	}
	fmt.Print(sum)
	for id > 0 {
		id--
		var file = files[id]
		for is, sb := range space {
			if sb.pos >= file.pos {
				space = space[:is]
				break
			}
			if file.len <= sb.len {
				files[id] = block{sb.pos, file.len}
				if file.len == sb.len {
					space = append(space[:is], space[is+1:]...)
				} else {
					space[is] = block{sb.pos + file.len, sb.len - file.len}
				}
				break
			}
		}
	}
	sum = 0
	for id, file := range files {
		for pos := file.pos; pos < file.pos + file.len; pos++ {
			sum += int64(id) * int64(pos)
		}
	}
	fmt.Println("", sum)
}
