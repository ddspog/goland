package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/ddspog/goland/contracts"
)

var np = "\n\n"
var help = "Syntax: producer_consumer_sim [nProd int] [nCons int] [bufferSize int]\n" +
	"nProd -> number of producers to use.\n" +
	"nCons -> number of consumers" +
	"bufferSize -> size of the buffer."

// Simulate producer and consumer operations arount a limited buffer.
// The nProd and nCons argument limit the number of producers and consumers, bufferSize limit buffer.
func main() {
	defer contracts.ShowHelp(help)
	nProd, nCons, bufferSize := parseArguments()

	program(nProd, nCons, bufferSize)
}

// Verify if parameters are valid.
func parseArguments() (int, int, int) {
	contracts.Require(func() bool {
		return len(os.Args) > 3
	}, "The program receives number of producers, consumers and buffer size."+np+help)

	var nProd, errProd = strconv.Atoi(os.Args[1])
	var nCons, errCons = strconv.Atoi(os.Args[2])
	var bufferSize, errBuffer = strconv.Atoi(os.Args[3])

	contracts.Require(func() bool {
		if errProd == nil && errCons == nil && errBuffer == nil {
			return true
		}
		return false
	}, "Arguments must be int."+np+help)

	return nProd, nCons, bufferSize
}

// Struct to encapsulate buffer data, contain an id of producer and value produced.
type item struct {
	producer int
	value    int
}

// The program function build producers and consumers to work around a limited buffer.
// Producers produces elements ten times the size of buffer, and consumers prints the id of producer and value found
// on the buffer.
func program(nProd int, nCons int, bufferSize int) {
	var buffer = make(chan item, bufferSize)
	var done = make(chan bool, nProd)

	for i := 0; i < nProd; i++ {
		go produce(i, bufferSize*10, buffer, done)
	}

	for i := 0; i < nCons; i++ {
		go consume(buffer)
	}

	for i := 0; i < nProd; i++ {
		<-done
	}

	close(buffer)
}

// A function to produce items, putting id and value information.
// When it's done, report it to done channel.
func produce(id int, max int, buffer chan item, done chan bool) {
	for i := 0; i < max; i++ {
		produced := new(item)
		produced.producer = id
		produced.value = i
		buffer <- *produced
	}

	done <- true
}

// A function to produce items, printing their information.
// It finishes when no more item it's produced.
func consume(buffer chan item) {
	for elem := range buffer {
		fmt.Printf("%d %d\n", elem.producer, elem.value)
	}
}
