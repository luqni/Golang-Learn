/*
Program Name : First Go Program
Version : 1.0
Description : Go Program basic program structure.
Author : Muhammad Luqni Baehaqi
Date Create : 10-08-2019
*/
package main //(package declaration) merupakan deklarasi paket.
/* mendefinisikan nama paket di mana program ini akan dimasukkan. 
Program "Go" berjalan dalam paket, sehingga diperlukan bagian ini untuk mendefinisikan suatu paket.
Setiap paket memiliki nama dan jalur yang terkait dengannya.
*/
import "fmt" //Preprocessor Statements
/*
bagian ini digunakan untuk mendifinisikan semua preprocessor statement yang akan di import
import "statement" digunakan untuk memberitahu kompiler
 */
 func main() { //entry point
	fmt.Println("Hello, Word!")
 }