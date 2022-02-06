package controller

import (
	"datang-yaa/repository"
	"datang-yaa/service"
	"fmt"
	"net/http"
	_ "strconv"

	//"time"

	"github.com/labstack/echo/v4"
)

type (
	AdminController interface {
		GetAdmin(echo echo.Context) error
		GetOrders(echo echo.Context) error
		GetAdminDetail(echo echo.Context) error
	}

	adminController struct {
		adminService service.AdminService
		customerRepo repository.CustomerRepository
		orderRepo    repository.OrderRepository
	}
)

func NewAdminController(
	adminService service.AdminService,
	customerRepo repository.CustomerRepository,
	orderRepo repository.OrderRepository,
) AdminController {

	a := adminController{}
	a.adminService = adminService
	a.customerRepo = customerRepo
	a.orderRepo = orderRepo

	return a
}

func (a adminController) GetAdmin(c echo.Context) error {
	customers, _ := a.customerRepo.FindAll()
	return c.Render(http.StatusOK, "listCustomer", map[string]interface{}{
		"active":    "listCustomer",
		"customers": customers,
	})
}

func (a adminController) GetOrders(c echo.Context) error {
	orders, _ := a.orderRepo.FindAll()
	fmt.Println(orders)
	return c.Render(http.StatusOK, "listOrder", map[string]interface{}{
		"active": "listOrder",
		"orders": orders,
	})
}

func (a adminController) GetAdminDetail(echo echo.Context) error {

	var err error
	return err
}
