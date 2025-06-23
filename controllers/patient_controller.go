package controllers

import (
	"hospitalapp/config"
	"hospitalapp/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPatients(c *gin.Context) {
	var patients []models.Patient
	config.DB.Find(&patients)
	c.JSON(http.StatusOK, patients)
}

func GetpatientById(c *gin.Context) {
	id := c.Param("id")
	var patient models.Patient
	if err := config.DB.Find(&patient, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error" : "patient not found",
		})
		return
	} 
	c.JSON(http.StatusOK, patient)
}

func CreatePatient(c *gin.Context) {
	var patient models.Patient
	err := c.BindJSON(&patient)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : "invalid data",
		})
		return
	}
	config.DB.Create(&patient)
	c.JSON(http.StatusCreated, gin.H{
		"id" : patient.ID,
		"message" : "patient added",
	})
}

func UpdatePatient(c *gin.Context) {
	id := c.Param("id")
	var patient models.Patient

	// Find patient by ID
	if err := config.DB.First(&patient, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Patient not found",
		})
		return
	}

	// Create a temporary struct to avoid overwriting fields like ID, CreatedAt
	var input models.Patient
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid input for update",
		})
		return
	}

	// Update fields (you can do this more safely with a DTO struct)
	patient.Name = input.Name
	patient.Age = input.Age
	patient.Disease = input.Disease
	patient.Address = input.Address
	patient.AssignedTo = input.AssignedTo

	config.DB.Save(&patient)

	c.JSON(http.StatusOK, patient)
}


func DeletePatient(c *gin.Context) {
	id := c.Param("id")
	var patient models.Patient

	if err := config.DB.Find(&patient, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H {
			"error" : "patient not found",
		})
		return
	} 
	config.DB.Delete(&models.Patient{}, id)
	c.JSON(http.StatusOK, gin.H{
		"error" : "patient deleted",
	})
}