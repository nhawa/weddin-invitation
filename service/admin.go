package service

import (
	"datang-yaa/repository"
)

type (
	AdminService interface {
	}

	adminService struct {
		customerRepo repository.CustomerRepository
	}
)

func NewAdminService(
	customerRepo repository.CustomerRepository,
) AdminService {

	l := adminService{}
	l.customerRepo = customerRepo

	return l
}
