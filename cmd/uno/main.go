package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/fatih/color"
	"github.com/psykhi/uno/pkg/processor"
	"os"
)

func main() {
	printAll := flag.Bool("all", false, "Print all lines and highlight new lines")
	maxDiffRatio := flag.Float64("d", 0.3, "The maximum difference ratio between the input line and the other lines seen")
	flag.Parse()
	fmt.Println(maxDiffRatio)

	scanner := bufio.NewScanner(os.Stdin)
	p := processor.NewProcessor(*maxDiffRatio)
	l := processor.Line{}
	for scanner.Scan() {
		l.Input = scanner.Bytes()
		l.IsNew = false
		l = p.Process(l)
		if l.IsNew {
			if *printAll {
				color.Red("%s", l.Input)
			} else {
				fmt.Printf("%s\n", l.Input)
			}
			continue
		}
		if *printAll {
			fmt.Printf("%s\n", l.Input)
		}

	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}
}
