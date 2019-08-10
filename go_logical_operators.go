/*
Program Name : First Go Program
Version : 1.0
Description : Go Program basic program structure.
Author : Muhammad Luqni Baehaqi
Date Create : 10-08-2019
*/
package main
import "fmt"

func main(){
	var a int = 20
	var b int = 10
	var c int = 25

	var flag bool = false
	var result bool
	fmt.Println("Learn - Go Logical Operators")
	result = (a > b) && (a > c)
	fmt.Printf("(a>b) && (a>c) : %t\n", result)
	result = (a > b) || (a > c)
	fmt.Printf("(a.b) || (a>c) : %t\n", result)
	result = !flag
	fmt.Printf("!flag : %t\n", result)	
}