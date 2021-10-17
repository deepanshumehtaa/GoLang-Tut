/* 
https://youtu.be/-Zg276gE1ZY
https://perennialsky.medium.com/understand-handle-handler-and-handlefunc-in-go-e2c3c9ecef03
Interface: 	method signature
		In Go language, the interface is a custom type that is used to specify a set of one or more method signatures and the interface is abstract, 
		so you are not allowed to create an instance of the interface. 
		But you are allowed to create a variable of an interface type and this variable can be assigned with a c.
		
		it is necessary to implement all the methods declared in the interface for implementing an interface.



   Rules:
	1.

*/

package main


type Rect struct {
	length, breadth float64
}
var rect_obj Rect;

func (r Rect) area() (float64) {
	return r.length*r.breadth
}

func (r Rect) perimeter() (float64) {
	return 2*(r.length + r.breadth)
}

// INTERFACE
type Shape interface{
	area() float64
}

func main() {
	rect_obj = Rect{
		length: 10,
		breadth: 10,
	}
	println(rect_obj)
	println(rect_obj.area())
	println(rect_obj.perimeter())

} // main ends



// CASE 2. 


























