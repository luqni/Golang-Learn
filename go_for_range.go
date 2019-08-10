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
	langs := []string{"Go","C","C++","Java"}
	fmt.Println("Learn - Go For Each Loop")

	for i,s := range langs{
		fmt.Println(i,s)
	}
}