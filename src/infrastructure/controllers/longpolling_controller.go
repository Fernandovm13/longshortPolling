package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"holamundo/src/core"
	productApp "holamundo/src/products/application"
	categoryApp "holamundo/src/categories/application"
)

type LongPollingController struct {
	listProductUseCase  *productApp.ListProductUseCase
	listCategoryUseCase *categoryApp.ListCategoryUseCase
}

func NewLongPollingController(lp *productApp.ListProductUseCase, lc *categoryApp.ListCategoryUseCase) *LongPollingController {
	return &LongPollingController{
		listProductUseCase:  lp,
		listCategoryUseCase: lc,
	}
}

// LongPoll espera hasta que se notifique un cambio en productos o categorías o se cumpla el timeout,
// y luego retorna la lista de ambos.
func (c *LongPollingController) LongPoll(ctx *gin.Context) {
	select {
	case <-core.ProductNotifier:
	case <-core.CategoryNotifier:
	case <-time.After(30 * time.Second):
		// Timeout: no se recibió notificación
	}
	products, err := c.listProductUseCase.Execute()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	categories, err := c.listCategoryUseCase.Execute()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"products":   products,
		"categories": categories,
	})
}
