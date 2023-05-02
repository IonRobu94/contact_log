package models

type RequestContact struct {
	FirstName       string `json:"first_name" gorm:"not null;column:first_name"`
	LastName        string `json:"last_name" gorm:"not null;column:last_name"`
	DateOfBirth     string `json:"date_of_birth" gorm:"column:date_of_birth"`
	Nationality     string `json:"nationality" gorm:"column:nationality"`
	Address         string `json:"address" gorm:"column:address"`
	TelephoneNumber string `json:"telephone_number" gorm:"not null;column:telephone_number"`
	Email           string `json:"email" gorm:"not null;unique; column:email"`
	CreatedAt       string `json:"created_at" gorm:"column:created_at"`
	UpdatedAt       string `json:"updated_at" gorm:"column:updated_at"`
}
type ResponseContact struct {
	RequestContact
	ID uint `gorm:"primary_key" json:"id"`
}

func (ResponseContact) TableName() string {
	return "contacts"
}
