package main
import "fmt"

func main(){
 a:=5
 b:= &a

 fmt.Println(a, b)
 fmt.Printf("%T\n",b)

 // use * to read valur from memory address

 fmt.Println(*b)

 //change value

 *b = 10
 fmt.Println(a)
}