package repository

import "gorm.io/gorm"

type GORMRepository struct {
	db *gorm.DB
}

func NewGORMRepository(db *gorm.DB) Repository {
	return &GORMRepository{db: db}
}

func (r *GORMRepository) Create(todo *Todo) error {
	return r.db.Create(todo).Error
}

func (r *GORMRepository) Get(id int64) (*Todo, error) {
	var todo Todo
	if err := r.db.First(&todo, id).Error; err != nil {
		return nil, err
	}
	return &todo, nil
}

func (r *GORMRepository) List() ([]*Todo, error) {
	var todos []*Todo
	if err := r.db.Find(&todos).Error; err != nil {
		return nil, err
	}
	return todos, nil
}

func (r *GORMRepository) Update(todo *Todo) error {
	return r.db.Save(todo).Error
}

func (r *GORMRepository) Delete(id int64) error {
	return r.db.Delete(&Todo{}, id).Error
}
