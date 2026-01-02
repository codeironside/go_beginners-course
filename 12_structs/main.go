package main
import ("fmt" 
"strconv")


//Define person structs
type Person struct {
	firstName string
	lastName string
	city string
	gender string
	age int
}

//Greating method (value receiver)
func(p Person)greet()string{
	return"hello, name name is "+p.firstName + " "+ p.lastName + " and I AM " + strconv.Itoa(p.age)
}
//has borthday methods (pointer receiver)

func(p *Person) hasBirthday(){
	p.age++
}

// getsMarried (pointer receiver)

func (p *Person)getsMarried(spouseLastName string){
	if p.gender =="m"{
		return
	}else{
		p.lastName = spouseLastName 
	}
}

func main(){
person1:=Person{firstName:"Samatha", lastName:"Smith", city:"Boston", gender:"f", age:25}

// alternative
person2:=Person{"bob", "johnson", "accra","m", 35}
// fmt.Println(person1)

// fmt.Println(person1.firstName)
// person1.age++
person1.hasBirthday()
person1.getsMarried("Williams")
person2.getsMarried("Thompson")
fmt.Println(person1.greet())
fmt.Println(person2.greet())

}