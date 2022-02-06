package repository

import (
	"time"

	"datang-yaa/model"
	"datang-yaa/utils"

	"github.com/jmoiron/sqlx"
)

type (
	CustomerRepository interface {
		Repository
		FindAll() (customers []model.Customer, err error)
		FindOneById() (customer model.Customer, err error)
	}
	customerRepository struct {
		BaseRepository
	}
)

var (
	findOneClientByIdStmt *sqlx.Stmt
	findClientsStmt       *sqlx.Stmt
)

func NewCustomerRepository(db *sqlx.DB) CustomerRepository {
	c := customerRepository{}
	c.DB = db
	c.contextTimeout = time.Second * 2

	return c
}

func (c customerRepository) FindAll() (customers []model.Customer, err error) {
	customers = []model.Customer{}
	query := `SELECT id, first_name, last_name, phone_number, enabled, created_at
		FROM customer
	`

	err = c.DB.Select(&customers, query)

	if err != nil {
		utils.LogError(err)
	}

	return
}

func (c customerRepository) FindOneById() (customer model.Customer, err error) {
	ctx, cancel := c.GetContext()
	defer cancel()

	if findOneClientByIdStmt == nil {
		findOneClientByIdStmt, err = c.DB.Preparex("select * from oauth_client where id=$1 LIMIT 1")

		if err != nil {
			return customer, err
		}
	}
	// customer.CreatedAt.Format()
	err = findOneClientByIdStmt.GetContext(ctx, &customer)
	if err != nil {
		utils.LogError(err)
	}

	return
}
