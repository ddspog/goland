// +build !contracts

// Package contracts encapsulate Preconditions, Postconditions and Invariants to methods that will only apply,
// when the buld requires to check contracts via tags.
package contracts

import "fmt"

// Require check if test evaluates as true, else it panics calling Precondition error via message.
func Require(test func() bool, message string) {}

// Ensure check if test evaluates as true, else it panics calling Postcondition error via message.
func Ensure(test func() bool, message string) {}

// Assure check if test evaluates as true, else it panics calling Invariant error via message.
func Assure(test func() bool, message string) {}

// ShowHelp recovers when program panics and show a message of help
func ShowHelp(helpMessage string) {
	if r := recover(); r != nil {
		if helpMessage == "" {
			fmt.Println("An error has ocurred.")
		} else {
			fmt.Println("An error has ocurred.\n\n" + helpMessage)
		}
	}
}
