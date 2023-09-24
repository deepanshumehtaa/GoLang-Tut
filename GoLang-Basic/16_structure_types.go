package main
import "fmt"

type Rect struct {
	length int
	breadth float64
	isTrue bool
}

// the way how u define struct methods
func (r *Rect) getArea() (float64, ) {
    return float64(r.length) * r.breadth
}

type Triangle struct {
	dimensions *Rect
	hypotnuse float64
}

func (tri *Triangle) getArea() (float64, ) {
    return float64(tri.dimensions.length) * tri.dimensions.breadth * tri.hypotnuse
}

func main() {
    
    var p0 Rect = Rect{11, 22, true}
    fmt.Println(p0)
    
    p1 := Rect{10, 20, false}
    fmt.Println(p1)
    
    // Recommended (Also, using this way gives u the power to use some parameter as optional)
    p2 := Rect{length: 21, breadth: 22, isTrue: true}
    fmt.Println("length and breadth are: ", p2.length, p2.breadth, p2.isTrue)
    
    p3 := Rect{length: 21, breadth: 22}
    fmt.Println("length and breadth with Area: ", p3.length, p3.breadth, p3.isTrue, p3.getArea())
    
    
    // level 2
    p11 := Rect{length: 11, breadth: 22}
    t1 := Triangle{dimensions: &p11, hypotnuse: 200}  // magic is here dude
    fmt.Println("Triangle Length: ", t1.dimensions.length)
    fmt.Println("Triangle Area: ", t1.getArea())

}
