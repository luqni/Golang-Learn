/*
Program Name : First Go Program
Version : 1.0
Description : Go Program basic program structure.
Author : Muhammad Luqni Baehaqi
Date Create : 10-08-2019
*/
package main
import "fmt"

/*
syntax untuk mendeklarasikan variable dalam golang

-> var <nama> <type>
	atau
-> var <nama> <type> = <expression>
*/

func main(){
	var hello_str string
	hello_str="Hello, World!"
	var nama string = "luqni"
	fmt.Println(hello_str)
	fmt.Println(nama)
}