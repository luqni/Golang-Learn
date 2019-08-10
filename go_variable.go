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