package main

import "fmt"
import "time"

var xsize, ysize = 101, 103
type robot struct {x, y, vx, vy int}

func move(robots []robot) {
	for i, _ := range robots {
		var r = &robots[i]
		r.x += r.vx
		if r.x < 0 { r.x += xsize } else if r.x > xsize - 1 { r.x -= xsize }
		r.y += r.vy
		if r.y < 0 { r.y += ysize } else if r.y > ysize - 1 { r.y -= ysize }
	}
}

func quadrants(robots []robot) (q1, q2, q3, q4 int) {
	for _, r := range robots {
		if r.x < xsize / 2 {
			if r.y < ysize / 2 { q1++ } else if r.y > ysize / 2 { q3++ }
		} else if r.x > xsize / 2 {
			if r.y < ysize / 2 { q2++ } else if r.y > ysize / 2 { q4++ }
		}
	}
	return
}

func main() {
	var robots, easter []robot
	for {
		var r robot
		if n, _ := fmt.Scanf("p=%d,%d v=%d,%d", &r.x, &r.y, &r.vx, &r.vy); n < 4 { break }
		robots = append(robots, r)
		easter = append(easter, r)
	}
	if len(robots) < 50 {
		xsize, ysize = 11, 7
	}
	for s := 1; s <= 100; s++ {
		move(robots)
	}
	q1, q2, q3, q4 := quadrants(robots)
	fmt.Println(q1 * q2 * q3 * q4)
	var s = 1
	for {
		move(easter)
		q1, q2, q3, q4 := quadrants(easter)
		if q1 > len(easter) / 2 || q2 > len(easter) / 2 || q3 > len(easter) / 2 || q4 > len(easter) / 2 {
			for y := 0; y < ysize; y++ {
				for x := 0; x < xsize; x++ {
					var found = false
					for _, e := range easter {
						if e.x == x && e.y == y { found = true; break }
					}
					if found { fmt.Print("#") } else { fmt.Print(".") }
				}
				fmt.Println()
			}
			fmt.Println(s, "second")
			time.Sleep(time.Second)
		}
		s++
	}
}
