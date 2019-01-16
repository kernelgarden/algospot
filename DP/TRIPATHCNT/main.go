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

	caseNum := nextInt()
	for c := 0; c < caseNum; c++ {
		n := nextInt()

		cache := make([][]int, n)
		pathCache := make([][]int, n)
		tri := make([][]int, n)
		for i := 0; i < n; i++ {
			tri[i] = make([]int, i + 1)
			cache[i] = make([]int, i + 1)
			pathCache[i] = make([]int, i + 1)
			for j := 0; j <= i; j++ {
				tri[i][j] = nextInt()
				cache[i][j] = -1
				pathCache[i][j] = -1
			}
		}

		path(tri, cache, 0, 0)
		fmt.Fprintln(out, maxCnt(cache, pathCache, 0, 0))
	}

	out.Flush()
}

func path(tri, cache [][]int, y, x int) int {
	if y + 1 == len(tri) {
		cache[y][x] = tri[y][x]
		return cache[y][x]
	}

	if cache[y][x] != -1 {
		return cache[y][x]
	}

	cache[y][x] = tri[y][x] + max(path(tri, cache, y + 1, x), path(tri, cache, y + 1, x + 1))
	return cache[y][x]
}

// 경로를 다시 셀 때도 메모이제이션을 적용시켜야 훨씬 빠르다.
// 경로의 갯수를 셀 때도 탐색의 경우의 수가 매우 크기 때문
func maxCnt(pathMap, cache [][]int, y, x int) int {
	if y + 1 == len(pathMap) {
		return 1
	}

	if cache[y][x] != -1 {
		return cache[y][x]
	}

	next1 := pathMap[y + 1][x]
	next2 := pathMap[y + 1][x + 1]

	var pathNum int
	if next1 > next2 {
		pathNum = maxCnt(pathMap, cache, y + 1, x)
	} else if next1 < next2 {
		pathNum = maxCnt(pathMap, cache, y + 1, x + 1)
	} else {
		pathNum = maxCnt(pathMap, cache, y + 1, x) + maxCnt(pathMap, cache, y + 1, x + 1)
	}

	cache[y][x] = pathNum
	return cache[y][x]
}

func max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}
