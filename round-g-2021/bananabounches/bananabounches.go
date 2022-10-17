package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/the-zeitgeist/kickstart/reusable"
)

func main() {
	fns := func(s *bufio.Scanner) func(tc int) {
		return func(tc int) {
			s.Scan()
			sf := strings.Fields(s.Text())
			k, _ := strconv.Atoi(sf[1])

			s.Scan()
			b := sort.StringSlice(strings.Fields(s.Text()))
			fmt.Println(b)

			bm := make([][]int, len(b))
			ut := make([]int, len(b))

			for i, bi := range b {
				ibi, _ := strconv.Atoi(bi)
				bm[i] = []int{ibi, i}
			}

			sort.SliceStable(bm, func(i, j int) bool {
				return bm[i][0] > bm[j][0]
			})
			fmt.Println(bm)

			var r int
			var t int

			for i := 0; i < len(b); i++ {
				if r == k {
					break
				}

				if r+bm[0][0] <= k {
					if validateUsedTree(&ut, bm[0][1]) {
						r += bm[0][0]
						t++
						fmt.Println("is valid", r)
					}
					bm = bm[1:]
				}
				fmt.Println(bm)
			}

			// fmt.Printf("bm: %+v\n", bm)
			fmt.Printf("Case #%v: ", tc)
			if r < k {
				fmt.Printf("%v\n", -1)
				return
			}
			fmt.Printf("%v\n", t)

			// fmt.Printf("r: %+v\n", r)
			//
			// fmt.Printf("n: %v\n", n)
			fmt.Printf("k: %v\n", k)
			fmt.Printf("ut: %v\n", ut)
		}
	}

	reusable.ReadTestCases("bananabounches.txt", fns)
	// readStd(fns)
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

func validateUsedTree(ut *[]int, i int) (isValid bool) {
	if len(*ut) < 2 {
		return true
	}

	if i == 0 {
		if (*ut)[1] == 0 || (*ut)[2] == 0 {
			return true
		}
	}

	if i == len(*ut)-1 {
		if (*ut)[i-1] == 0 || (*ut)[i-2] == 0 {
			return true
		}
	}

	if (*ut)[i+1] == 0 || (*ut)[i-1] == 0 {
		if (*ut)[i+1] == 0 && (*ut)[i-1] == 0 {
			return true
		}

		if (*ut)[i+1] == 1 {
			if len(*ut)-1 == i+1 {
				return true
			}

			if (*ut)[i+2] == 0 {
				return true
			}
		}

		if i-1 == 0 {
			return true
		}

		if (*ut)[i-2] == 0 {
			return true
		}
	}

	return isValid
}
