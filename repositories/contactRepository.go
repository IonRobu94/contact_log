package repositories

import (
	"contact_log/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

func getEnvConfig() string {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	dbname := os.Getenv("DB_NAME")
	dataSource := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname
	return dataSource
}

func GetContacts() ([]models.ResponseContact, error) {
	dataSource := getEnvConfig()

	db, err := gorm.Open(mysql.Open(dataSource), &gorm.Config{})
	if err != nil {
		log.Fatalf("Errore nel connettersi al database: %s", err)
	}

	var connections []models.ResponseContact
	if err := db.Find(&connections).Error; err != nil {
		log.Fatalf("Errore nell'eseguire la query: %s", err)
		return nil, err
	}

	return connections, nil
}

func GetContactsByID(id uint) (*models.ResponseContact, error) {
	dsn := getEnvConfig()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Errore nel connettersi al database: %s", err)
	}

	var contact models.ResponseContact
	if err := db.Where("id = ?", id).First(&contact).Error; err != nil {
		log.Fatalf("Errore nell'eseguire la query: %s", err)
		return nil, err
	}

	return &contact, nil
}
func CreateContact(requestContact models.RequestContact) (*models.ResponseContact, error) {
	dsn := getEnvConfig()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Errore nel connettersi al database: %s", err)
	}

	contact := models.ResponseContact{
		RequestContact: requestContact,
	}

	if err := db.Create(&contact).Error; err != nil {
		log.Fatalf("Errore nell'eseguire la query: %s", err)
		return nil, err
	}

	return &contact, nil
}
func UpdateContact(updatedContact models.ResponseContact) (*models.ResponseContact, error) {
	dsn := getEnvConfig()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Errore nel connettersi al database: %s", err)
	}

	if err := db.Model(&models.ResponseContact{}).Where("id = ?", updatedContact.ID).Omit("id", "created_at").Updates(&updatedContact).Error; err != nil {
		log.Fatalf("Errore nell'eseguire la query: %s", err)
		return nil, err
	}

	return &updatedContact, nil
}
