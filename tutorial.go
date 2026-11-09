package main

import "fmt"

func main() {
	var i string = "Hello"
	var j int = 344
	var x bool = true
	var y float32 = 89.3

	fmt.Printf("The value of i is: %s \n", i)
	fmt.Printf("The value of j to base 2 is: %b \n", j)
	fmt.Printf("The value of x is: %t \n", x)
	fmt.Printf("The value of exponent is y is: %e \n", y)
}

// 1. Integer formatting verbs
//b base 2
//d base 10
//+d base 10 and shows sign
//o base 8
//O base 8, leading Oo
//x base 16 lower
//X base 16 upper
//#x base 16, leading with 0x
//+4d padding with space (width 4, right justified)
//-4d padding with space )with 4, left justified)
//04d padding with zeros (width 4)

// test work
/* func main() {
	var i int = 400

	fmt.Printf("Base 2 for i is: %b \n", i)
	fmt.Printf("Base 10 for i is: %d \n", i)
	fmt.Printf("Base 16 for i is: %x for lower", i)
	fmt.Printf("Base 16 for i is: %X for upper", i)
	fmt.Printf("Base 16 for i leading with 0x is: %#x", i)
	fmt.Printf("width 4 padding right for i is: %4d", i)
	fmt.Printf("Width 4 padding left for i is: %-4d", i)
	fmt.Printf("Paidding zero for i is: %04d, i")
	fmt.Printf("Plus sign for i is: %+d", i)
	fmt.Printf("Base 8 for i is: %o", i)
	fmt.Printf("Base 8 for is starting with Oo is: %O", i)

} */

// 2. String Formating Verbs
// %s prints the string value
// %q prints the string with double quotes
// %8s prints the plain string (width 8, right justified)
// %-8s print the plain string (width 8, left justified)
// %x print as hex dump
// % x prints as hex dump with spaces
// A hex dump is a human-readable display of raw computer data (from a file, memory/RAM, or storage device)
// shown in hexadecimal (base-16) format.

// test work
/* func main() {
	var i string = "Hello"

	fmt.Printf("The value of i in plain string is: %s \n", i)
	fmt.Printf("The value of i in plain string with double quotes is: %q \n", i)
	fmt.Printf("The value of i in 8 width, right justified is: %8s \n", i)
	fmt.Printf("The value of i in 8 width, left justified is: %-8s \n", i)
	fmt.Printf("The value of i in hex dump is: %x \n", i)
	fmt.Printf("The value of i in hex dump with spaces is: % x \n", i)
} */

// 3. Boolean Formating Verbs
// %t prints the value of the boolean operator (true or false). Works like the %v

//test work
/* func main () {
	var i bool = true
	var j bool = false

	fmt.Printf("The value of i is: %t \n", i)
	fmt.Printf("The value of j is: %t \n", j)
} */

// 4. Float Formating Verbs
// %e prints e as exponent
// %f prints the decimal point, no exponent
// %.2f defualt width, precision 2
// %6.2f width 6, precision 2
// %g exponent as needed , only necessary digits

//test work
/* func main() {
	var i float32 = 3.141

	fmt.Printf("The exponent for i is: %e \n", i)
	fmt.Printf("The decimal point for i with no exponent is: %f \n", i)
	fmt.Printf("The default width, precision 2 for i is: %.2f \n", i)
	fmt.Printf("The width 6, precision 2 for i is: %6.2f \n", i)
	fmt.Printf("The exponent only, with necessary digits for i is: %g", i)
} */
