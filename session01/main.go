package main

import "fmt"

type CustomType struct{}

func main() {
	/*	a := "t1"
		b := 5
		switch {
		case a == "t":
			fmt.Println("a == t")
		case b == 3:
			fmt.Println("b == 3")
		case b == 5, a == "test string":
			fmt.Println("a = test string ; or b = 5")
		default:
			fmt.Println("default case")}*/

	var d interface{}
	//interface类型相当于java的Object类型

	//e := byte(1)
	e := CustomType{}
	d = &(e)
	switch t := d.(type) {
	case byte:
		fmt.Println("byte is byte", t)
	case *int:
		fmt.Println("*int is int", t)
	case *byte:
		fmt.Println("*byte is byte", t)
	case string:
		fmt.Println("string is string", t)
	case *string:
		fmt.Println("*string is string", t)
	case CustomType:
		fmt.Println("CustomType is", t)
	case *CustomType:
		fmt.Println("*CustomType is", t)
	default:
		fmt.Println("default case", t)
	}
}
