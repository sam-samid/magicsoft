package main

import "fmt"

func main() {
	array := []int{1, 4, 5, 6, 8, 2}
	fmt.Println("==============================GARPH============================")
	graph(array)
	fmt.Println("==========================INSERTIONSORT========================")
	insertsort(array)
	fmt.Println("============================REVERSESORT========================")
	a := insertsort(array)
	reversesort(a)
}

func graph(array []int) {
	firstarray := array[0]

	for index := 0; index <= 5; index++ {
		if array[index] > firstarray {
			firstarray = array[index]
		}
	}

	for row := firstarray; row >= 1; row-- {
		for col := 0; col <= 5; col++ {
			if array[col] >= row {
				fmt.Printf(" | ")
			} else {
				fmt.Printf("   ")
			}
		}
		fmt.Printf("\n")
	}

	for index := 0; index <= 5; index++ {
		fmt.Printf(" %d ", array[index])
	}
	fmt.Printf("\n")
}

func insertsort(array []int) (sortarray []int) {
	for i := 1; i < 6; i++ {
		j := i
		for j > 0 {
			if array[j-1] > array[j] {
				array[j-1], array[j] = array[j], array[j-1]
				graph(array)
			}
			j = j - 1
		}
	}
	sortarray = array
	return sortarray
}

func reversesort(array []int) {
	for i := 1; i < 6; i++ {
		j := i
		for j > 0 {
			if array[j-1] < array[j] {
				array[j-1], array[j] = array[j], array[j-1]
				graph(array)
			}
			j = j - 1
		}
	}
}
