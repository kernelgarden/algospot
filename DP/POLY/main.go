package main

import (
	"bufio"
	"fmt"
	"os"
)

var scanner = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

var MOD = 10000000

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
		sum := 0
		n := nextInt()
		cache := make([][]int, n + 1)
		for i := 0; i < n + 1; i++ {
			cache[i] = make([]int, n + 1)
			for j := 0; j < n + 1; j++ {
				cache[i][j] = -1
			}
		}

		for i := 1; i <= n; i++ {
			sum = modAdd(sum, poly(n, i, cache), MOD)
		}

		fmt.Fprintln(out, sum)
	}

	out.Flush()
}

func poly(n, first int, cache [][]int) int {
	// 쓸데 없는 기저사례를 잡고 있어서 헤맸다.
	// n과 first의 범위는 체크할 필요가 없다. 어차피 first는 조건문에서 걸러지고,
	// n이 0인 경우는 first == n 이 조건에 의해 체크할 필요가 없어진다.
	if first == n {
		return 1
	}

	if cache[n][first] != -1 {
		//fmt.Printf("n: %v, first: %v, val: %v\n", n, first, cache[n][first])
		return cache[n][first]
	}

	sum := 0
	// 첫번쨰 라인과 두번째 라인의 조합 가능한 폴리의 수는 (second + (first - 1))
	for second := 1; second <= (n - first); second++ {
		subSum := poly(n-first, second, cache)
		// 경우를 곱해야하는 경우를 놓쳐서 헤맸다.
		// 현재 함수 내에서 첫번째 라인과 둘쨰 라인의 조합의 수 * 다음 재귀로 만들어지는 조합의 수를 해야하는데 더했음..
		sum = modAdd(sum, (second + first - 1) * subSum, MOD)
	}

	cache[n][first] = sum
	//fmt.Printf("n: %v, first: %v, val: %v\n", n, first, cache[n][first])
	return cache[n][first]
}

func modAdd(a, b, mod int) int {
	aP := a % mod
	bP := b % mod

	return (aP + bP) % mod
}
