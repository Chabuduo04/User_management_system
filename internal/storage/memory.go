package storage

import (
	"sync"
	"github.com/Chabuduo04/User_management_system/internal/model"
)

type MemoryStorage struct {
	users map[int]*model.User
	mu    sync.RWMutex // 保证并发安全
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		users: make(map[int]*model.User),
	}
}

func (m *MemoryStorage) Create(user *model.User) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	
	m.users[user.ID] = user
	return nil
}

func (m *MemoryStorage) Update(id int, user *model.User) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.users[id] = user
	return nil
}

func (m *MemoryStorage) Delete(id int) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	delete(m.users, id)
	return nil
}

func (m *MemoryStorage) GetByID(id int) (*model.User, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	return m.users[id], nil
}

func (m *MemoryStorage) List() ([]*model.User, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	res := make([]*model.User, len(m.users))
	for _, user := range m.users {
		res = append(res,user)
	}
	return res, nil
}