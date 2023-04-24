/* although the + operators is often used to add together two values, it can also be used 
to add a varaible and a value, or variable and another varaiable*/
 
package main
import ("fmt")
func main() {

var (
     sum1 = 100+80 
	 sum2 = sum1 + 270
	 sum3 = sum2 + sum2
)

     fmt.Println(sum3)
}