/*
Program Name : First Go Program
Version : 1.0
Description : Go Program basic program structure.
Author : Muhammad Luqni Baehaqi
Date Create : 10-08-2019
*/
package main
import "fmt"

type person struct {
	name string
	age int
}

func main(){
	fmt.Println(person{"Bambang", 20})
	fmt.Println(person{name: "Aji", age:30})
	fmt.Println(person{name: "Joko", age:40})
	fmt.Println(person{name: "Urip", age:10}) 

	s:= person{name:"Setia", age: 25}
	fmt.Println(s.name)

	sp:=&s
	fmt.Println(sp.age) //pointer
}