package main
import ("fmt")

func familyName(fname string, age int) {
	fmt.Println("Hello", age, "year old", fname, "Refsnes")
}
func main() {
	familyName("Maggie", 4)
	familyName("Meg", 10)
	familyName("Fafa", 30)
}
