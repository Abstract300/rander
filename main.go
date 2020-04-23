package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

var logger = log.New(os.Stdout, "LOGGER: ", log.Lshortfile)

const SourceFile = "./dictionary"

func main() {
	f, err := os.Open(SourceFile)
	if err != nil {
		logger.Fatal("Could not open source file", err)
	}
	generator := NewWordGenerator(bufio.NewReader(f))
	word := generator.GenerateWord() + generator.GenerateWord()
	fmt.Println("Your random word: ", word)
}

type WordGenerator struct {
	words []string
	rand *rand.Rand
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
		rand: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}
