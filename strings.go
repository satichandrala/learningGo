/* strings are defined using double quotes*/
package main

import "fmt"

func main() {
	var name = "pruefung"
	// [0:2] letters in those positions
	fmt.Println(len(name), name[0:2])

	// creating a copy of the string name
	var newstring = name[:]
	//assigning a string to a new variable
	// Strings are immutable : Even if we assign new value to first, second is still going to be test
	var first = "test"
	var second = first
	first = "another test"
	var word = first + " " + second
	fmt.Println("newstring:", newstring, "first:", first, "second:", second, "concatenated first and second:", word)
}
