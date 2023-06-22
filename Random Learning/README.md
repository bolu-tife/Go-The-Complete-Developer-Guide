# Random Knowledge 
**Naming returns Functions** - A return statement without arguments returns the named return values. This is known as a "naked" return. Naked return statements should be used only in short functions They can harm readability in longer functions.
```
    func functionName () (x,y int){

        return //returns x and y  and returns the zero value of the datatype - (x = 0)
}
```

**Guard clauses** - returning early from conditions/functions

**Short assignment statement (:=)** is only available inside a function

**Zero values** - Variables declared without an explicit initial value are given their zero value.

**Constants** - Constants are declared like variables, but with the const keyword. Constants can be character, string, boolean, or numeric values eg const Pi = 3.14