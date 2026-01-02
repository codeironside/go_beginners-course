package main
import "fmt"

func main(){
//  emails := make(map[string]string)

 //assign kv

//  emails["bob"]="bob@gmail.com"
//  emails["sharon"]="sharon@gmail.com"

//Decalre map and kv
emails:= map[string]string{"bob":"bob@gmail.com", "sharon":"sharon@email.com"}

 fmt.Println(emails)
 fmt.Println(emails["bob"])

 //Delete from map
 delete(emails, "bob")
 fmt.Println(emails)
}