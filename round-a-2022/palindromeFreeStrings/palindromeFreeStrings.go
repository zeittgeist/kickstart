package main

import (
	"fmt"
	"strings"
)

func main() {
	fns := func(tc int) {
		// n, _ := strconv.Atoi(reusable.ScanFromBuffio())
		// s := reusable.ScanFromBuffio()
		var n int
		var s string
		fmt.Scanf("%d", &n)
		fmt.Scan(&s)
		if n < 5 {
			fmt.Printf("Case #%d: POSSIBLE\n", tc)
			return
		}

		p := checkAllCobinations(s)
		if p {
			fmt.Printf("Case #%d: POSSIBLE\n", tc)
			return
		}

		fmt.Printf("Case #%d: IMPOSSIBLE\n", tc)
	}

	// reusable.ReadTestCases("rtc.txt", fns)
	// reusable.ReadTestCases("testCases.txt", fns)
	readStd(fns)
}

func checkAllCobinations(s string) bool {
	// fmt.Println(s)
	if containsSubpalindrome(s) {
		// fmt.Println(s, "has palindromes")
		return false
	}

	if !strings.Contains(s, "?") {
		// fmt.Println(s, "is a solution")
		return true
	}

	if i := strings.Index(s, "?"); i > 3 {
		if pv := binaryMapMemo[string(s[i-4:i])]; len(pv) == 1 {
			// fmt.Println("Limited tree for ", s)
			return checkAllCobinations(strings.Replace(s, "?", pv[0], 1))
		}
	}
	return checkAllCobinations(strings.Replace(s, "?", "0", 1)) || checkAllCobinations(strings.Replace(s, "?", "1", 1))

}

func readStd(fn func(tc int)) {
	var tc int
	fmt.Scanf("%d", &tc)

	for i := 0; i < tc; i++ {
		fn(i + 1)
	}
}

func containsSubpalindrome(s string) bool {
	bp := []string{
		"00000",
		"10001",
		"10101",
		"00100",
		"01010",
		"01110",
		"11011",
		"11111",
		"001100",
		"010010",
		"011110",
		"100001",
		"101101",
		"110011",
	}

	for _, v := range bp {
		if strings.Contains(s, v) {
			return true
		}
	}

	return false
}

func isPalindrome(s string) bool {
	i := 0
	j := len(s) - 1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}

	return true
}

var binaryMapMemo = map[string][]string{
	"0000": {"1"},
	"0001": {"0", "1"},
	"0010": {"1"},
	"0011": {"0", "1"},
	"0100": {"0", "1"},
	"0101": {"1"},
	"0110": {"0", "1"},
	"0111": {"1"},
	"1000": {"0"},
	"1001": {"0", "1"},
	"1010": {"0"},
	"1011": {"0", "1"},
	"1100": {"0", "1"},
	"1101": {"0"},
	"1110": {"0", "1"},
	"1111": {"0"},
}
