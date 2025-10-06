package handlers

import (
	"net/http"

	"gym-membership/config"
	"gym-membership/models"

	"github.com/gin-gonic/gin"
)

// ========================== CREATE ==========================
func CreateTrainer(c *gin.Context) {
	var input models.Trainer
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat trainer"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Trainer berhasil dibuat",
		"data":    input,
	})
}

// ========================== GET ==========================
func GetTrainers(c *gin.Context) {
	var trainers []models.Trainer
	if err := config.DB.Find(&trainers).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data trainer"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Daftar trainer berhasil diambil",
		"data":    trainers,
	})
}

// ========================== UPDATE ==========================
func UpdateTrainer(c *gin.Context) {
	id := c.Param("id")
	var trainer models.Trainer

	if err := config.DB.First(&trainer, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Trainer tidak ditemukan"})
		return
	}

	var input models.Trainer
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	trainer.Name = input.Name
	trainer.Expertise = input.Expertise

	if err := config.DB.Save(&trainer).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memperbarui trainer"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Trainer berhasil diperbarui",
		"data":    trainer,
	})
}

// ========================== DELETE ==========================
func DeleteTrainer(c *gin.Context) {
	id := c.Param("id")
	var trainer models.Trainer

	if err := config.DB.First(&trainer, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Trainer tidak ditemukan"})
		return
	}

	if err := config.DB.Delete(&trainer).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghapus trainer"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Trainer berhasil dihapus",
	})
}
