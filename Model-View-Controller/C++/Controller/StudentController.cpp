#include <bits/stdc++.h>

using namespace std;

#include "../Model/Student.cpp"
#include "../View/StudentView.cpp"


// Controller
class StudentController {
    private:
        Student model;
        StudentView view;
    
    public:
        StudentController(Student model, StudentView view) {
            this->model = model;
            this->view = view;
        }

        void setStudentName(string name) {
            model.setName(name);
        }

        string getStudentName() {
            return model.getName();
        }

        void setStudentRollNo(string rollNo) {
            model.setRollNo(rollNo);
        }

        string getStudentRollNo() {
            return model.getRollNo();
        }

        void updateView() {
            view.printStudentDetails(model.getName(), model.getRollNo());
        }
};