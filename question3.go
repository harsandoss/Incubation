package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}
func reverseString(str string) string {
	byte_str := []rune(str)
	for i, j := 0, len(byte_str)-1; i < j; i, j = i+1, j-1 {
		byte_str[i], byte_str[j] = byte_str[j], byte_str[i]
	}
	return string(byte_str)
}
func main() {
	var n int
	var String string
	fmt.Print("Enter the number of string to be added =")
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		scanner := bufio.NewScanner(os.Stdin)

		fmt.Printf("Enter the string%d :", i)
		fmt.Scan(&String)

		scanner.Scan()

		a := SortString(String)
		b := reverseString(a)

		fmt.Print(b)
	}
}
