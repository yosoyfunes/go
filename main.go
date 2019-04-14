// Go version: go1.10.1 linux/amd64
package main

import "fmt"

// Strings support formatting in your declaration
const HELLO_USER string = "Hello %s %s, welcome to fantastic Go's world\n"
// A const not used not produce a error
// Go inference the type
const ONE = 1

func main(){
  // Declaration
  var name string
  // Assignation
  name = "Name"
  // Declaration and assignation, Go inference the type
  lastname := "Lastname"
  // Declaration and assignation, Go inference the type
  var helloWorld = "Hello world!"
  // Multiple assignation, equivalent to:
  //  var a = 1
  //  var b = "Two"
  //  var c = true
  var (
    a = 1
    b = "Two"
    c = true
  )

  fmt.Print("What's your name? ")
  fmt.Scanf("%s", &name)
  fmt.Print("What's your lastname? ")
  fmt.Scanf("%s", &lastname)
  fmt.Printf(HELLO_USER, name, lastname)
  fmt.Println(helloWorld, a, b, c)
}