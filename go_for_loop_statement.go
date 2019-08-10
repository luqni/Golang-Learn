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
	sum := 0
	fmt.Println("Learn - Go for loop")
	
	for i := 1; i<5; i++ {
		sum += i
	}
	fmt.Println(sum)
}