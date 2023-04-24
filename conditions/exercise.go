/* if staement
this example prints "Hello World" if x is greater than y

package main
import("fmt")

func main() {
	var x = 50
	var y =10
	if x > y {
		fmt.Print("Hello World")
	}
}


/* using the if else statement
print Good day if time:=20 is less than 18 else print Good evening*/

package main
import ("fmt")

func main() {
	time := 20
	if (time < 18) {
		fmt.Println("Good day.")
	} else {
		fmt.Println("Good evening")
	}
}

/* nested if 
this example shows how to use nested if statements*/

package main
imprt ("fmt")

func main(){
	fmt.Println("Num s more than 10.")
	if num > 15 {
		fmt.Println("Num is also more than 15.")
	}
} else {
	fmt.Println("Num is less than 10.")
}