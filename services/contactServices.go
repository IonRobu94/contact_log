package services

import (
	"contact_log/models"
	"contact_log/repositories"
)

func GetContacts() ([]models.ResponseContact, error) {
	return repositories.GetContacts()
}
func GetContactsByID(id uint) (*models.ResponseContact, error) {
	return repositories.GetContactsByID(id)
}
func CreateContact(requestContact models.RequestContact) (*models.ResponseContact, error) {
	return repositories.CreateContact(requestContact)
}
func UpdateContact(updatedContact models.ResponseContact) (*models.ResponseContact, error) {
	return repositories.UpdateContact(updatedContact)
}
