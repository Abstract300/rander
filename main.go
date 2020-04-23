package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"
	"time"
)

var logger = log.New(os.Stdout, "LOGGER: ", log.Lshortfile)

const SourceFile = "./dictionary"

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

	generator := NewWordGenerator(bufio.NewReader(f))
	words := generator.GenerateWords(wordCount)
	table := tabwriter.NewWriter(os.Stdout, 0, 8, 4, '\t', 0)
	defer table.Flush()
	fmt.Fprintf(table, "%s\t%s\t\n", "Index", "Word")
	fmt.Fprintf(table, "%s\t%s\t\n", "-----", "----")
	for i, j := range words {
		fmt.Fprintf(table, "%d\t%s\t\n", i, j)
	}
}

type WordGenerator struct {
	words []string
	rand  *rand.Rand
}

func (g *WordGenerator) GenerateWords(n int) []string {
	var words []string
	for i := 0; i < n; i++ {
		words = append(words, g.GenerateWord()+g.GenerateWord())
	}
	return words
}

func (g *WordGenerator) GenerateWord() string {
	return strings.Title(g.words[g.rand.Intn(len(g.words))])
}

func NewWordGenerator(reader io.Reader) *WordGenerator {
	words := make([]string, 0)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}
	return &WordGenerator{
		words: words,
		rand:  rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}
