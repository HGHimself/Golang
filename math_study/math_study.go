package main

import "fmt"

/*
 n^a = n^(a/2) * n^(a/2) for any even a
 n^a = n^(a/2) * n^(a/2) * n for any odd a
 here I believe you do about half as many multiplication
 so O(n/2) ?
 O(log n) function calls
*/
func power_int(b int, a int) int  {
	fmt.Printf("base:%v, exp:%v\n", b, a)
	if a == 0  { return 1 }
	n := power_int(b, a >> 1)
	if a % 2 == 0  { return n * n }
	if a > 0 { return b * n * n }
	return n * n / b
}

func main()  {
	fmt.Println("Hello World!!")
	exp := 16
	base := 2
	res := power_int(base, exp)
	fmt.Printf("%v^%v = %v\n", base, exp, res)
}
