package main

import (
	"bufio"
	"fmt"
	"math"
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

	for i := 0; i < caseNum; i++ {
		n, m := nextInt(), nextInt()
		nums1 := make([]int, n)
		nums2 := make([]int, m)

		for j := 0; j < n; j++ {
			nums1[j] = nextInt()
		}

		for j := 0; j < m; j++ {
			nums2[j] = nextInt()
		}

		fmt.Fprintln(out, getJLIS(n,m, nums1, nums2))
	}

	out.Flush()
}

func getJLIS(n, m int, nums1, nums2 []int) int {
	cache := make([][]int, n)
	for i := 0; i < n; i++ {
		cache[i] = make([]int, m)
		for j := 0; j < m; j++ {
			cache[i][j] = -1
		}
	}

	longest := 0
	for i := -1; i < n; i++ {
		for j := -1; j < m; j++ {
			fmt.Printf("[Start] => i: %v, j: %v\n", i, j)
			longest = max(longest, jlis(nums1, nums2, i, j, cache))
		}
	}

	return longest
}

func jlis(nums1, nums2 []int, start1, start2 int, cache [][]int) int {
	indexA := start1 + 1
	indexB := start2 + 1

	if indexA >= len(nums1) || indexB >= len(nums2){
		return 0
	}

	fmt.Printf("indexA: %v, indexB: %v\n", indexA, indexB)

	if cache[indexA][indexB] != -1 {
		fmt.Printf("[DEBUG] cache hit! start1:%v, start2: %v\n", indexA, indexB)
		return cache[indexA][indexB]
	}

	var num1, num2 int
	if start1 == -1 {
		num1 = math.MinInt64
	} else {
		num1 = nums1[start1]
	}

	if start2 == -1 {
		num2 = math.MinInt64
	} else {
		num2 = nums2[start2]
	}

	maxNum := max(num1, num2)
	longest := 2

	for i := start1 + 1; i < len(nums1); i++ {
		if nums1[i] > maxNum {
			longest = max(longest, jlis(nums1, nums2, i, start2, cache) + 1)
			fmt.Printf("[DEBUG] longest:%v, i:%v , j:%v \n", longest, i, start2)
		}
	}

	for i := start2 + 1; i < len(nums2); i++ {
		if nums2[i] > maxNum {
			longest = max(longest, jlis(nums1, nums2, start1, i, cache) + 1)
			fmt.Printf("[DEBUG] longest:%v, i:%v, j:%v \n", longest, start1, i)
		}
	}

	cache[indexA][indexB] = longest
	return cache[indexA][indexB]
}

func max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}