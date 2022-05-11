package main

import "fmt"

var ints = map[string]int64{
	"first":  34,
	"second": 12,
}

// Initialize a map for the float values
var floats = map[string]float64{
	"first":  35.98,
	"second": 26.99,
}

// SumInts adds together the values of m

func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

// SumFloats adds together the values of m.

func Sumfloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}

// In this code, we have declared two functions to add together the values of a map and return the sum.
//   - SumFloats takes a map of string to float64 values
// SumInts takes a map of string to int64 values

/*
With generics, you can write one function here instead of two. Next, you’ll add a single generic
function for maps containing either integer or float values.
*/

// Adding a generic function to handle multiple types

/*To support values of either type, that single function will need a way to declare what types it
supports. Calling code, on the other hand, will need a way to specify whether it is calling with an
integer or float map.

To support this, you’ll write a function that declares type parameters in addition to its ordinary
function parameters. These type parameters make the function generic, enabling it to work with
arguments of different types. You’ll call the function with type arguments and ordinary function
arguments.

Each type parameter has a type constraint that acts as a kind of meta-type for the type parameter.
Each type constraint specifies the permissible type arguments that calling code can use for the
respective type parameter.

Keep in mind that a type parameter must support all the operations the generic code is performing
on it. For example, if your function’s code were to try to perform string operations (such as
indexing) on a type parameter whose constraint included numeric types, the code wouldn’t compile.


In the code you’re about to write, you’ll use a constraint that allows either integer or float
types.

*/

func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}

	return s
}

/*
   □ Specify for the K type parameter the type constraint comparable. Intended specifically for
     cases like these, the comparable constraint is predeclared in Go. It allows any type whose
     values may be used as an operand of the comparison operators == and !=. Go requires that
     map keys be comparable. So declaring K as comparable is necessary so you can use K as the
     key in the map variable. It also ensures that calling code uses an allowable type for map
     keys.
   □ Specify for the V type parameter a constraint that is a union of two types: int64 and
     float64. Using | specifies a union of the two types, meaning that this constraint allows
     either type. Either type will be permitted by the compiler as an argument in the calling
     code.
*/

func main() {

	fmt.Printf("Generic Sums: %v and %v\n",
		SumIntsOrFloats[string, int64](ints),
		SumIntsOrFloats[string, float64](floats))

}
