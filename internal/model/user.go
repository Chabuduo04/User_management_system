package model

//用户实体定义
type User struct {
	ID	int	`json:"id"`
	Name	string	`json:"name"`
	Email	string	`json:"email"`
	Password	string	`json:"-"`
	Roles	[]string	`json:"roles"`
}

//用户创建请求DTO
type CreateUserRequest struct {
	Name	string	`json:"name" validate:"required"`
	Email	string	`json:"email" validate:"email"`
	Password	string	`json:"password" validate:"min=8"`
	Roles	[]string	`json:"roles"`
}
