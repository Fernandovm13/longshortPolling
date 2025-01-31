package controllers

import (
	"github.com/gin-gonic/gin"
	"holamundo/src/categories/application"
	"holamundo/src/categories/domain/entities"
	"net/http"
	"strconv"
)

type CategoryController struct {
	createUseCase *application.CreateCategoryUseCase
	listUseCase   *application.ListCategoryUseCase
	updateUseCase *application.UpdateCategoryUseCase
	deleteUseCase *application.DeleteCategoryUseCase
}

func NewCategoryController(create *application.CreateCategoryUseCase, list *application.ListCategoryUseCase, update *application.UpdateCategoryUseCase, deleteUC *application.DeleteCategoryUseCase) *CategoryController {
	return &CategoryController{
		createUseCase: create,
		listUseCase:   list,
		updateUseCase: update,
		deleteUseCase: deleteUC,
	}
}

func (c *CategoryController) CreateCategory(ctx *gin.Context) {
	var category entities.Category
	if err := ctx.ShouldBindJSON(&category); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	if err := c.createUseCase.Execute(&category); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create category"})
		return
	}
	ctx.JSON(http.StatusCreated, category)
}

func (c *CategoryController) ListCategories(ctx *gin.Context) {
	categories, err := c.listUseCase.Execute()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch categories"})
		return
	}
	ctx.JSON(http.StatusOK, categories)
}

func (c *CategoryController) UpdateCategory(ctx *gin.Context) {
	var category entities.Category
	if err := ctx.ShouldBindJSON(&category); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	if err := c.updateUseCase.Execute(&category); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update category"})
		return
	}
	ctx.JSON(http.StatusOK, category)
}

func (c *CategoryController) DeleteCategory(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if err := c.deleteUseCase.Execute(int32(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete category"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Category deleted"})
}
