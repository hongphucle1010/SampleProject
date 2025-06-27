package model

type Student struct {
	ID    int    `json:"id,omitempty" example:"220000" bson:"_id"`
	Name  string `json:"name,omitempty" example:"Le Hong Phuc"`
	Email string `json:"email,omitempty" example:"hongphucle1010@gmail.com"`
	Dob   string `json:"dob,omitempty" example:"2004-10-10T00:00:00Z"`
	Gpa   float64 `json:"gpa,omitempty" example:"4.0"`
}

type CreateStudentRequest struct {
	Name  string  `json:"name" binding:"required" example:"Le Hong Phuc"`
	Email string  `json:"email" binding:"required" example:"hongphucle1010@gmail.com"`
	Dob   string  `json:"dob" binding:"required" example:"2004-10-10T00:00:00Z"`
	Gpa   float64 `json:"gpa" binding:"required" example:"4.0"`
}
