// Module code is organied into packages
// Main is a package that starts go program
package main

// import whole package, can not limit some parts
import "fmt"

func main() {
	fmt.Println("Hello, World!")
}

// Following peace of code will not work, because of semicolon insertion rule
// check with go vet
/*
func semicolon_rule()
{
	fmt.Println("test")
}
*/
