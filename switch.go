/**
 @author:way
 @date:2021/11/24
 @note
**/

package main

import "fmt"

func main() {
	var a = "hell"
	switch a {
	case "hello":
		println("hello")
	case "hi":
		println("hi")
	default:
		fmt.Println(0)
	}

	var a1 ="mum"
	switch a1 {
	case "mum","daddy":
		fmt.Println("flm")
	}

	var s = "hello"
	switch  {
	case s=="hello":
		fmt.Println("hello")
		fallthrough
	case s=="world":
		fmt.Println("world")
	}
}
