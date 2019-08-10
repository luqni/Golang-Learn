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
	var a,b,c int
	a = 50
	b = 10
	
	c = a & b // & bitwise AND integers
	fmt.Println(c)
	
	c = a | b // | bitwise OR integers
	fmt.Println(c)

	c = a ^ b // ^ bitwise aOR integers
	fmt.Println(c)

	c = a &^ b // &^ bit clear (AND OR) integers
	fmt.Println(c)

}