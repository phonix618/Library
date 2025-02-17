package main

import (
	"fmt"
	"sort"
	"strconv"
)

var books [][]string

func main() {
	for {
		fmt.Println("Welcome to the library :)")
		fmt.Println("1. Add Your Book ")
		fmt.Println("2. Display Books ")
		fmt.Println("3. Sort Books By Year Of Publication ")
		fmt.Println("4. Exit ")
		fmt.Println("Enter Your Choice: ")
		var choice int
		fmt.Scan(&choice)
		switch choice {
		case 1:
			addBook()
		case 2:
			displayBooks()
		case 3:
			sortBooksByYear()
		case 4:
			fmt.Println("Exit ")
			return
		default:
			fmt.Println("Invalid Number :( ")
		}
	}

}
func addBook() {
	var title, author, year string
	fmt.Print("Enter Title:")
	fmt.Scanln(&title)
	fmt.Print("Enter Author:")
	fmt.Scanln(&author)
	fmt.Print("Enter Year:")
	fmt.Scanln(&year)
	books = append(books, []string{title, author, year})
	fmt.Println("Books Added!")
}
func displayBooks() {
	if len(books) == 0 {
		fmt.Println("No Books Found!")
		return
	}
	for _, book := range books {
		fmt.Printf("Title: %s, Author: %s, Year Of Publication: %s\n", book[0], book[1], book[2])
	}
}
func sortBooksByYear() {
	sort.Slice(books, func(i, j int) bool {
		year1, _ :=
			strconv.Atoi(books[i][2])
		year2, _ :=
			strconv.Atoi(books[j][2])
		return year1 < year2
	})
	fmt.Println("Sorted Books By Year Of Publication :)")
}
