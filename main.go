package main

import (
	"fmt"

	"github.com/MydroX/go-visual-sort/internal/sorting"
)

func main() {
	//Numbers need to be positive non null integer
	numbers := []int{2, 45, 5, 1, 6, 11, 3, 4, 7, 6}
	res := sorting.Sort(numbers)

	fmt.Println(res)
}
