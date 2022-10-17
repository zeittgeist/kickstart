package main

import (
	"fmt"
	"math"
)

func main() {
	fns := func(tc int) {
		var n int

		fmt.Scanf("%d", &n)
		var c = []string{}

		countPalindromicFactors(n, &c)
		fmt.Printf("%#v", c)
		fmt.Printf("Case #%d: %d\n", tc, len(c))

	}

	readStd(fns)
}

func readStd(fn func(tc int)) {
	var tc int
	fmt.Scanf("%d", &tc)

	for i := 0; i < tc; i++ {
		fn(i + 1)
	}
}

func countPalindromicFactors(n int, c *[]string) {
	for k := 1; k <= int(math.Sqrt(float64(n))); k++ {
		if n%k == 0 {
			ns := fmt.Sprint(n / k)
			ks := fmt.Sprint(k)
			if isPalindrome(ns) {
				*c = append(*c, ns)
			}

			if ns == ks {
				continue
			}

			if isPalindrome(ks) {
				*c = append(*c, ks)
			}
		}
	}
}

func isPalindrome(s string) bool {
	i := 0
	j := len(s) - 1
	for i < j {
		if s[i] != s[j] {
			return false
		}

		i++
		j--
	}

	return true
}
