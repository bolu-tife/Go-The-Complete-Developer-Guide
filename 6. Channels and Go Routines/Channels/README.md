# Concurrency

Concurency is not parallelism

Concurrency relates to an application that is processing more than one task at the same time. It involves having multiple threads executing code. Concurrency is creates the illusion of parallelism, however actually the chunks of a task aren’t processed in a parallel way, but inside the application, there are more than one task is being processed at a time. Ideally, when a thread is blocked(go routine) another task is picked up while still awaiting the blocked task.

Parallelism involves multiple threads(go routines) running at the same time. It requires running multiple CPU's


## Go routines
A goroutine is a lightweight thread of execution.
we use the `go` to invoke new funtions in a go routine

Every go program has a (defuault) main routine. Every other one created are called child routines(the ones that we add the go keywords in front)

The main goroutine must be running for any other(child) goroutines to run. If the main goroutine terminates, then the program will exit and no other goroutine will run.

They are used to handle blocking code

## Channels
Channels are ways to communicate with go routines. They are the pipes that connect concurrent goroutines. You can send values into channels from one goroutine and receive those values into another goroutine. You can send the channel values into functions or saved them into variables
 
They are typed.

### Syntax
You make use of the `chan` keyword
make(chan dataType) eg `make(chan string)`

## Infinite for loop Syntax
`for {statements}`

## Function literals
They are unnamed function used to wrap a little chunk of code. They are called lambda( ruby, python) or anonymous functions(javascript, php) in other programming languages. 
eg
```
    func(x int) int { 
        x = x * x
        return x
    }(2) // returns 4 
```

## Note:
it's best not to reference the same variable in two different routines
instead rely on go as pass by value language and pass in it into a function (can be a function literal; Passing values[not reference types] into a function creates copies of that value)