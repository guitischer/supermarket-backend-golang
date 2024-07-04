package controllers

import (
	"net/http"

	"supermarket-api/src/database"
	"supermarket-api/src/models"

	"github.com/gin-gonic/gin"
)

type CreateProductTypeRequest struct {
	Name string `json:"name" binding:"required"`
}

func GetAllProductTypes(ctx *gin.Context) {
	var productTypes []models.ProductType
	database.DB.Find(&productTypes)

	ctx.JSON(http.StatusOK, gin.H{"data": productTypes})
}

func CreateProductType(ctx *gin.Context) {
	var input CreateProductTypeRequest
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create book
	productType := models.ProductType{Name: input.Name}
	database.DB.Create(&productType)

	ctx.JSON(http.StatusOK, gin.H{"data": productType})
}
