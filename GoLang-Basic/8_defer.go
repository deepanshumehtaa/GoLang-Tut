/* defer

   How comiler works?
   >> 	Compilers are utility programs that take your code and transform it into executable machine code files. 
	When you run a compiler on your code
		Preprocessor reads the source code (the C++ file you just wrote) and searches for any preprocessor #directives. 
		Preprocessor directives cause the preprocessor to change your code in some way (by usually adding some library or another C++ file).
	then, compiler works through the preprocessed code line by line translating(tokenizing, parsing) into the appropriate machine language instruction.
	This will also uncover any syntax errors that are present in your source code and will throw an error to the command line.
	if no errors are present, the compiler creates an object file with the machine language binary necessary to run on your machine. 
	While the object file that the compiler just created is likely enough to do something on your computer, 
	it still isn’t a working executable of your C++ program. There is a final important step to reach an executable program.
	
	Every cpp file get genrated into object file.

   Rules:
	1. defer pushes the function call at the starting of the stack, thus executed at the end.
	2. by end means end of all the function present in the main.
	3. NOTE: defer call in the loop 3 times and the every result of defer stores in the starting of stack thus o/p: 321
	4. In Go language, multiple defer statements are allowed in the same program and they are executed in LIFO(Last-In, First-Out) or STACK.
	5. Defer statements are generally used to ensure that the files are closed when your work is finished with them, or to close the channel,
	   or to catch the panics in the program.
	6. Unlike variables, we can left fuction decleared but not used.
	
	Application:
	1. 
	
*/

package main

func multiplication(a1 int, a2 int) int {
  
    res:= a1 * a2
    return res
    
}
  
func show() { println("Hello!, GeeksforGeeks") }


func main() {

	for i :=1; i<=3; i++ {
	println(" first in the loop")
	}
	
	println("\n with defer ")
	
	for i :=1; i<=3; i++ {
	defer print(i)
	println(" in the loop")
	}
	
	
	for i :=1; i<=3; i++ {
	print(i)
	println(" Last loop")
	}
	
	
	ans := multiplication(23, 45);
	println("THE MULTIPLICATION:", ans)
	
	


} // main ends


