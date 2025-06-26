package service

import (
	"sms/app/model"
	"sms/app/repository"
	"sort"
)

type StudentQuery struct {
	Name      string
	Email     string
	MinGpa    *float64
	MaxGpa    *float64
	SortBy    string // e.g., "name", "gpa"
	SortOrder string // "asc" or "desc"
	Page      int
	PageSize  int
}

type IStudentService interface {
	GetStudents(query StudentQuery) ([]model.Student, int, error) // students, total count, error
	GetStudentByID(id int) (model.Student, error)
	UpdateStudent(student model.Student) (model.Student, error)
	DeleteStudent(id int) error
	AddStudent(req model.CreateStudentRequest) (model.Student, error)
}

type StudentService struct {
	repository repository.IStudentRepository
}

func NewStudentService(repository repository.IStudentRepository) IStudentService {
	return &StudentService{repository}
}

func (s *StudentService) GetStudents(query StudentQuery) ([]model.Student, int, error) {
	students, err := s.repository.GetStudents()
	if err != nil {
		return nil, 0, err
	}

	// Filtering
	filtered := make([]model.Student, 0)
	for _, student := range students {
		if query.Name != "" && student.Name != query.Name {
			continue
		}
		if query.Email != "" && student.Email != query.Email {
			continue
		}
		if query.MinGpa != nil && student.Gpa < *query.MinGpa {
			continue
		}
		if query.MaxGpa != nil && student.Gpa > *query.MaxGpa {
			continue
		}
		filtered = append(filtered, student)
	}

	total := len(filtered)

	// Sorting
	switch query.SortBy {
	case "name":
		if query.SortOrder == "desc" {
			sort.Slice(filtered, func(i, j int) bool { return filtered[i].Name > filtered[j].Name })
		} else {
			sort.Slice(filtered, func(i, j int) bool { return filtered[i].Name < filtered[j].Name })
		}
	case "gpa":
		if query.SortOrder == "desc" {
			sort.Slice(filtered, func(i, j int) bool { return filtered[i].Gpa > filtered[j].Gpa })
		} else {
			sort.Slice(filtered, func(i, j int) bool { return filtered[i].Gpa < filtered[j].Gpa })
		}
	}

	// Pagination
	start := query.Page * query.PageSize
	end := start + query.PageSize
	if query.PageSize == 0 {
		return filtered, total, nil
	}
	if start > total {
		return []model.Student{}, total, nil
	}
	if end > total {
		end = total
	}
	return filtered[start:end], total, nil
}

func (s *StudentService) GetStudentByID(id int) (model.Student, error) {
	return s.repository.GetStudentByID(id)
}

func (s *StudentService) UpdateStudent(student model.Student) (model.Student, error) {
	return s.repository.UpdateStudent(student)
}

func (s *StudentService) DeleteStudent(id int) error {
	return s.repository.DeleteStudent(id)
}

func (s *StudentService) AddStudent(req model.CreateStudentRequest) (model.Student, error) {
	return s.repository.AddStudent(req)
}
