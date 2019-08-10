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
	var x = 2
	fmt.Println("Learn - Go Switch with fallthrough")

	switch x {
	case  1:
		fmt.Println("1")
		fallthrough
	case 2:
		fmt.Println("2")
		fallthrough
	case 3:
		fmt.Println("3")
	}
}