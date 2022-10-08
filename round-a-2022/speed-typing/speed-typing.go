package main

import (
	"fmt"
)

func main() {
	fns := func(tc int) {
		var i, p string

		fmt.Scan(&i)
		fmt.Scan(&p)

		ii := []rune(i)
		pp := []rune(p)

		var n, m int

		for n < len(ii) && m < len(pp) {
			if ii[n] == pp[m] {
				n++
				m++
				continue
			}

			m++
		}

		if n == len(ii) {
			fmt.Printf("Case #%v: %v\n", tc, len(pp)-len(ii))
			return
		}

		fmt.Printf("Case #%v: IMPOSSIBLE\n", tc)
	}

	// reusable.ReadTestCases("test-case2.txt", fns)
	// reusable.ReadTestCases("test-case.txt", fns)
	readStd(fns)
}

func readStd(fn func(tc int)) {
	var tc int
	fmt.Scanf("%d", &tc)

	for i := 0; i < tc; i++ {
		fn(i + 1)
	}
}
