package main

import (
	"fmt"
	"math"
)

func main() {
	// ZERO VALUES
	// Values of variables that are declared, but has no value assigned
	var intZero int
	fmt.Printf("Int zero value is %v\n", intZero)

	// LITERALS
	// Literals are any value written directly in the source code
	integerLiteral := 10
	integerUnderscoreLiteral := 10_234 // same as 10234 for readability
	integerByteLiteral := 0b0101       // also 0o - octat, 0x - hexademical
	floatLiteral := 10.234
	stringLiteral := "string\nnewline"
	runeLiteral := 's'
	runeMultibyteLiteral := '🪴'
	boolLiteral := true
	fmt.Println(
		"Literals examples:",
		integerLiteral, integerUnderscoreLiteral, integerByteLiteral, floatLiteral, stringLiteral, runeLiteral, runeMultibyteLiteral, boolLiteral,
	)
	// Literals are untyped until they are used (in functions, vars or context)

	// TYPES
	// Type tells compiler how to store and operate on literals

	// Boolean type
	var boolDefault bool
	fmt.Printf("Boolean default values is %v\n", boolDefault) // false is zero value

	// Integer type
	// int8 (-128 to 127) uint8 (0 to 255) are minimal
	// int64 and uint64 are maximal sizes
	var int8Example int8
	fmt.Println(int8Example) // 0 is zero value
	var uintExample int      // default int is int32, but not an alias to it, distinct type; uint is positive only
	int8Example, uintExample = 127, 250
	fmt.Printf("Overflow bug with sum of different int types with int8 conversion = %v\n", int8Example+int8(uintExample)) // answer is 121 which is incorrect
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
	fmt.Println("Both floats default values are", float32Example, float64Example) // zero value is 0
	// Floats are rarely used due to their unexact nature. It is better to not compare them too.

	// String and runes
	// Strings in Go are immutable (can not be changed) sequence of bytes, usually with UTF-8 encoded text
	s := "Go"
	fmt.Printf("%v length is %v\n", s, len(s)) // 2
	s = "🙂"
	fmt.Printf("%v length is %v\n", s, len(s)) // 4
	// Sequence of bytes, not characters

	// Rune is a type for a single code point which is an aliast to int32
	s = "🙂"
	fmt.Printf("Index of string provides bytes like index 0 here 🙂: %v\n", s[0])
	fmt.Println("Range iteration with proper value format provides runes:")
	for _, r := range "á🙂" {
		fmt.Printf("%c\n", r)
	}

	// Type conversions
	// Go does not allow automatic type promotion, only explicit conversion
	fmt.Printf("Sum of different int types with correct int32 conversion = %v\n", int(int8Example)+uintExample)
	// No type can be converted to a bool
	// As literals are untyped you can combine any compatible type, so they will be proceeded before compilation
	var untypedSum float64 = 200.5 * 2
	fmt.Println("Multiplication of int and float using untyped literal:", untypedSum)

	// VARIABLE DECLARATION
	// _ - intenionally unused variable
	var _ int = 10 // best for initializing zero value and redifining default literal type
	var _ = 15     // package level with default literal type of the value
	z := 10        // same as previous one; most used, works only on function level
	fmt.Println("z =", z)
	z = 12 // reassign value only with =
	fmt.Println("z =", z)

	// Multiple variable assignments at once
	var _, _ int = 1, 2
	var _, _ = 0, "hello"

	// Put var before multiple declarations
	var (
		_        = 15
		_ string = "hey"
	)

	// Const
	// Immutable variable declaration (usually done on package level)
	const constExample1 int = 13
	const (
		constExample2 = 14
		constExample3 = constExample1 + 10 // but z + 10 will cause an error, because consts are calculated at compille time
	)
	fmt.Println(constExample3)

	// Unused variables
	// Every variable in go needs to be read at least once
	readThisVar := 15
	fmt.Println(readThisVar)
	// Package level variables and consts are exceptions as they are calculated at compile time and if not used, will be included in compiled binary

	// Naming
	// snakeCase, short names, k and v in for range, i and j in for loop, Capital beginning only for items accessible outside the package
	// can not use short variable names = code is overcomplicated

	// EXCERCISES
	fmt.Println("Excercise 1:")
	var i1 int = 20
	var j1 float32 = float32(i1)
	fmt.Println(i1, j1)

	fmt.Println("Excercise 2:")
	const value = 2
	var i2 int = value
	var f2 float32 = value
	fmt.Println(i2, f2)

	fmt.Println("Excercise 3:")
	var b3 byte = math.MaxUint8
	var smallI int32 = math.MaxInt32
	var bigI uint64 = math.MaxInt64
	b3 += 1
	smallI += 1
	bigI += 1
	fmt.Println(b3, smallI, bigI)
}
