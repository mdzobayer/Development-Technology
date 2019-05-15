package view

import "fmt"

//--------View-----------
type StudentView struct {
}

func (s *StudentView) PrintStudentDetails(studentName, studentRollNo string) {
	fmt.Println("Student: ")
	fmt.Println("Name: ", studentName)
	fmt.Println("Roll No: ", studentRollNo)
}
