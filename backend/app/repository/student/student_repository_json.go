package student

import (
	"fmt"
	"sample/app/model"
	"sample/app/utils"

	"github.com/spf13/viper"
)

type JsonStudentRepository struct{}

func NewJsonStudentRepository() IStudentRepository {
	return &JsonStudentRepository{}
}

// GetStudents implements IStudentRepository.
func (r *JsonStudentRepository) GetStudents() ([]model.Student, error) {
	students, err := utils.ReadJsonFile[model.Student](viper.GetString("database.path"))
	if err != nil {
		return []model.Student{}, err
	}

	return students, nil
}

// AddStudent implements IStudentRepository.
func (r *JsonStudentRepository) AddStudent(req model.CreateStudentRequest) (model.Student, error) {
	students, err := utils.ReadJsonFile[model.Student](viper.GetString("database.path"))
	if err != nil {
		return model.Student{}, err
	}
	student := model.Student{
		ID:    utils.GenerateNextID(students, func(s model.Student) int { return s.ID }),
		Name:  req.Name,
		Email: req.Email,
		Dob:   req.Dob,
		Gpa:   req.Gpa,
	}
	students = append(students, student)
	return student, utils.WriteJsonFile(viper.GetString("database.path"), students)
}

// UpdateStudent implements IStudentRepository.
func (r *JsonStudentRepository) UpdateStudent(student model.Student) (model.Student, error) {
	students, err := utils.ReadJsonFile[model.Student](viper.GetString("database.path"))
	if err != nil {
		return model.Student{}, err
	}
	updated := false
	for i, s := range students {
		if s.ID == student.ID {
			students[i] = student
			updated = true
			break
		}
	}
	if !updated {
		return model.Student{}, fmt.Errorf("student with ID %d not found", student.ID)
	}
	return student, utils.WriteJsonFile(viper.GetString("database.path"), students)
}

// DeleteStudent implements IStudentRepository.
func (r *JsonStudentRepository) DeleteStudent(id int) error {
	students, err := utils.ReadJsonFile[model.Student](viper.GetString("database.path"))
	if err != nil {
		return err
	}
	newStudents := make([]model.Student, 0, len(students))
	deleted := false
	for _, s := range students {
		if s.ID == id {
			deleted = true
			continue
		}
		newStudents = append(newStudents, s)
	}
	if !deleted {
		return fmt.Errorf("student with ID %d not found", id)
	}
	return utils.WriteJsonFile(viper.GetString("database.path"), newStudents)
}

// GetStudentByID implements IStudentRepository.
func (r *JsonStudentRepository) GetStudentByID(id int) (model.Student, error) {
	students, err := utils.ReadJsonFile[model.Student](viper.GetString("database.path"))
	if err != nil {
		return model.Student{}, err
	}
	for _, s := range students {
		if s.ID == id {
			return s, nil
		}
	}
	return model.Student{}, fmt.Errorf("student with ID %d not found", id)
}
