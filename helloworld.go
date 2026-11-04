package main 

import "fmt"

func main () {
	var name1 string = "Phil"
	var num1 float32 = 45

	fmt.Printf("name1 has the value: %v and type: %T \n", name1 , name1)
	fmt.Printf("num1 has the value: %v and type: %T \n", num1, num1)
}
// Go compact code
// package main; import "fmt" func main () {fmt.Println("Hello World!");}

// := not used outside a function
// var can be used outside a function
// const can be used outside a function
// \n creates new lines 
// %v is used to print the VALUE of arguments 
// %T is used to print the TYPE of arguments 