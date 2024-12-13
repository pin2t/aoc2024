package main

import "bufio"
import "os"
import "fmt"

type point struct {x, y int64}

func play(a, b point, prize point) (win bool, spent int64) {
	spent = 400
	for atimes := int64(0); atimes <= 100; atimes++ {
		for btimes := int64(0); btimes <= 100; btimes++ {
			if a.x * atimes + b.x * btimes == prize.x && a.y * atimes + b.y * btimes == prize.y {
				win = true
				if atimes * 3 + btimes < spent {
					spent = atimes * 3 + btimes
				}
			}
		}
	}
	return
}

func play2(a, b point, prize point) (win bool, spent int64) {
	var btimes = (prize.y * a.x - prize.x * a.y) / (a.x * b.y - b.x * a.y);
	var atimes = (prize.x - btimes * b.x) / a.x;
	win = (atimes * a.x + btimes * b.x == prize.x && atimes * a.y + btimes * b.y == prize.y)
	spent = atimes * 3 + btimes
	return
}

func main() {
	var scanner = bufio.NewScanner(os.Stdin)
	var totals [2]int64
	for scanner.Scan() {
		var a, b, prize point
		fmt.Sscanf(scanner.Text(), "Button A: X+%d, Y+%d", &a.x, &a.y)
		scanner.Scan()
		fmt.Sscanf(scanner.Text(), "Button B: X+%d, Y+%d", &b.x, &b.y)
		scanner.Scan()
		fmt.Sscanf(scanner.Text(), "Prize: X=%d, Y=%d", &prize.x, &prize.y)
		scanner.Scan()
		if win, spent := play(a, b, prize); win {
			totals[0] += spent
		}
		if win, spent := play2(a, b, point{prize.x + 10000000000000, prize.y + 10000000000000}); win {
			totals[1] += spent
		}
	}
	fmt.Println(totals)
}
