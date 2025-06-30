package models

type Task struct {
	ID     uint   `gorm:"primaryKey"`
	Title  string `gorm:"not null"`
	Done   bool
	UserID uint
	User   User `gorm:"foreignKey:UserID"`
}

// package models

// type Task struct {
// 	ID     uint   `gorm:"primaryKey"`
// 	Title  string
// 	Done   bool
// 	UserID uint
// }
