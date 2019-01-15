package main

import (
	"bufio"
	"fmt"
	"os"
)

var scanner = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func nextInt() int {
	scanner.Scan()
	r := 0
	f := 1
	for _, c := range scanner.Bytes() {
		if c == '-' {
			f = -1
			continue
		}
		r *= 10
		r += int(c - '0')
	}
	return r * f
}

func main() {
	scanner.Split(bufio.ScanWords)

	cache := make([]int, 101)
	for i := 0; i < 101; i++ {
		cache[i] = -1
	}

	caseNum := nextInt()
	for i := 0; i < caseNum; i++ {
		n := nextInt()
		fmt.Fprintln(out, tiling(n, cache))
	}

	out.Flush()
}

func tiling(n int, cache []int) int {
	if n == 0 || n == 1 {
		return 1
	}

	if cache[n] != -1 {
		return cache[n]
	}

	cache[n] = modAdd(tiling(n - 1, cache), tiling(n - 2, cache), 1000000007)
	return cache[n]
}

func modAdd(a, b int, num int) int {
	aP := a % num
	bP := b % num

	return (aP + bP) % num
}
