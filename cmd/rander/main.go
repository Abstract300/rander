package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"text/tabwriter"

	"github.com/abstract300/rander"
)

var logger = log.New(os.Stdout, "LOGGER: ", log.Lshortfile)

const SourceFile = "../../assets/dictionary"

func main() {
	flag.Parse()
	if len(flag.Args()) > 1 || len(flag.Args()) <= 0 {
		logger.Fatal("Illegal arguments.")
	}
	wordCount, err := strconv.Atoi(flag.Arg(0))
	if err != nil {
		logger.Fatal("Argument isn't an integer")
	}

	f, err := os.Open(SourceFile)
	if err != nil {
		logger.Fatal("Could not open source file", err)
	}

	generator := rander.NewWordGenerator(bufio.NewReader(f))
	words := generator.GenerateWords(wordCount)
	table := tabwriter.NewWriter(os.Stdout, 0, 8, 4, '\t', 0)
	defer table.Flush()
	fmt.Fprintf(table, "%s\t%s\t\n", "Index", "Word")
	fmt.Fprintf(table, "%s\t%s\t\n", "-----", "----")
	for i, j := range words {
		fmt.Fprintf(table, "%d\t%s\t\n", i, j)
	}
}
