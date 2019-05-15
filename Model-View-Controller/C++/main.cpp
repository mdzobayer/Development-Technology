#include "Controller/StudentController.cpp"


int main() {

    Student model("Zobayer", "CSE 059 07156");
    StudentView view;
    StudentController controller(model, view);

    controller.updateView();
    controller.setStudentName("Zobayer Mahmud");
    controller.updateView();

    return (0);
}