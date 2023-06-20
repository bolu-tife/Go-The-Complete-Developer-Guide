## Variables
A variable is a storage location, with a specific type and an associated name.
It is used to store data

### initialization
1. var variableName  dataType = "some word"
    var - keyword for declaring a variable
    card - name of the variable
    dataType - data type of hte variable eg int, string, bool, float64

2. variableName := "some word"
    here - go infers the type (coercion - check)

### assignment 
 variableName = "new word"


## Functions
A function is a block of code that performs a specific task

### Syntax
- func functionName(){} - returns void
- func functionName() dataType {} eg func newApple() string {return "newApple}


### Receiver Functions
- a function that belongs to an instance/type 
- func (d definedDataType) functionName() { }
- eg // func (d deck) print() {  d deck is a receiver - allows every variable of type deck have access to it ref './deck.go'


## Array vs slice
- array - fixed length of things
- slice - array that can grow and shrink eg cards := []string{'hello', 'world}
- both must contain same dataTypes

### Adding an element to a slice
cards = append(cards,"new word") -- this actually takes the elements in the old slice, appends the new item and creates a new slice with all the elements

## For loops
A loop is used to repeat a block of code.

### Syntax 
1. `for index, element := range elements{ statement(s) }`
2. `for initialization; condition; update { statement(s)} ` 


## Type conversion 
dataTypeYouWant(ValueToBeChanged)

## Tests
1. Tests files should end with _test.go
2. Test functions names should start follow PascalCase and should start with TestFunctionName

## Note 
1. Go is Strongly typed 
2. Go is not an Object Oriented Language
3. Go uses zero based indexing