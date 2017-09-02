package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/ddspog/goland/collections"
	"github.com/ddspog/goland/contracts"
)

var np = "\n\n"
var help = "Syntax: translate_notes [level int] [list ...string]\n" +
	"level -> number of threads to use.\n" +
	"list -> latin musical notes like: Do, Re, Mi..."

// Maps musical notes on latin version to english using concurrency.
// The list argument it's a list of latin musical notes and level argument it's the level of concurrency.
// It outputs a mapping of the latin notes to english version.
func main() {
	defer contracts.ShowHelp(help)
	level, list := verifyParameters()

	var result = program(level, list)

	verifyResult(result)
	fmt.Println(result)
}

// Verify if parameters are valid.
func verifyParameters() (int, []string) {
	contracts.Require(func() bool {
		return len(os.Args) > 2
	}, "The program receives a concurrency level and a list of arguments."+np+help)

	var level, err = strconv.Atoi(os.Args[1])

	contracts.Require(func() bool {
		if err == nil {
			return true
		}

		fmt.Print(err.Error() + "\n")
		return false
	}, "Level argument must be an int."+np+help)
	contracts.Require(func() bool {
		var musicalNotes = collections.NewSet("Do", "Re", "Mi", "Fa", "Sol", "La", "Si")

		for _, note := range os.Args[2:] {
			if _, ok := musicalNotes[note]; !ok {
				return false
			}
		}

		return true
	}, "List must only contains latin musical notes."+np+help)

	var list = os.Args[2:]

	return level, list
}

// Verify if results are valid.
func verifyResult(result []string) {
	contracts.Ensure(func() bool {
		var musicalNotes = collections.NewSet("A", "B", "C", "D", "E", "F", "G")

		for _, note := range result {
			if _, ok := musicalNotes[note]; !ok {
				return false
			}
		}

		return true
	}, "Result must only contains english musical notes.")
}

// The program function maps musical notes on latin version to english using concurrency.
// The list argument it's a list of latin musical notes and level argument it's the level of concurrency.
// It returns a mapping of the latin notes to english version.
func program(level int, list []string) []string {
	return list
}
