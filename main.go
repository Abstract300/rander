package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

var logger = log.New(os.Stdout, "LOGGER: ", log.Lshortfile)

const WORDCOUNT = 54763
const source = "dictionary"

func main() {
	rndSrc := rand.NewSource(time.Now().UnixNano())
	rndNum := rand.New(rndSrc)
	num1, num2 := rndNum.Int()%WORDCOUNT, rndNum.Int()%WORDCOUNT

	word, err := generateWord(num1, num2, source)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Your Random word: ", word)
}

type Word struct {
	First string
	Last  string
}

type WordError struct {
	msg string
	err error
}

func (we WordError) Error() string {
	return fmt.Sprintf("What: %s | How: %s", we.msg, we.err)
}

func generateWord(num1, num2 int, filename string) (string, error) {
	var word Word
	we := WordError{}
	wordList := make([]string, 0)

	file, err := readWords(filename)
	if err != nil {
		we.msg = "Couldn't fetch file data that was read."
		we.err = err
		return "", we
	}
	scanner := bufio.NewScanner(strings.NewReader(string(file)))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		wordList = append(wordList, scanner.Text())
	}

	word.First = wordList[num1]
	word.Last = wordList[num2]

	return word.First + word.Last, nil
}

type readWordsError struct {
	msg string
	err error
}

func (rwe readWordsError) Error() string {
	return fmt.Sprintf("What: %s | How: %s", rwe.msg, rwe.err)
}

func readWords(fileName string) ([]byte, error) {
	var b []byte
	rwerr := readWordsError{}
	f, err := os.Open(fileName)
	if err != nil {
		rwerr.err = err
		rwerr.msg = "Couldn't open the target file"
		return b, rwerr
	}

	fileInfo, err := f.Stat()
	if err != nil {
		rwerr.err = err
		rwerr.msg = "Cannot stat on the target file."
		return b, err
	}

	size := fileInfo.Size()

	b = make([]byte, size)

	_, err = f.Read(b)
	if err != nil {
		rwerr.err = err
		rwerr.msg = "Cannot read content from file."
		return b, err
	}

	return b, nil
}

func (w Word) String() string {
	return fmt.Sprintf("%s%s", w.First, w.Last)
}
