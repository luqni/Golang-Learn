package main
import "fmt"

func main(){
	var a int = 30
	var b int = 5

	fmt.Printf("Belajar - Go Assignment Operators")
	a+=b
	fmt.Printf("a+=b :%d\n", a)
	a-=b
	fmt.Printf("a-=b :%d\n", a)
	a*=b
	fmt.Printf("a*=b :%d\n", a)
	a/=b
	fmt.Printf("a/=b :%d\n", a)
	a%=b
	fmt.Printf("a%%=b :%d\n", a)
}