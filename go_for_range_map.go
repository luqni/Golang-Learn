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

	fmt.Println("Learn - Go For Each with Map")
	fruits := map[string]string{"A":"Apple","B":"Banna","C":"Cherry"}

	for key, value := range fruits {
		fmt.Printf("%s -> %s\n", key, value)
	}
}