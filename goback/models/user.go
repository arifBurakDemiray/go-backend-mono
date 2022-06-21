package models

// User struct
type User struct {
	ID       string `gorm:"column:id;primary_key;not null" json:"id"`
	Password string `gorm:"column:password" json:"-"`
	MaxToDo  int    `gorm:"column:max_todo" json:"max_todo"`
	Email    string `gorm:"column:email" json:"email"`
}

// TableName func
func (User) TableName() string {
	return "users"
}
