package main

import "fmt"

func main() {
	var studentScore int = 101

	fmt.Println(studentScore)
	if studentScore >= 80 && studentScore <= 100 {
		fmt.Println("A")
	} else if studentScore >= 65 && studentScore < 80 {
		fmt.Println("B+")
	} else if studentScore >= 50 && studentScore < 65 {
		fmt.Println("B")
	} else if studentScore >= 35 && studentScore < 50 {
		fmt.Println("C")
	} else if studentScore >= 0 && studentScore < 35 {
		fmt.Println("D")
	} else {
		fmt.Println("Invalid")
	}
}
