package main

import "bufio"
import "os"
import "regexp"
import "fmt"

func main() {
	var cmd = regexp.MustCompile("mul\\(\\d+,\\d+\\)|do\\(\\)|don't\\(\\)")
	var scanner = bufio.NewScanner(os.Stdin)
	var muls [2]int
	var do = true
	for scanner.Scan() {
		for _, s := range cmd.FindAllString(scanner.Text(), -1) {
			if s[0:3] == "mul" {
				var n1, n2 int
				fmt.Sscanf(s, "mul(%d,%d)", &n1, &n2)
				muls[0] += n1 * n2
				if do { muls[1] += n1 * n2 }
			} else {
				do = s == "do()"
			}
		}
	}
	fmt.Println(muls)
}
