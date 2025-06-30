package models

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Tasks    []Task `gorm:"foreignKey:UserID"`
}

// package models

// type User struct {
// 	ID       uint   `gorm:"primaryKey"`
// 	Username string `gorm:"unique"`
// 	Password string
// 	Tasks    []Task
// }
