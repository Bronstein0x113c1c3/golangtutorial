package main

import (
	"bufio"
	"bytes"
	"fmt"
)

type ByteCounter int
type WordCounter int
type LineCounter int

func (c *ByteCounter) Write(data []byte) (int, error) {
	*c += ByteCounter(len(data))
	return len(data), nil
}
func (w *WordCounter) Write(data []byte) (int, error) {
	words := 0
	the_scanner := bufio.NewScanner(bytes.NewBuffer(data))
	the_scanner.Split(bufio.ScanWords)
	for the_scanner.Scan() {
		words++
	}
	(*w) += WordCounter(words)
	return words, nil

}
func (l *LineCounter) Write(data []byte) (int, error) {
	lines := 0
	the_scanner := bufio.NewScanner(bytes.NewBuffer(data))
	for the_scanner.Scan() {
		lines++
	}
	(*l) += LineCounter(lines)
	return lines, nil
}

func main() {
	var w WordCounter
	var l LineCounter
	fmt.Fprintf(&w, "Hello, %s from the Death", "Bronstein")
	fmt.Fprintf(&l, "Hello, %s from the Death", "Bronstein")
	fmt.Println(w, " ", l)

}
