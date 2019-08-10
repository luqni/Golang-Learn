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

	var count = 0
	fmt.Println("Learn - Go Break Statement")

	for count <= 10 {
		count++
		if (count == 5){
			break
		}
		fmt.Printf("Inside loop %d\n", count)
	}
	fmt.Println("Out of loop")
}