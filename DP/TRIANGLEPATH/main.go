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

func nextString() string {
	scanner.Scan()

	s := make([]byte, 0, 10)

	for _, c := range scanner.Bytes() {
		s = append(s, c)
	}

	return string(s[:])
}

func main() {
	scanner.Split(bufio.ScanWords)

	caseNum := nextInt()

	for i := 0; i < caseNum; i++ {
		n := nextInt()
		tri := make([][]int, n)

		for j := 1; j <= n; j++ {
			tri[j - 1] = make([]int, j)
			for k := 0; k < j; k++ {
				tri[j - 1][k] = nextInt()
			}
		}

		fmt.Fprintln(out, GetMaxPath(n, tri))
	}

	out.Flush()
}

func GetMaxPath(n int, tri [][]int) int {
	cache := make([][]int, n)
	for i := 1; i <= n; i++ {
		cache[i - 1] = make([]int, i)
	}

	return maxPath(0, 0, n, tri, cache)
}

func maxPath(y, x, n int, tri, cache [][]int) int {
	if y == n - 1 {
		return tri[y][x]
	}

	if x > y {
		return 0
	}

	if cache[y][x] != 0 {
		return cache[y][x]
	}

	cache[y][x] = max(maxPath(y + 1, x + 1, n, tri, cache), maxPath(y + 1, x, n, tri, cache)) + tri[y][x]
	return cache[y][x]
}

func max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}
