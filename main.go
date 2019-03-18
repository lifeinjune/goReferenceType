package main

import "fmt"

func main() {
	h := []string{"hi", "there", "how", "are", "you"}
	updateString(h)
	fmt.Println(h)

}
func updateString(s []string) {
	s[0] = "bye"
}
