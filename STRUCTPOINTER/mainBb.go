package main

import "fmt"

type Student struct {
ID int
Name string
GPA float32

}
func (student *Student) graduate() {
	student.Name = student.Name + " S.AP"
}
func main() {
	student := Student{1, "Muhazi Ramadhan", 3.37}
	fmt.Println(student.Name)

	student.graduate()
	fmt.Println(student.Name)
}

