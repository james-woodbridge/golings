// slices1
// Make me compile!

package main

import "fmt"

func main() {
	a := make([]int, 5, 10) // play with length and capacity
	a[0] = 0
	a[1] = 1
	a[2] = 2
	a[3] = 3
	a[4] = 4

	fmt.Println("length of 'a':", len(a))
	fmt.Println("capacity of 'a':", cap(a))
}
