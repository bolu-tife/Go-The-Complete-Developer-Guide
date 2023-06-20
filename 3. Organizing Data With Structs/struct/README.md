# Structures (struct)
A struct (short for structure) is used to create a collection of members of different data types, into a single variable.

## Syntax 
type structName struct {
    attribute1 dataType
    attribute2 dataType
}

## Sample ways to declare
1. structName{"sample1", "sample2" } 
    - // this way isn't too encouraged
    - Golang automatically attaches attribute1 and attribute1 to sample1 and sample2 respectively - but what if there's a name swap 👀- chaos
2. structName{attribute1: "sample1", attribute2: "sample2" }
3. var variableName structName - all attributes would attain a default value

## Pointers
Go is a pass by value language. 
**Note** : - whenever we pass a value into a function, go copies that data and place it into a new computer memory

A Variable has 
- an address, where it is stored and 
- a value, what is stored

address -> value 
address = &variable 
value = *address (or the variable by default)

&variable - return the memory address of this variable
*pointer - returns the value of what is stored in that address

variable =  *(&variable) - returns the value stored in the memory address received

Note a *type eg *person in the './main.go' - means we are working on a type of pointer type 
*person means a type of pointer person

### Value types in Go 
- Examples - int, float, string, bool, struct
- they create copies of the value when dealing with them in functions 
- need to consider pointers when dealing with them

### Reference Types in Go
- Examples - slices, maps, channels, pointers, functions 
- they create copies but they are referencing the same value
- no need to consider pointers

When we create a slice - Go creates an array and a structure that records the length of the slice, capacity of the slice and a reference to the array