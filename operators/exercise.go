// arithmetic operators-this example multiplies 10 by 5 and prints result//

package main
import ("fmt")

func main () {
	fmt.print (10*5)
}

assignment operators-this example adds a value to a variable

package main
import ("fmt")

func main() {
	var x = 10
	x +=5
	fmt.Println(x)
}

//comparison operators-the return value of a comparison is either true(1) or false(0) //

package main
import ("fmt") 

func main(){
	var x = 5
	var y = 3
	fmt.Println(x>y) // result is 1 (true) because 5 is greater than 3
}
