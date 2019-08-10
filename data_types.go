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
	var varByte byte ='a'
	var varRune rune = 'b'

	fmt.Printf("%c = %d and %c = %U\n", varByte, varByte, varRune, varRune)
}