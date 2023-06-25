package main

import (
	"fmt"
	"io"
	"os"
)

type Writer struct {
	w     io.Writer
	count int64
}

func main() {
	w, count := CountingWriter(os.Stdout)
	fmt.Fprintf(w, "Bonjour! %s \n", "Jonathan")
	fmt.Println(*count)
}

func (w *Writer) Write(p []byte) (int, error) {
	n, err := w.w.Write(p)
	w.count += int64(n)
	return n, err
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	wr := Writer{w, 0}
	return &wr, &wr.count
}
