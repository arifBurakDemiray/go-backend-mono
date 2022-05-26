package models

// User struct
type Role struct {
	ID   string `gorm:"column:id;primary_key;not null" json:"id"`
	Name string `gorm:"column:name" json:"name"`
}

// TableName func
func (Role) TableName() string {
	return "roles"
}
