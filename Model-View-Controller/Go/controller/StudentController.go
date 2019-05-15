package controller

import (
	"Development-Technology/Model-View-Controller/Go/model"
	mdl "Development-Technology/Model-View-Controller/Go/model"
	vew "Development-Technology/Model-View-Controller/Go/view"
)

//-------------Controller------------

type StudentController struct {
	model model.Student
	view  vew.StudentView
}

func (s *StudentController) StudentController(model mdl.Student, view vew.StudentView) {
	(*s).model = model
	(*s).view = view
}

func (s *StudentController) SetStudentName(name string) {
	(*s).model.SetName(name)
}
func (s *StudentController) SetStudentRollNo(rollNo string) {
	(*s).model.SetRollNo(rollNo)
}
func (s *StudentController) GetStudentName() string {
	return (*s).model.GetName()
}
func (s *StudentController) GetStudentRollNo() string {
	return (*s).model.GetRollNo()
}
func (s *StudentController) UpdateView() {
	(*s).view.PrintStudentDetails((*s).model.GetName(), (*s).model.GetRollNo())
}
