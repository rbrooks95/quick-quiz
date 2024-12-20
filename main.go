package main

import (
	"encoding/csv"
	"fmt"
	"os"
)


func main()  {
	fd, error := os.Open("questions.csv")
	fmt.Println("hello")

	if error != nil {
    fmt.Println(error)
}

fmt.Println("Successfully opened the CSV file")
defer fd.Close()
fileRead := csv.NewReader(fd)
questions,error := fileRead.ReadAll()

if error != nil {
	fmt.Println(error)
}
fmt.Println(questions)
for i := 0; i < len(questions); i++{
	fmt.Println(questions[i])
}
}
