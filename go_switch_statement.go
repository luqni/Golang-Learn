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
	fmt.Println("learn - Go Switch statement.")

	switch dayOfWeek {
	case 1:
		fmt.Println("Hari Senin")
	case 2:
		fmt.Println("Hari Selasa")
	case 3:
		fmt.Println("Hari Rabu")
	case 4:
		fmt.Println("Hari Kamis")
	case 5:
		fmt.Println("Hari Jum'at")
	case 6:
		fmt.Println("Hari Sabtu")
	case 7:
		fmt.Println("Hari Ahad")
	default:
		fmt.Println("Invalid Weekday")
	}
}