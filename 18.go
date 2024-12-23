package main

import "bufio"
import "os"
import "fmt"
import "slices"
import "strconv"

const size = 70
type pos struct {x, y int}
type dir struct {dx, dy int}

func steps(corrupted map[pos]bool) int {
	type state struct {p pos; steps int}
	var queue []state
	var enqueue = func (p pos, steps int) {
		for i := 0; i < len(queue); i++ {
			if queue[i].steps >= steps {
				queue = slices.Insert(queue, i, state{p, steps})
				return
			}
		}
		queue = append(queue, state{p, steps})
	}
	var dequeue = func () (p pos, steps int) {
		p, steps = queue[0].p, queue[0].steps
		queue = queue[1:]
		return
	}
	enqueue(pos{0, 0}, 0)
	var end = pos{size, size}
	var processed = make(map[pos]bool)
	for len(queue) > 0 {
		var p, steps = dequeue()
		if p == end { return steps }
		if processed[p] { continue }
		processed[p] = true
		for _, d := range []dir{{1, 0}, {0, 1}, {-1, 0}, {0, -1}} {
			var next = pos{p.x + d.dx, p.y + d.dy}
			if next.x >=0 && next.x <= size && next.y >= 0 && next.y  <= size && !corrupted[next] {
				enqueue(next, steps + 1)
			}
		}
	}
	return -1
}

func main() {
	var scanner = bufio.NewScanner(os.Stdin)
	var corrupted = make(map[pos]bool)
	var bytes []pos
	for scanner.Scan() {
		var x, y int
		fmt.Sscanf(scanner.Text(), "%d,%d", &x, &y)
		bytes = append(bytes, pos{x, y})
	}
	for i := 0; i < 1024; i++ { corrupted[bytes[i]] = true }
	fmt.Print(steps(corrupted))
	for i := 1024; i < len(bytes); i++ {
		corrupted[bytes[i]] = true
		if steps(corrupted) == -1 {
			fmt.Println("", strconv.Itoa(bytes[i].x) + "," + strconv.Itoa(bytes[i].y))
			break
		}
	}
}
