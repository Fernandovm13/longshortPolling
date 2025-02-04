package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"holamundo/src/core"
	"holamundo/src/products/application"
	"holamundo/src/products/domain/entities"
)

type ProductController struct {
	createUseCase *application.CreateProductUseCase
	listUseCase   *application.ListProductUseCase
	updateUseCase *application.UpdateProductUseCase
	deleteUseCase *application.DeleteProductUseCase
}

func NewProductController(create *application.CreateProductUseCase, list *application.ListProductUseCase, update *application.UpdateProductUseCase, deleteUC *application.DeleteProductUseCase) *ProductController {
	return &ProductController{createUseCase: create, listUseCase: list, updateUseCase: update, deleteUseCase: deleteUC}
}

func (c *ProductController) CreateProduct(ctx *gin.Context) {
	var product entities.Product
	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.createUseCase.Execute(&product); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// Notifica el cambio en productos
	core.NotifyProductUpdate()
	ctx.JSON(http.StatusCreated, product)
}

func (c *ProductController) ListProducts(ctx *gin.Context) {
	products, err := c.listUseCase.Execute()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, products)
}

// Short polling: devuelve inmediatamente la lista de productos.
func (c *ProductController) ListProductsShortPolling(ctx *gin.Context) {
	products, err := c.listUseCase.Execute()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, products)
}


func (c *ProductController) ListProductsLongPolling(ctx *gin.Context) {
	select {
	case <-core.ProductNotifier:
	case <-time.After(30 * time.Second):
	}
	products, err := c.listUseCase.Execute()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, products)
}

func (c *ProductController) UpdateProduct(ctx *gin.Context) {
	var product entities.Product
	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.updateUseCase.Execute(&product); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// Notifica el cambio en productos
	core.NotifyProductUpdate()
	ctx.JSON(http.StatusOK, product)
}

func (c *ProductController) DeleteProduct(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if err := c.deleteUseCase.Execute(int32(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo eliminar el producto"})
		return
	}
	// Notifica el cambio en productos
	core.NotifyProductUpdate()
	ctx.JSON(http.StatusOK, gin.H{"message": "Producto eliminado"})
}
