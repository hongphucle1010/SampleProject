package student

import "sample/app/model"

type IStudentRepository interface {
	GetStudents() ([]model.Student, error)
	AddStudent(req model.CreateStudentRequest) (model.Student, error)
	UpdateStudent(student model.Student) (model.Student, error)
	DeleteStudent(id int) error
	GetStudentByID(id int) (model.Student, error)
}
