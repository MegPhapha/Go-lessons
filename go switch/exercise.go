//single case switch

package main
import ("fmt")

func main() {
	var day = 2
	switch day {
	case(1):
		fmt.Print("Saturday")
	case(2):
		fmt.Print("Sunday") // result is Sunday

	}
}

// switch statement 2

package main
import ("fmt")

func main() {
	var day = 4
	switch day {
	case(1):
		fmt.Print("Saturday")
	case(2):
		fmt.Print("Sunday")
	default:
		fmt.Print("Weekday")  // result is Weekday
	}
}



// multi case switch

packagemain
import ("fmt")

func main() {
	day := 6

	switch day {
	case 1,3,5:
		fmt.Println("odd weekday")
	case 2,4:
		fmt.Println("Even weekday")
	case 6,7:
		fmt.Println("Weekend")  // result is Weekend
	default:
		fmt.Println("Invalid day of day number")

	}
}