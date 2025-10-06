package handlers

import (
	"gym-membership/config"
	"gym-membership/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ========================== CREATE ==========================
func CreateMembership(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User tidak ditemukan di token"})
		return
	}

	var input struct {
		Name        string  `json:"name" binding:"required"`
		Description string  `json:"description"`
		Price       float64 `json:"price" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	membership := models.Membership{
		Name:        input.Name,
		Description: input.Description,
		Price:       input.Price,
		UserID:      userID.(uint),
	}

	if err := config.DB.Create(&membership).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat membership"})
		return
	}

	
	config.DB.Preload("User").First(&membership, membership.ID)

	c.JSON(http.StatusCreated, gin.H{
		"message": "Membership berhasil dibuat",
		"data":    membership,
	})
}

// ========================== GET ==========================
func GetMemberships(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User tidak ditemukan di token"})
		return
	}

	var memberships []models.Membership
	if err := config.DB.Preload("User").Where("user_id = ?", userID.(uint)).Find(&memberships).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data membership"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Daftar membership berhasil diambil",
		"data":    memberships,
	})
}

// ========================== UPDATE ==========================
func UpdateMembership(c *gin.Context) {
	userID, _ := c.Get("user_id")
	id := c.Param("id")

	var membership models.Membership
	if err := config.DB.Where("id = ? AND user_id = ?", id, userID.(uint)).First(&membership).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Membership tidak ditemukan"})
		return
	}

	var input struct {
		Name        string  `json:"name" binding:"required"`
		Description string  `json:"description"`
		Price       float64 `json:"price" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	membership.Name = input.Name
	membership.Description = input.Description
	membership.Price = input.Price

	if err := config.DB.Save(&membership).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memperbarui membership"})
		return
	}

	config.DB.Preload("User").First(&membership, membership.ID)

	c.JSON(http.StatusOK, gin.H{
		"message": "Membership berhasil diperbarui",
		"data":    membership,
	})
}

// ========================== DELETE ==========================
func DeleteMembership(c *gin.Context) {
	userID, _ := c.Get("user_id")
	id := c.Param("id")

	var membership models.Membership
	if err := config.DB.Where("id = ? AND user_id = ?", id, userID.(uint)).First(&membership).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Membership tidak ditemukan"})
		return
	}

	if err := config.DB.Delete(&membership).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghapus membership"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Membership berhasil dihapus",
	})
}
