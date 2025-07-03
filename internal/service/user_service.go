package service

import (
	"errors"
	"github.com/Chabuduo04/User_management_system/internal/model"
	"github.com/Chabuduo04/User_management_system/internal/storage"
)

type UserService struct {
	storage storage.UserStorage
}

func NewUserService(storage storage.UserStorage) *UserService {
	return &UserService{storage: storage}
}

func (s *UserService) Register(req *model.CreateUserRequest) (*model.User, error) {
	// 参数校验
	if len(req.Password) < 8 {
		return nil, errors.New("password too short")
	}

	user := &model.User{
		ID:       generateID(), // 假设有ID生成函数
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
		Roles:    req.Roles,
	}

	if err := s.storage.Create(user); err != nil {
		return nil, err
	}
	return user, nil
}

var id int = 0
func generateID() int{
	id++
	return id
}