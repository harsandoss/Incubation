package main

import "fmt"

func main() {
	var i, j, n int
	fmt.Println("Enter the nmber: ")
	fmt.Scan(&n)
	for i = n; i > 1; i-- {
		for j = n; j >= 1; j-- {
			if j > i {
				fmt.Printf("%d", j)
			} else {
				fmt.Printf("%d", i)
			}
		}
		for j = 2; j <= n; j++ {
			if j > i {
				fmt.Printf("%d", j)
			} else {
				fmt.Printf("%d", i)
			}
		}
		fmt.Println(" ")
	}
	for i = 1; i <= n; i++ {
		for j = n; j >= 1; j-- {
			if j > i {
				fmt.Printf("%d", j)
			} else {
				fmt.Printf("%d", i)
			}
		}
		for j = 2; j <= n; j++ {
			if j > i {
				fmt.Printf("%d", j)
			} else {
				fmt.Printf("%d", i)
			}
		}
		fmt.Println(" ")
	}
	return
}
