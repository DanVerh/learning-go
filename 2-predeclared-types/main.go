package main

import "fmt"

func main() {
// Zero values
// Values of variables that are declared, but has no value assigned
 var intZero int
 fmt.Printf("Int zero value is %v\n", intZero)

// Literals
// Literals are any value written directly in the source code
	integerLiteral := 10
	integerUnderscoreLiteral := 10_234 // same as 10234 for readability
	integerByteLiteral := 0b0101 // also 0o - octat, 0x - hexademical
	floatLiteral := 10.234
	stringLiteral := "string\nnewline"
	runeLiteral := 's'
	runeMultibyteLiteral := '🪴'
	boolLiteral := true
	fmt.Println(
		"Literals examples:",
		integerLiteral, integerUnderscoreLiteral, integerByteLiteral, floatLiteral, stringLiteral, runeLiteral,runeMultibyteLiteral, boolLiteral,
	)
// Literals are untyped until they are used (in functions, vars or context)

// Types
// Type tells compiler how to store and operate on literals

// Boolean type
	var boolDefault bool
	fmt.Printf("Boolean default values is %v\n", boolDefault) // false is zero value

// Integer type
	// int8 (-128 to 127) uint8 (0 to 255) are minimal
	// int64 and uint64 are maximal sizes
	var int8Example int8
	fmt.Println(int8Example) // 0 is zero value
	var uintExample int // default int is int32, but not an alias to it, distinct type; uint is positive only
	int8Example, uintExample = 127, 250
	fmt.Printf("Overflow bug with sum of different int types with int8 conversion = %v\n", int8Example + int8(uintExample)) // answer is 121 which is incorrect
	fmt.Printf("Sum of different int types with correct int32 conversion = %v\n", int(int8Example) + uintExample)
	/*
	 - use just in int in most cases
	 - for specific requirements like network protocols mention size
	 - where multiple implementation type is required use generics (in past both int64 and uint64 functions)
	*/

// Float type
	/*
	Quick note on floats:
	- float value = sign × mantissa × 2^exponent
	- mantissa is an integer and point "floats" based on exponent
	- Floats are not exact because binary can exactly represent only fractions whose denominator is a power of 2. Decimal fractions like 0.1 require infinite binary digits, so they must be rounded to fit into finite bits, which causes small precision errors.
	*/
	var float32Example float32
	var float64Example float64
	fmt.Println(float32Example, float64Example) // zero value is 0
	// Floats are rarely used due to their unexact nature. It is better to not compare them too.


}