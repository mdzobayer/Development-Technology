package model

//--------------Model----------
type Student struct {
	name   string
	rollNo string
}

func (s *Student) Student(name, rollNo string) {
	(*s).name = name
	(*s).rollNo = rollNo
}

func (s *Student) GetRollNo() string {
	return (*s).rollNo
}
func (s *Student) GetName() string {
	return (*s).name
}
func (s *Student) SetName(name string) {
	(*s).name = name
}
func (s *Student) SetRollNo(rollNo string) {
	(*s).rollNo = rollNo
}
