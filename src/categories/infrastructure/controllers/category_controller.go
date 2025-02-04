package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"holamundo/src/core"
	"holamundo/src/categories/application"
	"holamundo/src/categories/domain/entities"
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
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo crear la categoría"})
		return
	}
	// Notifica el cambio en categorías
	core.NotifyCategoryUpdate()
	ctx.JSON(http.StatusCreated, category)
}

func (c *CategoryController) ListCategories(ctx *gin.Context) {
	categories, err := c.listUseCase.Execute()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudieron obtener categorías"})
		return
	}
	ctx.JSON(http.StatusOK, categories)
}

// Short polling: devuelve inmediatamente la lista de categorías.
func (c *CategoryController) ListCategoriesShortPolling(ctx *gin.Context) {
	categories, err := c.listUseCase.Execute()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudieron obtener categorías"})
		return
	}
	ctx.JSON(http.StatusOK, categories)
}

func (c *CategoryController) ListCategoriesLongPolling(ctx *gin.Context) {
	select {
	case <-core.CategoryNotifier:
	case <-time.After(30 * time.Second):
	}
	categories, err := c.listUseCase.Execute()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudieron obtener categorías"})
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
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo eliminar la categoría"})
		return
	}
	// Notifica el cambio en categorías
	core.NotifyCategoryUpdate()
	ctx.JSON(http.StatusOK, category)
}

func (c *CategoryController) DeleteCategory(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if err := c.deleteUseCase.Execute(int32(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo eliminar la categoría"})
		return
	}
	// Notifica el cambio en categorías
	core.NotifyCategoryUpdate()
	ctx.JSON(http.StatusOK, gin.H{"message": "Categoria eliminada"})
}
