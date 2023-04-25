// in this example, we store the return value in a variable called total //
package main
import ("fmt")

func myFunction(x int, y int) (result int) {
	result = x + y
	return
}

func main() {
	total := myFunction(1, 2)
	fmt.Println(total)
}
