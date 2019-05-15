#include <bits/stdc++.h>

using namespace std;

// Student Model
class Student {
    private:
        string rollNo;
        string name;
    public:
        Student() {
            rollNo = "";
            name = "";
        }
        Student(string name, string rollNo) {
            this->rollNo = rollNo;
            this->name = name;
        }
        string getRollNo() {
            return rollNo;
        }
        string getName() {
            return name;
        }
        void setName(string name) {
            this->name = name;
        }
        void setRollNo(string rollNo) {
            this->rollNo = rollNo;
        }
};
