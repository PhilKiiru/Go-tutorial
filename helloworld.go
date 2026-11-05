package main 

import "fmt"

func main () {
	var i int = 56
	var j string = "Boy!"

	//i 
	fmt.Printf("The value of i is: %v and its type is: %T\n", i,i)
	fmt.Printf("The percentage of i is: %v%% and it's Go-format is: %#v\n", i,i)

	//j
	fmt.Printf("The value of j is: %v and its type is: %T\n", j,j)
	fmt.Printf("The percentage of j is: %v%% and its Go-format is: %#v", j,j)
}

// Go compact code
// package main; import "fmt" func main () {fmt.Println("Hello World!");}

// := not used outside a function
// var can be used outside a function
// const can be used outside a function
// \n creates new lines
// Printf = format with %v, %T, %d, %s, etc., and you add \n yourself
// %v is used to print the VALUE of arguments
// %T is used to print the TYPE of arguments
// %% is used to show the percentage 
// %#v is used to print the Go-syntax format of the value 
