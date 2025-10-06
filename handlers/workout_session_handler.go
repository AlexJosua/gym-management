package handlers

import (
	"net/http"

	"gym-membership/config"
	"gym-membership/models"

	"github.com/gin-gonic/gin"
)

// ========================== CREATE ==========================
func CreateWorkoutSession(c *gin.Context) {
	var input struct {
		UserID    uint   `json:"user_id" binding:"required"`
		TrainerID uint   `json:"trainer_id" binding:"required"`
		Date      string `json:"date" binding:"required"`
		Duration  int    `json:"duration" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	session := models.WorkoutSession{
		UserID:    input.UserID,
		TrainerID: input.TrainerID,
		Date:      input.Date,
		Duration:  input.Duration,
	}

	if err := config.DB.Create(&session).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat session"})
		return
	}

	config.DB.Preload("User").Preload("Trainer").First(&session, session.ID)

	c.JSON(http.StatusCreated, gin.H{
		"message": "Workout session berhasil dibuat",
		"data":    session,
	})
}

// ========================== GET ==========================
func GetWorkoutSessions(c *gin.Context) {
	var sessions []models.WorkoutSession
	if err := config.DB.Preload("User").Preload("Trainer").Find(&sessions).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil session"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Daftar workout sessions berhasil diambil",
		"data":    sessions,
	})
}

// ========================== UPDATE ==========================
func UpdateWorkoutSession(c *gin.Context) {
	id := c.Param("id")
	var session models.WorkoutSession

	if err := config.DB.First(&session, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Workout session tidak ditemukan"})
		return
	}

	var input struct {
		UserID    uint   `json:"user_id"`
		TrainerID uint   `json:"trainer_id"`
		Date      string `json:"date"`
		Duration  int    `json:"duration"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.UserID != 0 {
		session.UserID = input.UserID
	}
	if input.TrainerID != 0 {
		session.TrainerID = input.TrainerID
	}
	if input.Date != "" {
		session.Date = input.Date
	}
	if input.Duration != 0 {
		session.Duration = input.Duration
	}

	if err := config.DB.Save(&session).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memperbarui session"})
		return
	}

	config.DB.Preload("User").Preload("Trainer").First(&session, session.ID)

	c.JSON(http.StatusOK, gin.H{
		"message": "Workout session berhasil diperbarui",
		"data":    session,
	})
}

// ========================== DELETE ==========================
func DeleteWorkoutSession(c *gin.Context) {
	id := c.Param("id")
	var session models.WorkoutSession

	if err := config.DB.First(&session, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Workout session tidak ditemukan"})
		return
	}

	if err := config.DB.Delete(&session).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghapus session"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Workout session berhasil dihapus",
	})
}
