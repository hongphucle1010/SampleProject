package controller

import (
	"sample/app/model"
	"sample/app/service"
	"sample/pkg/response"
	"strconv"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type StudentController struct {
	Ctx            iris.Context
	StudentService service.IStudentService
}

// @Summary Get all students
// @Description Get all students with optional filtering, sorting, and pagination
// @Tags students
// @Produce  json
// @Param name query string false "Filter by name"
// @Param email query string false "Filter by email"
// @Param min_gpa query number false "Minimum GPA"
// @Param max_gpa query number false "Maximum GPA"
// @Param sort_by query string false "Sort by field (name, gpa)"
// @Param sort_order query string false "Sort order (asc, desc)"
// @Param page query int false "Page number (starts from 0)"
// @Param page_size query int false "Page size"
// @Success 200 {object} response.SuccessResponse[model.PaginatedResponse[model.Student]]
// @Failure 500 {object} response.ErrorResponse
// @Router /students/ [get]
func (c *StudentController) Get() mvc.Result {
	// Parse query parameters
	name := c.Ctx.URLParam("name")
	email := c.Ctx.URLParam("email")
	minGpaStr := c.Ctx.URLParam("min_gpa")
	maxGpaStr := c.Ctx.URLParam("max_gpa")
	sortBy := c.Ctx.URLParamDefault("sort_by", "")
	sortOrder := c.Ctx.URLParamDefault("sort_order", "asc")
	page := c.Ctx.URLParamIntDefault("page", 0)
	pageSize := c.Ctx.URLParamIntDefault("page_size", 10)

	var minGpa, maxGpa *float64
	if minGpaStr != "" {
		if v, err := strconv.ParseFloat(minGpaStr, 64); err == nil {
			minGpa = &v
		}
	}
	if maxGpaStr != "" {
		if v, err := strconv.ParseFloat(maxGpaStr, 64); err == nil {
			maxGpa = &v
		}
	}

	query := service.StudentQuery{
		Name:      name,
		Email:     email,
		MinGpa:    minGpa,
		MaxGpa:    maxGpa,
		SortBy:    sortBy,
		SortOrder: sortOrder,
		Page:      page,
		PageSize:  pageSize,
	}

	students, total, err := c.StudentService.GetStudents(query)
	if err != nil {
		return response.ErrorResponse{
			Code:    iris.StatusInternalServerError,
			Message: "Failed to fetch students",
			Details: err.Error(),
		}
	}

	return response.SuccessResponse[model.PaginatedResponse[model.Student]]{
		Code:    iris.StatusOK,
		Message: "Successfully fetched students",
		Data: model.PaginatedResponse[model.Student]{
			Data:     students,
			Total:    total,
			Page:     page,
			PageSize: pageSize,
		},
	}
}

// @Summary Get student by ID
// @Description Get a student by their ID
// @Tags students
// @Produce  json
// @Param id path int true "Student ID"
// @Success 200 {object} response.SuccessResponse[model.Student]
// @Failure 404 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /students/{id} [get]
func (c *StudentController) GetBy(id int) mvc.Result {
	student, err := c.StudentService.GetStudentByID(id)
	if err != nil {
		return response.ErrorResponse{
			Code:    iris.StatusNotFound,
			Message: "Student not found",
			Details: err.Error(),
		}
	}
	return response.SuccessResponse[model.Student]{
		Code:    iris.StatusOK,
		Message: "Successfully fetched student",
		Data:    student,
	}
}

// @Summary Create a new student
// @Description Add a new student
// @Tags students
// @Accept  json
// @Produce  json
//
//	@Param student body model.CreateStudentRequest true "Student to add" \
//	 @example {"dob": "2004-10-10T00:00:00Z", "email": "hongphucle1010@gmail.com", "gpa": 4.0, "name": "Le Hong Phuc"}
//
// @Success 201 {object} response.SuccessResponse[model.Student]
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /students/ [post]
func (c *StudentController) Post() mvc.Result {
	var req model.CreateStudentRequest
	if err := c.Ctx.ReadJSON(&req); err != nil {
		return response.ErrorResponse{
			Code:    iris.StatusBadRequest,
			Message: "Invalid request body",
			Details: err.Error(),
		}
	}
	created, err := c.StudentService.AddStudent(req)
	if err != nil {
		return response.ErrorResponse{
			Code:    iris.StatusInternalServerError,
			Message: "Failed to create student",
			Details: err.Error(),
		}
	}
	return response.SuccessResponse[model.Student]{
		Code:    iris.StatusCreated,
		Message: "Student created successfully",
		Data:    created,
	}
}

// @Summary Update a student
// @Description Update a student by ID
// @Tags students
// @Accept  json
// @Produce  json
// @Param id path int true "Student ID"
// @Param student body model.Student true "Student data to update"
// @Success 200 {object} response.SuccessResponse[model.Student]
// @Failure 400 {object} response.ErrorResponse
// @Failure 404 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /students/{id} [put]
func (c *StudentController) PutBy(id int) mvc.Result {
	var student model.Student
	if err := c.Ctx.ReadJSON(&student); err != nil {
		return response.ErrorResponse{
			Code:    iris.StatusBadRequest,
			Message: "Invalid request body",
			Details: err.Error(),
		}
	}
	student.ID = id
	updated, err := c.StudentService.UpdateStudent(student)
	if err != nil {
		return response.ErrorResponse{
			Code:    iris.StatusNotFound,
			Message: "Student not found or update failed",
			Details: err.Error(),
		}
	}
	return response.SuccessResponse[model.Student]{
		Code:    iris.StatusOK,
		Message: "Student updated successfully",
		Data:    updated,
	}
}

// @Summary Delete a student
// @Description Delete a student by ID
// @Tags students
// @Produce  json
// @Param id path int true "Student ID"
// @Success 204 {object} nil
// @Failure 404 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /students/{id} [delete]
func (c *StudentController) DeleteBy(id int) mvc.Result {
	err := c.StudentService.DeleteStudent(id)
	if err != nil {
		return response.ErrorResponse{
			Code:    iris.StatusNotFound,
			Message: "Student not found or delete failed",
			Details: err.Error(),
		}
	}
	c.Ctx.StatusCode(iris.StatusNoContent)
	return nil
}
