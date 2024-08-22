package repository

type Todo struct {
    ID          int64  `gorm:"primaryKey"`
    Title       string `gorm:"size:255"`
    Description string `gorm:"size:255"`
    Completed   bool
}

type Repository interface {
    Create(todo *Todo) error
    Get(id int64) (*Todo, error)
    List() ([]*Todo, error)
    Update(todo *Todo) error
    Delete(id int64) error
}
