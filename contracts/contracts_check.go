// +build contracts

// Package contracts encapsulate Preconditions, Postconditions and Invariants to methods that will only apply,
// when the buld requires to check contracts via tags.
package contracts

// Require check if test evaluates as true, else it panics calling Precondition error via message.
func Require(test func() bool, message string) {
	if !test() {
		if message == "" {
			panic("Precondition error.")
		} else {
			panic("Pre error -> " + message)
		}
	}
}

// Ensure check if test evaluates as true, else it panics calling Postcondition error via message.
func Ensure(test func() bool, message string) {
	if !test() {
		if message == "" {
			panic("Postcondition error.")
		} else {
			panic("Post error -> " + message)
		}
	}
}

// Assure check if test evaluates as true, else it panics calling Invariant error via message.
func Assure(test func() bool, message string) {
	if !test() {
		if message == "" {
			panic("Invariant error.")
		} else {
			panic("Inv error -> " + message)
		}

	}
}

// ShowHelp recovers when program panics and show a message of help
func ShowHelp(helpMessage string) {}
