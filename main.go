package main

import "fmt"

func main() {
	h := []string{"hi", "there", "how", "are", "you"} //create slice of string
	updateString(h)                                   //call function to update string by passing argument of slice of string
	fmt.Println(h)                                    //print slice
	/*because slice is made out of two different memory which are
	  	1. one with information about slice like length of the slice, capacity of the slice, pointer to actual array of the slice store the values in
	  	2. array of the slice that contain actual value
	  so when slice is passed as argument, it copy and passes the information about the slice not the actual value of the slice
	  therefor, copy of the slice still going to point into the same actual value of the array
	*/
}
func updateString(s []string) {
	s[0] = "bye"
}
