package models

type User struct {
	ID    uint   `json:"id" gorm:"primary_key;auto_increment;not_null"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   uint8  `json:"age"`
}
