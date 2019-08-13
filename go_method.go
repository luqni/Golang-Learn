/*
Program Name : First Go Program
Version : 1.0
Description : Go Program basic program structure.
Author : Muhammad Luqni Baehaqi
Date Create : 10-08-2019
*/
package main
import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() int{
	return r.width * r.height
}

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main(){
	r := rect{width:10, height:5}

	fmt.Println("Area", r.area())
	fmt.Println("Perim", r.perim())

	rp := &r
	fmt.Println("area", rp.area())
	fmt.Println("perim", rp.perim()o)
}