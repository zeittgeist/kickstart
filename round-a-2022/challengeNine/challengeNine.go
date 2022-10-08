package main

import (
	"fmt"
	"strings"
)

func main() {
	fns := func(tc int) {
		// ns := reusable.ScanFromBuffio(s)
		var ns string
		fmt.Scan(&ns)
		sum := 0
		for _, v := range ns {
			sum += int(v - '0')
		}
		mv := 9 - (sum % 9)

		if mv == 9 {
			fmt.Printf("Case #%v: %v\n", tc, fmt.Sprint(ns[:1], 0, ns[1:]))
			return
		}

		i := strings.IndexFunc(ns, func(r rune) bool {
			return int(r-'0') > mv
		})

		if i == -1 {
			fmt.Printf("Case #%v: %v\n", tc, fmt.Sprint(ns, mv))
			return
		}

		fmt.Printf("Case #%v: %v\n", tc, fmt.Sprint(ns[:i], mv, ns[i:]))
	}

	// reusable.ReadTestCases("tc1.txt", fns)
	readStd(fns)
}

func readStd(fn func(tc int)) {
	var tc int
	fmt.Scanf("%d", &tc)

	for i := 0; i < tc; i++ {
		fn(i + 1)
	}
}
