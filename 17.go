package main

import "bufio"
import "os"
import "fmt"
import "strconv"
import "regexp"
import "math"

func run(program []int, inita int64, initb int64, initc int64) []int {
	var a, b, c, ip = inita, initb, initc, 0
	var output []int
	var combo = func(op int) int64 {
		switch op {
		case 0, 1, 2, 3: return int64(op)
		case 4: return a
		case 5: return b
		case 6: return c
		}
		panic("wrong operand" + string(op))
	}
	ip = 0
	for {
		if ip >= len(program) { break }
		var instr = program[ip]
		var op = program[ip + 1]
		switch instr {
		case 0: a = int64(float64(a) / (math.Pow(2, float64(combo(op)))))
		case 1: b = b ^ int64(op)
		case 2: b = combo(op) % 8
		case 3: if a > 0 { ip = op; ip -= 2 }
		case 4: b = b ^ c
		case 5: output = append(output, int(combo(op) % 8))
		case 6: b = int64(float64(a) / (math.Pow(2, float64(combo(op)))))
		case 7: c = int64(float64(a) / (math.Pow(2, float64(combo(op)))))
		}
		ip += 2
	}
	return output
}

func run2(program []int, target []int, inita int64) int64 {
	if len(target) == 0 {
		return inita
	}
	for digit := int64(0); digit < 8; digit++ {
		var a =  (inita << 3) | digit
		var b, c int64
		var out int
		var combo = func(op int) int64 {
			switch op {
			case 0, 1, 2, 3: return int64(op)
			case 4: return a
			case 5: return b
			case 6: return c
			}
			panic("wrong operand" + string(op))
		}
		for ip := 0; ip < len(program) - 2; ip += 2 {
			var instr = program[ip]
			var op = program[ip + 1]
			switch instr {
			case 1: b = b ^ int64(op)
			case 2: b = combo(op) % 8
			case 4: b = b ^ c
			case 5: out = int(combo(op) % 8)
			case 6: b = int64(float64(a) / (math.Pow(2, float64(combo(op)))))
			case 7: c = int64(float64(a) / (math.Pow(2, float64(combo(op)))))
			}
		}
		if out == target[len(target) - 1] {
			var prev = run2(program, target[:len(target)-1], a)
			if prev != -1 {
				return prev
			}
		}
	}
	return -1
}

func main() {
	var reNumber = regexp.MustCompile("\\d+")
	var scanner = bufio.NewScanner(os.Stdin)
	var a, b, c int64
	scanner.Scan()
	fmt.Sscanf(scanner.Text(), "Register A: %d",  &a)
	scanner.Scan()
	fmt.Sscanf(scanner.Text(),"Register B: %d",  &b)
	scanner.Scan()
	fmt.Sscanf(scanner.Text(),"Register C: %d",  &c)
	scanner.Scan()
	scanner.Scan()
	var program []int
	for _, s := range reNumber.FindAllString(scanner.Text(), -1) {
		n, _ := strconv.ParseInt(s, 10, 32)
		program = append(program, int(n))
	}
	var output = run(program, a, b, c)
	for i, v := range output {
		if i > 0 { fmt.Print(",") }
		fmt.Print(v)
	}
	fmt.Println()
	fmt.Println(run2(program, program, 0))
}

/*
2,4,1,1,7,5,0,3,1,4,4,0,5,5,3,0

0:
b = a % 8
b = b ^ 1
c = a >> b
b = b ^ 4
b = b ^ c
a = a >> 3
out b
if a > 0 goto 0
*/