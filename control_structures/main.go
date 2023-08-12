package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var totalLC, totalWC, totalCC int

	for _, fname := range os.Args[1:] {

		var lc, wc, cc int

		file, err := os.Open(fname)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		scan := bufio.NewScanner(file)

		for scan.Scan() {
			s := scan.Text()

			wc += len(strings.Fields(s))
			cc += len(s)
			lc++
		}

		fmt.Printf("%4d %4d %4d %s\n", lc, wc, cc, fname)

		totalLC += lc
		totalWC += wc
		totalCC += cc

		file.Close()
	}

	fmt.Printf("%4d %4d %4d %s\n", totalLC, totalWC, totalCC, "total")
}
