package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"

	"github.com/ddspog/goland/sort"
)

func main() {
	var testArray1, testArray2 = generateArray(), generateArray()
	var resultBubble = measureSorts("BubbleSort", func() []int {
		sort.BubbleSort(testArray1)
		return testArray1
	})
	var resultParallel = measureSorts("ParallelSort", func() []int {
		sort.Parallel(testArray2, runtime.NumCPU())
		return testArray2
	})
	fmt.Println(resultBubble)
	fmt.Println(resultParallel)
}

func generateArray() []int {
	var array = make([]int, 100)
	for i := range array {
		array[i] = rand.Int() % 200
	}
	return array
}

func measureSorts(name string, method func() []int) string {
	var start = time.Now()
	method()
	var elapsed = time.Since(start)

	return fmt.Sprintf("Method: %q, Finished in %s\n", name, elapsed)
}
