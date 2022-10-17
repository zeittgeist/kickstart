package main

import (
	"fmt"
	"math"
)

func main() {
	fns := func(tc int) {
		var r, a, b int
		fmt.Scanf("%d %d %d", &r, &a, &b)
		fmt.Printf("Case #%d: %.6f\n", tc, calculateArea(r, a, b, true))
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

func calculateArea(r, a, b int, right bool) float64 {
	if r == 0 {
		return 0
	}

	nr := int(r / b)

	if right {
		nr = r * a
	}

	return math.Pi*math.Pow(float64(r), 2.0) + calculateArea(nr, a, b, !right)
}
