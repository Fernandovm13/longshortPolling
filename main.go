package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"

	productApp "holamundo/src/products/application"
	productRepo "holamundo/src/products/infrastructure/repositories"
	productCtrl "holamundo/src/products/infrastructure/controllers"
	productInfra "holamundo/src/products/infrastructure"

	categoryApp "holamundo/src/categories/application"
	categoryRepo "holamundo/src/categories/infrastructure/repositories"
	categoryCtrl "holamundo/src/categories/infrastructure/controllers"
	categoryInfra "holamundo/src/categories/infrastructure"

	longPollingCtrl "holamundo/src/infrastructure/controllers"
	"holamundo/src/core"
)

func simulateChanges() {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()
	for {
		<-ticker.C
		// Simula un cambio en productos
		core.NotifyProductUpdate()
		fmt.Println("[SIMULATION]  producto simulada")
		
		// Espera 3 segundos y simula un cambio en categorías
		time.Sleep(3 * time.Second)
		core.NotifyCategoryUpdate()
		fmt.Println("[SIMULATION]  categoría simulada")
	}
}

func main() {
	// Inicia la simulación de notificaciones constantes
	go simulateChanges()

	pRepo := productRepo.NewMySQLProductRepository()
	createProductUC := productApp.NewCreateProductUseCase(pRepo)
	listProductUC := productApp.NewListProductUseCase(pRepo)
	updateProductUC := productApp.NewUpdateProductUseCase(pRepo)
	deleteProductUC := productApp.NewDeleteProductUseCase(pRepo)
	pController := productCtrl.NewProductController(createProductUC, listProductUC, updateProductUC, deleteProductUC)

	cRepo := categoryRepo.NewMySQLCategoryRepository()
	createCategoryUC := categoryApp.NewCreateCategoryUseCase(cRepo)
	listCategoryUC := categoryApp.NewListCategoryUseCase(cRepo)
	updateCategoryUC := categoryApp.NewUpdateCategoryUseCase(cRepo)
	deleteCategoryUC := categoryApp.NewDeleteCategoryUseCase(cRepo)
	cController := categoryCtrl.NewCategoryController(createCategoryUC, listCategoryUC, updateCategoryUC, deleteCategoryUC)

	// Controlador unificado para long polling
	lpController := longPollingCtrl.NewLongPollingController(listProductUC, listCategoryUC)

	r := gin.Default()

	productInfra.SetupProductRoutes(r, pController)
	categoryInfra.SetupCategoryRoutes(r, cController)

	// Ruta única para long polling (retorna ambas listas)
	r.GET("/longpoll", lpController.LongPoll)

	r.Run(":8080")
}
