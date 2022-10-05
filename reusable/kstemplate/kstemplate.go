package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("No file name provided")
		os.Exit(1)
	}

	filename := os.Args[1]

	t := fmt.Sprint(
		"package main\n",
		"\n",
		"import (\n",
		"  \"bufio\"\n",
		"  \"os\"\n",
		"  \"fmt\"\n",
		"  \"github.com/the-zeitgeist/kickstart/reusable\"\n",
		")\n",
		"\n",
		"func main() {\n",
		"  fns := func(s *bufio.Scanner) func(tc int) {\n",
		"    return func(tc int) {\n",
		"      // your code goes here\n",
		"    }\n",
		"  }\n",
		"\n",
		"  reusable.ReadTestCases(\"yourtestfile.txt\", fns)\n",
		"  // readStd(fns)\n",
		"}\n",
		"\n",
		"func readStd(fns func(s *bufio.Scanner) func(tc int)) {\n",
		"  s := bufio.NewScanner(os.Stdin)\n",
		"  var tc int\n",
		"  fmt.Scanf(\"%d\", &tc)\n",
		"\n",
		"  fn := fns(s)\n",
		"\n",
		"  for i := 0; i < tc; i++ {\n",
		"          fn(i)\n",
		"  }\n",
		"}\n",
	)

	fp := fmt.Sprintf("%v.go", filename)

	ioutil.WriteFile(fp, []byte(t), 0776)
	fmt.Printf("Wrote file %v.go\n", filename)
}
