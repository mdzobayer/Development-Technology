// Reference 1: https://stackoverflow.com/questions/36967419/can-i-create-a-struct-inside-the-package-to-use-outside-in-main-program
// Reference 2: https://aggarwalarpit.wordpress.com/2017/07/08/creating-your-own-package-in-go/

package main

import (
	"Development-Technology/Model-View-Controller/Go/controller"
	"Development-Technology/Model-View-Controller/Go/model"
	"Development-Technology/Model-View-Controller/Go/view"
)

func main() {
	var model model.Student

	//type Student = mdl.Student

	model.Student("Zobayer", "CSE 059 07156")

	var view view.StudentView

	var controller controller.StudentController
	controller.StudentController(model, view)

	controller.UpdateView()
	controller.SetStudentName("Zobayer Mahmud")
	controller.UpdateView()

}
