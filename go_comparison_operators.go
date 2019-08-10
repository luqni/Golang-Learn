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

	if (a > b){
		fmt.Println("Learn - Go Relational Operators")
		fmt.Println("a is greater than b.")
	}else{
		fmt.Println("Learn - Go Relational Operator")
		fmt.Println("b is greater than a.")
	}
}