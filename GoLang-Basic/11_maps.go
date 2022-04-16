/* Maps https://youtu.be/ORqGqhpw33w?list=PLa6iDxjj_9qXaVIemjejvomRgKu3r7t3J
   Maps (GO) <-->  un ordered dictionary (python)

   Rules:
	1. Declaration (with make keyword):
		m := make(map[<key Datatype>]<value Datatype>, initial_size)
		key DataTypes can't be slices, maps or functions.
		but could be int, str, float, array, strct, pointers and interfaces.
		initial_size is just reserve the memory initially, but it doesn't restrict the growth of map.

	2. Accessing value with the key that not exist:
		- don't raise error
		- will return default datatype value, 0 if int.
		- := will use when u declare new variable from pool, and
		   =(equals to), when u overriding the declared variable

	3. Map type is like:
		map[string]int

	4. Functions in the Maps:
		i)   len() -> gives the length of the dict.
		ii)  the change in the map is -> pass by reference ie.
			map2 := map1
			any change in map1 will reflect in map2, becoz pointing the same memory location.


*/

package main

import (
	"fmt"
)

func main() {

	// Declaration
	// m1 := map[string]int{}
	// m2 := make(map[string]int, initial_size)

	var m1 = map[int]int{}        // with var keyword
	m2 := make(map[string]int, 1) // with make, with initial size as 1
	m3 := map[string]int{"a": 10, "b": 20}

	println(m3)     // 0xc000034678
	fmt.Println(m3) // map[a:10 b:20]

	// appending
	m1[10] = 110
	m3["d"] = 0

	// deleting
	delete(m3, "a")

	fmt.Println(m1) // map[10:100]
	fmt.Println(m2) // map[]
	fmt.Println(m3) // map[b:20 d:0]

	// Rule-2
	fmt.Println(m3["z"]) // 0

	// Rule3
	value, is_exist := m3["z"]
	fmt.Println(value, is_exist) // 0 false

	value, is_exist = m3["b"]
	fmt.Println(value, is_exist) // 20 true

	// Rule 4.2
	m3_refrence := m3
	fmt.Println(m3_refrence) // map[b:20 d:0]

	delete(m3_refrence, "b")
	m3_refrence["z"] = 1000
	m3_refrence["zz"] = 1001

	fmt.Println(m3)          // map[d:0 z:1000 zz:1001]
	fmt.Println(m3_refrence) // map[d:0 z:1000 zz:1001]

	// printing keys and valuse
	for k, v := range m3 {
	  fmt.Println(k, v)
	}

	// printing both key and value
	for k := range m3 {
	  fmt.Println(k)
	}

	// printing only value
	for _, v := range m3 {
	  fmt.Println(v)
	}

	// ************** array as key **************
	m4 := map[[3]int]int{{1, 2}: 10, {10, 1}: 20}
	fmt.Println(m4) // map[[1 2 0]:10 [10 1 0]:20]

} // main ends
