package models

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"unique"`
	Password string
	Tasks    []Task `gorm:"foreignKey:UserID"` // <- muito importante!
}
