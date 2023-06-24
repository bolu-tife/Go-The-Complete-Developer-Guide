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

**For Loops**

`for init; condition; increment{ statement }`
it can also be `for ; condition; {statement}`

For serves as while in go 
`for condition {statement}`

omiting the codition makes it an **infinite loop**
`for {}` 

**If Statements**

if statement can start with a short statement to execute before the condition.

```
    if v := math.Pow(x, n); v < lim {
		return v
	}
```

**Defer Statements**

A defer statement defers the execution of a function until the surrounding function returns. It also returns things in a LIFO{Last in First Out} order

