package main
import "fmt"

const xx = iota;

const( 	a = iota
	b = iota
	c = iota
	d = iota + 100
	e = iota * 2
	f = iota / 2
	g = iota + d
	h = iota - 100.5000
	
	// iota = 500  // assigning iota to some values will change all values to 500 and remain static
	i = iota
	j = iota
)

func main() {
	println("Constants")
	fmt.Printf("Value: %v and Type: %T\n", xx, xx)


	println("Constants - abc")
	fmt.Printf("Value: %v and Type: %T\n", a, a)
	fmt.Printf("Value: %v and Type: %T\n", b, b)
	fmt.Printf("Value: %v and Type: %T\n", c, c)
	
	
	println("Artmatics---------")
	fmt.Printf("Value: %v and Type: %T\n", d, d)
	fmt.Printf("Value: %v and Type: %T\n", e, e)
	fmt.Printf("Value: %v and Type: %T\n", f, f)
	fmt.Printf("Value: %v and Type: %T\n", g, g)
	fmt.Printf("Value: %v and Type: %T\n", h, h)
	fmt.Printf("Value: %v and Type: %T\n", i, i)
	fmt.Printf("Value: %v and Type: %T\n", j, j)


} // main ends

// Constants
// Value: 500 and Type: int
// Constants - abc
// Value: 500 and Type: int
// Value: 500 and Type: int
// Value: 500 and Type: int
// Artmatics---------
// Value: 600 and Type: int
// Value: 1000 and Type: int
// Value: 250 and Type: int
// Value: 1100 and Type: int
// Value: 399.5 and Type: float64
// Value: 500 and Type: int
// Value: 500 and Type: int
