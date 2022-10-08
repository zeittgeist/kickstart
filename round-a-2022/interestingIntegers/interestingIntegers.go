package main

import (
	"fmt"
)

func main() {
	fns := func(tc int) {
		// ab := reusable.ScanFromBuffio()
		// abs := strings.Fields(ab)
		// a, _ := strconv.Atoi(abs[0])
		// b, _ := strconv.Atoi(abs[1])
		//
		var a, b int
		fmt.Scanf("%d %d", &a, &b)

		count := 0

		for i := a; i <= b; i++ {
			if isInteresting(i) {
				count++
			}
		}

		fmt.Printf("Case #%d: %d\n", tc, count)
	}

	// reusable.ReadTestCases("rtc.txt", fns)
	// reusable.ReadTestCases("testCases.txt", fns)
	readStd(fns)
}

func readStd(fn func(tc int)) {
	var tc int
	fmt.Scanf("%d", &tc)

	for i := 0; i < tc; i++ {
		fn(i + 1)
	}
}

func isInteresting(v int) bool {
	var sum int
	prod := 1
	vc := v
	for vc > 0 && prod != 0 {
		mod := vc % 10
		vc /= 10
		prod *= mod
		sum += mod
	}

	return prod == 0 || prod%sum == 0
}
