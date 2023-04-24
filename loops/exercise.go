// this example will print the numbers from 0 to 4 //

package main
import ("fmt")

func main() {
	for i:=0; i < 5; i++ {
	  fmt.Println(i)	
	}
}


// this example counts to 100 by tens //

package main
import ("fmt")

func main() {
	for i:=0; i <=100; i+10 {
		fmt.Println(i)
	}
}

// this example skips the value of 3 //

package main
import ("fmt")

func main() {
	for i:=0; i < 5; i++ {
		if i == 3 {
			continue
		}
		fmt.Println(i)
	}
}
