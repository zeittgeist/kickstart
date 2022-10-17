package main

import (
	"fmt"
)

func main() {
	fns := func(tc int) {
		var m, n, p int
		fmt.Scanf("%d %d %d", &m, &n, &p)

		s := make([][]int, m)

		var is int

		for i := 0; i < m; i++ {
			si := make([]int, n)
			s[i] = si

			for j := 0; j < n; j++ {
				fmt.Scanf("%d", &is)
				s[i][j] = is
			}
		}

		ms := 0

		for d := 0; d < n; d++ {
			msd := 0
			for i := 0; i < m; i++ {
				if i == p-1 {
					continue
				}

				if val := s[i][d] - s[p-1][d]; val > msd {
					msd = val
				}
			}

			ms += msd
		}

		fmt.Printf("Case #%d: %d\n", tc, ms)
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
