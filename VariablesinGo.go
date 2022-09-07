/*Variables
  we use var to define variables in Go*/

/*variables can be defined at package level or inside the function*/

package main

import "fmt"

func main() {
	var age = 20

	fmt.Println("Hello, World!")
}

/*we can also declare variables without predefined vale by 
var age int, or var name string etc.., but when declaring values
we use age:=10 we use the short variable declaration operator :=*/
/*Names are case sensitive*/


/* Use const to declare a constant value throughout
const age = 20*/

/*Declare multiple variables in one line*/
var ges, name
var ges, name = 10, "Bree"

/* also you can just declar without var keyword
Just do ages,name := 10, "Bree"*/
/* If we don't declare a value to a variable, then 
int is 0, and string is empty string
We can also declare as following*/

var vay int

var vay = 15

