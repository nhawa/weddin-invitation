package main

import (
	"datang-yaa/controller"
	"datang-yaa/repository"
	"datang-yaa/service"
)

func (a *application) registerRoutes() {
	// repository
	customerRepo := repository.NewCustomerRepository(a.DB)
	orderRepo := repository.NewOrderRepository(a.DB)

	// Service
	adminService := service.NewAdminService(customerRepo)
	//categoryService := service.NewCategoryService(viper.GetString("majalah.host"))

	// Controller
	articleController := controller.NewArticleController()
	adminController := controller.NewAdminController(adminService, customerRepo, orderRepo)

	// Route
	a.echo.GET("/", articleController.GetArticleList)
	a.echo.GET("/admin", adminController.GetAdmin)
	a.echo.GET("/admin/orders", adminController.GetOrders)
	//a.echo.GET("/categories", categoryController.GetCategoryList)
}
