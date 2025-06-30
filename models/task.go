package models

type Task struct {
	ID     uint   `gorm:"primaryKey"`
	Title  string `gorm:"not null"`
	Done   bool
	UserID uint `gorm:"not null"`
	User   User `gorm:"not null"`
}

// package models

// type Task struct {
// 	ID     uint   `gorm:"primaryKey"`
// 	Title  string
// 	Done   bool
// 	UserID uint
// }
