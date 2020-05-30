package main

import (
	"math"

	"github.com/cheggaaa/pb/v3"
	"github.com/gammazero/workerpool"
)

func getRange(length int) (start int, end int) {
	var s int;
	var e int;

	for i := 0; i < length; i++ {
		s += int(math.Pow(26, float64(i)))
	}

	for i := 0; i < length + 1; i++ {
		e += int(math.Pow(26, float64(i)))
	}
	e--

	return s, e
}

func main() {
	start, end := getRange(2)
	bar := pb.StartNew(end - start + 1)
	wp := workerpool.New(5)
	for i := start; i <= end; i++ {
		i := i
		wp.Submit(func() {
			name := toLetters(i)
			checkName(name + "tide", func() {
				bar.Increment()
			})
		})
	}
	wp.StopWait()
	bar.Finish()
}
