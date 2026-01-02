package main
import "fmt"

func main(){
 //Arrays
//   var fruitArr[2]string

//   //Assign values
//   fruitArr[0]="Apple"
//   fruitArr[1]="Orange"

//Decalare and assign same time

// fruitArr := [2]string{"apple", "orange"}
//   fmt.Println(fruitArr)
//   fmt.Println(fruitArr[1])
fruitSlice := []string{"apple", "orange", "Grape", "cherry"}

fmt.Println(fruitSlice[1:3])

}