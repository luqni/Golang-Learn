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
	i,j := 42, 2701

	p:=&i // ngepointer variable i
	fmt.Println(*p) // membaca isi variable i melalui pointer yang sudah dibuat

	*p := 21 //set variable i melaui pointer
	fmt.Println(i) // lihat nilai baru dari variable i
	
	p = &j //ngepointer variable j
	*p = *p / 37 // membagi nilai variable j melalui pointer yang sudah di buat
	fmt.Println(j) // lihat nilai baru dari variable j

}