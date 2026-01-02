package main
import "fmt"

func main(){
 i:=1
 for i<=10{
	fmt.Println(i)
	i++
 }

 for i:=1; i<=10;i++{
	fmt.Printf("Number %d\n", i)
 }

 //Fizz buzz
 for i:=1; i<=100; i++{
	if i %15 == 0{
		fmt.Println("Fizzbuzz")
	} else if i %3 == 0{
		fmt.Println("Fizz")
	} else if i % 5 == 0{
		fmt.Println("buzz")
	} else {
		fmt.Println(i)
	}
 }
}