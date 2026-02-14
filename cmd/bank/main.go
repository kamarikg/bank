package main

import "fmt"

func main() {
	categories := []string{"auto", "food", "beauty", "mobile", "fun"}
	top3 := append(categories[0:3], "job")
	fmt.Println(top3)
	fmt.Println(len(top3))
	fmt.Println(cap(top3))
}