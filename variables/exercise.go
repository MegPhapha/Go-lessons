Create a variable named myNum and assign the value 50 to it.

package main 
import ("fmt")

func main () {
var myNum = 50
 fmt.Println(myNum)
}

 Different Types of Variables In Same Line

 package main
import ("fmt")

func main() {
  var a, b = 6, "Hello"
  c, d := 7, "World!"

  fmt.Println(a)
  fmt.Println(b)
  fmt.Println(c)
  fmt.Println(d)
}

Declaration In a Block

package main
import ("fmt")

func main() {
   var (
     a int
     b int = 1
     c string = "hello"
   )

  fmt.Println(a)
  fmt.Println(b)
  fmt.Println(c)
}

Naming Rules

Camel Case

myVariableName = "John"

Pascal Case
MyVariableName = "John"

Snake Case

my_variable_name = "John"
