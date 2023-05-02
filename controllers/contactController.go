package controllers

import (
	"contact_log/models"
	"contact_log/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Test(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "test",
	})
}

func GetContacts(c *gin.Context) {
	contacts, err := services.GetContacts()
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error while getting contacts",
		})
	} else {
		c.JSON(200, contacts)
	}
}

func GetContactByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid ID",
		})
		return
	}

	contact, err := services.GetContactsByID(uint(id))
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error while getting contact",
		})
	} else {
		c.JSON(200, contact)
	}
}

func CreateContact(c *gin.Context) {
	var requestContact models.RequestContact

	if err := c.BindJSON(&requestContact); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	contact, err := services.CreateContact(requestContact)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create contact"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": contact})
}

func UpdateContact(c *gin.Context) {
	var updatedContact models.ResponseContact
	if err := c.BindJSON(&updatedContact); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	contact, err := services.UpdateContact(updatedContact)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update contact"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": contact})
}
