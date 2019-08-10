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
	var age int
	election:
	fmt.Print("Enter your age : ")
	fmt.Scanln(&age)

	if (age <= 17){
		fmt.Println("You are not eligible to vote.")
		goto election
	}else{
		fmt.Println("You are eligible to vote")
	}
}