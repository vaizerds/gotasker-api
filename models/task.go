package models

type Task struct {
    ID       uint   `gorm:"primaryKey"`
    Title    string
    Done     bool
    UserID   uint
}
