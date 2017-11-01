package main

import (
	"fmt"
)

func main() {
	mySlice := make([]int, 0, 5)
	fmt.Println("-----------")
	fmt.Println(mySlice)
	fmt.Println(len(mySlice))
	fmt.Println(cap(mySlice))
	fmt.Println("-----------")

	for i := 0; i < 80; i++ {
		mySlice = append(mySlice, i)
		fmt.Println("Len:", len(mySlice), "Capacity:", cap(mySlice), "Value:", mySlice[i])
	}

	a := []int{1, 2, 3, 4}
	b := []int{5, 6, 7, 8}

	a = append(a, b...)
	fmt.Println(a)

	// multi-dimensional slice
	records := make([][]string, 0)

	student1 := make([]string, 4)
	student1[0] = "Foster"
	student1[1] = "Nathan"
	student1[2] = "100.00"
	student1[3] = "74.00"

	records = append(records, student1)

	student2 := make([]string, 4)
	student2[0] = "Gomez"
	student2[1] = "Liza"
	student2[2] = "92.00"
	student2[3] = "96.00"

	records = append(records, student2)

	fmt.Println(records)

	transactions := make([][]int, 0, 3)

	for i := 0; i < 3; i++ {
		transaction := make([]int, 0)
		for j := 0; j < 4; j++ {
			transaction = append(transaction, i+j)
		}
		transactions = append(transactions, transaction)
	}

	fmt.Println(transactions)

	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			fmt.Println(transactions[i][j])
		}
	}
}
