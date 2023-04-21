package main

import "fmt"

func slices() {
	slice := [][]int{
		{9, 8, 7},
		{6, 5, 4},
		{3, 2, 1},
	}

	var reverse [][]int

	for i := len(slice) - 1; i >= 0; i-- {
		item := slice[i]

		var reverseItems []int

		for j := len(item) - 1; j >= 0; j-- {
			reverseItems = append(reverseItems, item[j])
		}

		reverse = append(reverse, reverseItems)
	}

	fmt.Println(reverse)

	s := make([]string, 2, 2)
	s[0] = "a"
	s[1] = "b"
	fmt.Println(s)
}

func maps() {
	m := make(map[string]int)
	m["hola"] = 1
	m["mundo"] = 2

	for k, v := range m {
		fmt.Println(k, v)
	}

	m2 := map[int]int{
		1: 1,
		2: 2,
	}

	for k, v := range m2 {
		fmt.Println(k, v)
	}

	val, hasKey := m2[3]
	fmt.Println(val, hasKey)
}

func Index[T comparable](slice []T, item T) (int, bool) {
	for i, v := range slice {
		if v == item {
			return i, true
		}
	}

	return -1, false
}

func main() {
	fmt.Println(Index([]int{1, 2, 3, 4}, 9))
	fmt.Println(Index([]string{"a", "b"}, "a"))
}

//https://go.dev/tour/methods/3
//https://go.dev/tour/methods/4
//https://go.dev/tour/methods/8

//https://go.dev/tour/generics/1
//https://go.dev/tour/generics/2

//https://go.dev/tour/concurrency/1
