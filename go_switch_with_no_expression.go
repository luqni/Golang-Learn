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
	var nilai = 55

	fmt.Println("Learn - Go Switch with No Expression")

	switch {
	case nilai > 60:
		fmt.Println("Lulus dengan grade 'A' ")
	case nilai < 60 && nilai >=50:
		fmt.Println("Lulus dengan grade 'B' ")
	case nilai < 50 && nilai >= 35:
		fmt.Println("Lulus dengan grade 'C' ")
	default:
		fmt.Println("Kamu Gagal")
	}
}