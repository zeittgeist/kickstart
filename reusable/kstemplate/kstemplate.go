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
		"  fns := func(tc int) {\n",
		"      // your code goes here\n",
		"  }\n",
		"\n",
		"  reusable.ReadTestCases(\"yourtestfile.txt\", fns)\n",
		"  // readStd(fns)\n",
		"}\n",
		"\n",
		"func readStd(fn func(tc int)) {\n",
		"  var tc int\n",
		"  fmt.Scanf(\"%d\", &tc)\n",
		"\n",
		"  for i := 0; i < tc; i++ {\n",
		"          fn(i + 1)\n",
		"  }\n",
		"}\n",
	)

	fp := fmt.Sprintf("%s.go", filename)

	ioutil.WriteFile(fp, []byte(t), 0776)
	fmt.Printf("Wrote file %s.go\n", filename)
}
