package main

import "bufio"
import "os"
import "regexp"
import "strconv"
import "fmt"

func main() {
	var reMul = regexp.MustCompile("mul\\((\\d+),(\\d+)\\)|do\\(\\)|don't\\(\\)")
	var scanner = bufio.NewScanner(os.Stdin)
	var muls, dmuls int64
	var do = true
	for scanner.Scan() {
		for _, sub := range reMul.FindAllStringSubmatch(scanner.Text(), -1) {
			if sub[0][0:3] == "mul" {
				var n1, _ = strconv.ParseInt(sub[1], 10, 32)
				var n2, _ = strconv.ParseInt(sub[2], 10, 32)
				muls += n1 * n2
				if do { dmuls += n1 * n2 }
			} else {
				do = sub[0] == "do()"
			}
		}
	}
	fmt.Println(muls, dmuls)
}
