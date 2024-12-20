package main

import (
	"bufio"
	//"encoding/csv"
	"fmt"
	"os"
)


func main()  {
// file, err := os.Open("questions.csv")

// if err != nil{
// 	fmt.Println("there is a problem opening the file")
// }

// defer file.Close()

// reader := csv.NewReader(file)

// read, err := reader.ReadAll()

// if err != nil {
// 	fmt.Println("problem reading file")
// }

// for _, v := range read{
// 	fmt.Println(v[0],v[1])
// }
Input()
}


func Input()  {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter text....")
	text, _ := reader.ReadString('\n')

	fmt.Println("you entered", text)
}