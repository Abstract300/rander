package rander

import (
	"bufio"
	"io"
	"math/rand"
	"strings"
	"time"
)

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
