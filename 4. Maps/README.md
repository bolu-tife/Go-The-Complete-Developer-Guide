# Maps
Golang Maps is a collection of unordered pairs of key-value. It is widely used because it provides fast lookups and values that can retrieve, update or delete with the help of keys. It is similar to objects in Javascript and dictionaries in Python

## Note 
- Map keys must have the same types
- Map values must have the same type
- Maps are reference types

## How to create a map

1.  variableName := map[keyDataType]valueDataType{} eg `colors := map[string]string{ }` , `colors := map[string]int{ "red": 1, }`
2.  var variableName map[keyDataType]valueDataType eg `var colors map[string]string`
3.  variableName := make(map[keyDataType]valueDataType) eg `colors := make(map[string]string)`

## Adding to a map
variableName[key] = value eg `colors["red"] = "RED"`

## Removing a key
`delete(colors, "red")`

## Maps vs Structs
- you can mix and match key and values in structs
- map keys must be the same type unlike struct
- map values must be the same type unlike struct
- you cannot iterate over properties in a struct
- structs are value types while maps are reference types

