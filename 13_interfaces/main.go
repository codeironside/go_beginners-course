package main
import ("fmt"
"math")

//Define interface

type shape interface{
	area()float64
}

type Circle struct{
	x,y, radius float64
}
type rectangle struct{
	l, b float64
}
func (c Circle) area()float64{
	return math.Pi* c.radius * c.radius
}
func (r rectangle) area()float64{
	return r.l * r.b
}
func getArea(s shape) float64 {
	return s.area()
}

func main(){
 cicrcle := Circle{x:0,y:0,radius:5}
 rec:= rectangle{l:10, b:5}
 fmt.Printf("Circle Area:%f\n", getArea(cicrcle))
 fmt.Printf("Rectangle Area:%f\n", getArea(rec))
}