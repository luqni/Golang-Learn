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
	var dayOfWeek = 5

	switch dayOfWeek {
	case 1,2,3,4,5 :
		fmt.Println("Its a Weekday")

	case 6,7:
		fmt.Println("Its a Weekend")

	default:
		fmt.Println("Invalid day")
	}
}