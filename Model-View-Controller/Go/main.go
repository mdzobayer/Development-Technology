package main

import "fmt"

//--------------Model----------
type Student struct {
	name   string
	rollNo string
}

func (s *Student) Student(name, rollNo string) {
	(*s).name = name
	(*s).rollNo = rollNo
}
func (s *Student) getRollNo() string {
	return (*s).rollNo
}
func (s *Student) getName() string {
	return (*s).name
}
func (s *Student) setName(name string) {
	(*s).name = name
}
func (s *Student) setRollNo(rollNo string) {
	(*s).rollNo = rollNo
}

//--------View-----------
type StudentView struct {
}

func (s *StudentView) printStudentDetails(studentName, studentRollNo string) {
	fmt.Println("Student: ")
	fmt.Println("Name: ", studentName)
	fmt.Println("Roll No: ", studentRollNo)
}

//-------------Controller------------

type StudentController struct {
	model Student
	view  StudentView
}

func (s *StudentController) StudentController(model Student, view StudentView) {
	(*s).model = model
	(*s).view = view
}

func (s *StudentController) setStudentName(name string) {
	(*s).model.setName(name)
}
func (s *StudentController) setStudentRollNo(rollNo string) {
	(*s).model.setRollNo(rollNo)
}
func (s *StudentController) getStudentName() string {
	return (*s).model.getName()
}
func (s *StudentController) getStudentRollNo() string {
	return (*s).model.getRollNo()
}
func (s *StudentController) updateView() {
	(*s).view.printStudentDetails((*s).model.getName(), (*s).model.getRollNo())
}

func main() {
	var model Student

	model.Student("Zobayer", "CSE 059 07156")
	var view StudentView

	var controller StudentController
	controller.StudentController(model, view)

	controller.updateView()
	controller.setStudentName("Zobayer Mahmud")
	controller.updateView()

}
