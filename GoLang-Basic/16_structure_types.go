// the classes of GO

package main
import "fmt"

type Rect struct {
	length int
	breadth float64
	isTrue bool
}

// the way how u define struct methods
func (r *Rect) getArea() float64 {
    return float64(r.length) * r.breadth
}

type Square struct {
	length int
	breadth float64
	isTrue bool
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

}  // main ends
