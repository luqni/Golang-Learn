/*
Program Name : First Go Program
Version : 1.0
Description : Go Program basic program structure.
Author : Muhammad Luqni Baehaqi
Date Create : 10-08-2019
*/
package main
import "fmt"

func getSequence() func() int {
	i:=0
	return func() int {
		i+=1
		return i
	}
}

func main(){
/* nextNumber merupakan sebuah fungsi dengan i adalah 0 */
nextNumber := getSequence()

/* aktifkan nextNumber untuk mengisi variabel i dengan 1 dan mengembalikan hal yang sama */
fmt.Println(nextNumber())
fmt.Println(nextNumber())
fmt.Println(nextNumber())

fmt.Println("################")
/* buat sequence baru dan lihat hasilnya, i akan kembali bernilai 0 */
nextNumber1 := getSequence()
fmt.Println(nextNumber1())
fmt.Println(nextNumber1())
fmt.Println(nextNumber1())
}
