package main

import (
	"bufio"
	"fmt"
	"os"
)

var scanner = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

var MOD = 1000000007

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
	for c := 0 ; c < caseNum; c++ {
		n := nextInt()
		asymCache := make([]int, n + 1)
		tilingCache := make([]int, n + 1)
		for i := 0; i <= n; i++ {
			asymCache[i] = -1
			tilingCache[i] = -1
		}
		fmt.Fprintln(out, asymtiling(n, asymCache, tilingCache))
	}

	out.Flush()
}

func asymtiling(w int, cache []int, tilingCache []int) int {
	// width 가 2라면 어떤 경우에도 대칭이 발생한다.
	if w <= 2 {
		return 0
	}

	if cache[w] != -1 {
		return cache[w]
	}

	cache[w] = (2 * tiling(w - 3, tilingCache)) % MOD

	cache[w] = modAdd(cache[w], asymtiling(w - 2, cache, tilingCache), MOD)
	cache[w] = modAdd(cache[w], asymtiling(w - 4, cache, tilingCache), MOD)

	return cache[w]
}

func tiling(w int, cache []int) int {
	if w == 0 || w == 1 {
		return 1
	}

	if cache[w] != -1 {
		return cache[w]
	}

	cache[w] = modAdd(tiling(w - 1, cache), tiling(w - 2, cache), MOD)
	return cache[w]
}

func modAdd(a, b, mod int) int {
	aP := a % mod
	bP := b % mod

	return (aP + bP) % mod
}
