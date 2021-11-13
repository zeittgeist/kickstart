package reusable

import (
	"bufio"
	"os"
	"strconv"
)

type TestCaseFunction func(s *bufio.Scanner) func(tc int)

func ReadTestCases(inputFile string, tcf TestCaseFunction) {
	f, _ := os.Open(inputFile)
	s := bufio.NewScanner(f)
	s.Scan()
	tc, _ := strconv.Atoi(s.Text())

	fn := tcf(s)

	for i := 0; i < tc; i++ {
		fn(i)
	}
}
