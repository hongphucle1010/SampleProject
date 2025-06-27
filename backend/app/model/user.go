package model

type User struct {
	ID       int    `json:"id,omitempty" example:"100000" bson:"_id"`
	Username string `json:"username,omitempty" example:"phuc" bson:"username"`
	Password string `json:"password,omitempty" example:"123456" bson:"password"`
	Role     string `json:"role,omitempty" example:"admin" bson:"role"`
}

type RegisterRequest struct {
	Username string `json:"username" binding:"required" example:"phuc"`
	Password string `json:"password" binding:"required" example:"123456"`
	Role     string `json:"role" binding:"required" example:"admin"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required" example:"phuc"`
	Password string `json:"password" binding:"required" example:"123456"`
}

type Teacher struct {
	User
	FirstName string `json:"first_name,omitempty" example:"Phuc" bson:"first_name"`
	LastName  string `json:"last_name,omitempty" example:"Le" bson:"last_name"`
	Email     string `json:"email,omitempty" example:"phucle@example.com" bson:"email"`
	Phone     string `json:"phone,omitempty" example:"0909090909" bson:"phone"`
	Address   string `json:"address,omitempty" example:"KTX Bach Khoa, 497 Hoa Hao Street" bson:"address"`
}
