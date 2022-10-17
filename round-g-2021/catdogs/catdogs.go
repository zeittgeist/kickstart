package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/the-zeitgeist/kickstart/reusable"
)

func main() {
	fns := func(s *bufio.Scanner) func(tc int) {
		return func(tc int) {
			s.Scan()
			ds := strings.Fields(s.Text())
			d, _ := strconv.Atoi(ds[1])
			c, _ := strconv.Atoi(ds[2])
			m, _ := strconv.Atoi(ds[3])
			s.Scan()
			ss := strings.Split(s.Text(), "")

			ssc := ss[:]

			for i, a := range ss {
				if a == "C" {
					if c == 0 {
						break
					}
					c -= 1
				}
				if a == "D" {
					if d == 0 {
						break
					}
					d -= 1
					c += m
				}
				if i < len(ss) {
					ssc = ss[i+1:]
				}
			}

			res := "YES"

			if strings.Count(strings.Join(ssc, ""), "D") != 0 {
				res = "NO"
			}
			fmt.Printf("Case #%v: %v\n", tc+1, res)
		}
	}

	reusable.ReadTestCases("catdog.txt", fns)
	// s := bufio.NewScanner(os.Stdin)
	// s.Scan()
	// tc, _ := strconv.Atoi(s.Text())
	//
	// fn := fns(s)
	//
	// for i := 0; i < tc; i++ {
	// fn(i)
	// }
}
