package main

import "fmt"

type MyStruct struct {
   id int
	 text string
}

func main() {
	/* this is a comment I hope */

	var i int
	i = 12

	fmt.Println("Hello, World!")

	fmt.Println(i)

	foo(&i, "This is j")

	fmt.Println(i)

	var array []int
	array = make([]int, 3, 5)

	for j := 0; j < len(array); j++  {
		fmt.Println(array[j])
	}

	var s MyStruct
	s.id = 10
	s.text = "The id is ten here"

	fmt.Println(s.text)

}

//func function_name( [parameter list] ) [return_types]
func foo(i *int, j string) bool  {
	fmt.Println("We are in the subroutine y'all")

	*i++

	fmt.Println(i)
	fmt.Println(j)

	return true
}
