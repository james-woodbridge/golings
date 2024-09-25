// slices3
// Make me compile!

package main

import "fmt"

func main() {
	var names []string
	fmt.Println(len(names))
	fmt.Println(cap(names))
	names = append(names, "Tom")
	fmt.Println(len(names))
	fmt.Println(cap(names))
	fmt.Println(names)
}
