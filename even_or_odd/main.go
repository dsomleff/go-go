package main

import "fmt"

type numbers []int

func (n numbers) evenOrOdd() {
	for _, v := range n {
		if v%2 == 0 {
			fmt.Println(v, "is even")
		} else {
			fmt.Println(v, "is odd")
		}
	}
}

func main() {
	nums := make(numbers, 11)
	for i := range nums {
		nums[i] = i
	}

	nums.evenOrOdd()
}
