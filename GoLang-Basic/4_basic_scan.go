package main

import "fmt"

func main() {
	var name string
	var alphabet_count int

	n, err := fmt.Sscanf("LOVE is having 4 alphabets.", "%s is having %d alphabets.", &name, &alphabet_count)
		
	if err != nil {
        panic(err)
	}
	
	fmt.Printf("%d: %s, %d\n", n, name, alphabet_count)
	
	
	
	
	value1, value2, value3 := "", "", 0

    	// Scan the 3 strings into 3 local variables.
    	// ... Must pass in the address of the locals.
    	_, err1 := fmt.Sscan("Bird frog 100", &value1, &value2, &value3);

    	// If the error is nil, we have successfully parsed the values.
    	if err1 == nil {
        	// Print the values.
        	fmt.Println("VALUE1:", value1);
        	fmt.Println("VALUE2:", value2);
        	fmt.Println("VALUE3:", value3);
    	}

	/*
	

	*/

} // main ends

