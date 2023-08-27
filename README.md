# goTutorial

This repository was created for my own benefit so that I can better understand the intricacies of how Go operates.
However, I hope that anyone that comes across it can get some sort of benefit from the processes I go through to understand
the different concepts and inner workings within GoLang.

I approach Go coming from a background in C/C++, so I will most likely be comparing the functionalities that work similarly
or highlight their differences.

Completed:
  * Slices
  * Maps
  * Interfaces
  * iota
  * dots
  * Concurrency
    * Rate limiter with event-based architecture
  
Interested/To Do:
  * Pointers
  * Polymorphism - more interfaces
  
# Some Notes:

## Formatting
* Each package should have a block comment describing its purpose
* Every exported method should have a doc comment
* Start comment with name of function
* Run Gofmt to format files

## Interfaces
* An interface is a custom data type that can take the form of any other data type that implements all of its methods.
* This is Go's form if polymorphism. 

## Printing
* Go uses the fmt package to do all of its printing (i.e. fmt.Print)
* Print, Printf, and Println are all available to print and format strings
    * As well as Fprintf, Sprint siblings all have same variations for printing to io.Writer compatible type. Sprint to print to string
* Go string formatting: https://golang.org/pkg/fmt/
* Custom default value by implementing "String() string" interface

## Functions
* The final parameter in a function can prepended by the … operator. This makes the function variadic, meaning that the final parameter can be satisfied by 0 or more arguments
* Constants
* Constants can be defined using the const keyword and parentheses
    * Also, var() works as well
*Incremental constants can be created automatically with iota

## Misc

* Consecutive function parameters can be identified using a single type at end
* Functions can return more than one argument!!
* If variables are named in return statement, they act as variable declarations
    * Return without arguments can be used - "naked return"

* Receivers can be added to functions to associate them to a specific type or struct
* Pointers can be added to these receivers to modify the values of the type/struct that called the method
* Although function parameters that call for a function must be passed an address, receivers implicitly convert a variable v to &v if called.

