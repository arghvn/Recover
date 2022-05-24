package main

//Go makes it possible to recover from a panic, by using the recover built-in function.
// A recover can stop a panic from aborting the program and let it continue with execution instead.

//The panic () function in Go Language is similar to exceptions raised at runtime when an error is encountered.
// panic() is either raised by the program itself when an unexpected error occurs or the programmer
//throws the exception on purpose for handling particular errors.

//The panic() function is inbuilt in Go Language and when it is raised,
// the code prints a panic message, and the function crashes. If there are other goroutines ongoing,
//they function normally and return. Once all goroutines have returned, the program crashes.
//The panic() function, if used deliberately by a programmer, can be given an input of the error message
//that the programmer wishes to print on the screen when a panic occurs.

//An example of where this can be useful: a server wouldn’t want to crash if one of the client
//connections exhibits a critical error. Instead, the server would want to close that connection and
//continue serving other clients. In fact, this is what Go’s net/http does by default for HTTP servers.

import "fmt"

func mayPanic() {
	panic("a problem")
}

func main() {

	//recover must be called within a deferred function. When the enclosing function panics,
	//the defer will activate and a recover call within it will catch the panic.

	defer func() {
		if r := recover(); r != nil {

			//The return value of recover is the error raised in the call to panic.

			fmt.Println("Recovered. Error:\n", r)
		}
	}()

	mayPanic()

	fmt.Println("After mayPanic()")
}

//This code will not run, because mayPanic panics. The execution of main stops
//at the point of the panic and resumes in the deferred closure.

//output :

//Recovered. Error:
//a problem
