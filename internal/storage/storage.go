package storage

import "github.com/Chabuduo04/User_management_system/internal/model"

// 存储接口抽象
type UserStorage interface {
	Create(user *model.User) error
	Update(id int, user *model.User) error
	Delete(id int) error
	GetByID(id int) (*model.User, error)
	List() ([]*model.User, error)
}