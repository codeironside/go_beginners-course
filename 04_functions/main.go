package main
import "fmt"
 func greeting(name string )string{
	return "hello "+name
 }

 func getSum(num1 uint, num2 uint)uint{
	return num1 + num2
 }
func main(){
 fmt.Println(getSum(6,5))
}