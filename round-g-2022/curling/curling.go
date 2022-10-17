package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fns := func(tc int) {
		var rs, rh int
		var n, m int

		fmt.Scanf("%d %d", &rs, &rh)

		fmt.Scanf("%d", &n)
		ns := readPoints(n, rs, rh, "red")

		fmt.Scanf("%d", &m)
		ms := readPoints(m, rs, rh, "yellow")

		nms := append(ns, ms...)

		if len(nms) == 0 {
			fmt.Printf("Case #%d: %d %d\n", tc, 0, 0)
			return
		}

		sort.Slice(nms, func(i, j int) bool {
			return nms[i].Distance < nms[j].Distance
		})

		t := nms[0].team
		ps := 0

		for _, v := range nms {
			if v.team != t {
				break
			}

			ps += 1
		}

		if t == "red" {
			fmt.Printf("Case #%d: %d %d\n", tc, ps, 0)
			return
		}

		fmt.Printf("Case #%d: %d %d\n", tc, 0, ps)
	}

	readStd(fns)
}

func euclideanDistance(x, y, r int) float64 {
	return math.Sqrt(float64(x*x+y*y)) - float64(r)
}

func readPoints(l, rs, rh int, t string) []PointMap {
	pm := make([]PointMap, 0, l)
	var x, y int

	for i := 0; i < l; i++ {
		fmt.Scanf("%d %d", &x, &y)

		d := euclideanDistance(x, y, rs)
		if d > float64(rh) {
			continue
		}

		pm = append(pm, PointMap{d, []int{x, y}, t})
	}

	return pm
}

type PointMap struct {
	Distance float64
	Point    []int
	team     string
}

func readStd(fn func(tc int)) {
	var tc int
	fmt.Scanf("%d", &tc)

	for i := 0; i < tc; i++ {
		fn(i + 1)
	}
}
