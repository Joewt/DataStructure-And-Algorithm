package main


import "fmt"

func test(arr *[3]int){
	(*arr)[0] = 7
	fmt.Println(arr) //pri
}

func main() {
	x := [3]int{1, 2, 3}

	test(&x)

	fmt.Println(x) //prints [7 2 3]
}

