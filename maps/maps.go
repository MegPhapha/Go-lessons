Creating Maps Using var and :=
Syntax

var a = map[KeyType]ValueType{key1:value1, key2:value2,...}
b := map[KeyType]ValueType{key1:value1, key2:value2,...}

example

package main
import ("fmt")

func main() {
	var a = map[string]string{"brang": "Ford", "mode1": "Mustang", "year": "1964"}
	b := map[string]int{"Oslo": 1, "Bergen": 2, "Trondheim" 3, "Stavanger": 4}

	fmt.Printf("a\t%v\n", a)
	fmt.Printf("b\t%v\n", b)
}