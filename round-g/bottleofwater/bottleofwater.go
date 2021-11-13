package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	fns := func(s *bufio.Scanner) func(tc int) {
		return func(tc int) {
			s.Scan()
			k, _ := strconv.Atoi(s.Text())
			var xs sort.IntSlice
			var ys sort.IntSlice

			for i := 0; i < k; i++ {
				s.Scan()
				st := strings.Fields(s.Text())

				is := make([]int, len(st))

				for i, c := range st {
					is[i], _ = strconv.Atoi(c)
				}

				xs = append(xs, is[0], is[2])
				ys = append(ys, is[1], is[3])
			}

			xs.Sort()
			ys.Sort()

			fmt.Printf("Case #%v: %v %v\n", tc+1, xs[xs.Len()/2-1], ys[ys.Len()/2-1])
		}
	}

	// reusable.ReadTestCases("bottleofwater.txt", fns)
	readStd(fns)
}

func readStd(fns func(s *bufio.Scanner) func(tc int)) {
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	tc, _ := strconv.Atoi(s.Text())

	fn := fns(s)

	for i := 0; i < tc; i++ {
		fn(i)
	}
}
