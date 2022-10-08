package reusable

import (
	"bufio"
	"os"
	"strconv"
)

type TestCaseFunction func(tc int)

var s *bufio.Scanner

func ReadTestCases(inputFile string, fn TestCaseFunction) {
	f, _ := os.Open(inputFile)

	s = bufio.NewScanner(f)
	s.Scan()
	tc, _ := strconv.Atoi(s.Text())

	for i := 0; i < tc; i++ {
		fn(i + 1)
	}
	defer f.Close()
}

func ScanFromBuffio() string {
	s.Scan()
	return s.Text()
}
