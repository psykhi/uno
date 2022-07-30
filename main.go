package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/fatih/color"
	"github.com/psykhi/uno/pkg/processor"
	"log"
	"os"
	"strings"
)

func main() {
	printAll := flag.Bool("all", false, "Print all lines and highlight new lines in red (if the terminal supports it)")
	maxDiffRatio := flag.Float64("d", 0.2, "The maximum difference ratio between the input line and the other lines seen (between 0 and 1)")
	patterns := flag.Bool("p", false, "Show log patterns")

	flag.Parse()

	if *maxDiffRatio < 0 || *maxDiffRatio > 1 {
		log.Fatalf("Distance ratio must be between 0 and 1")
	}

	// Read from stdin or from file
	f := os.Stdin
	if len(os.Args) >= 2 {
		var err error
		lastArg := os.Args[len(os.Args)-1]
		if !strings.HasPrefix(lastArg, "-") {
			f, err = os.Open(lastArg)
			if err != nil {
				panic(err)
			}
			defer f.Close()
		}
	}
	scanner := bufio.NewScanner(f)

	p := processor.NewProcessor(*maxDiffRatio)
	l := processor.Line{}

	for scanner.Scan() {
		l.Input = scanner.Bytes()
		l.IsNew = false
		l = p.Process(l)
		toPrint := l.Input
		if *patterns {
			toPrint = []byte(strings.Join(l.Tokens, ""))
		}
		if l.IsNew {
			if *printAll {
				color.Red("%s", toPrint)
			} else {
				fmt.Printf("%s\n", toPrint)
			}
			continue
		}
		if *printAll {
			fmt.Printf("%s\n", toPrint)
		}
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}
}
