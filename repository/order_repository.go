package repository

import (
	"time"

	"datang-yaa/model"
	"datang-yaa/utils"

	"github.com/jmoiron/sqlx"
)

type (
	OrderRepository interface {
		Repository
		FindAll() (orders []model.Order, err error)
	}
	orderRepository struct {
		BaseRepository
	}
)

var (
	findAllOrder *sqlx.Stmt
)

func NewOrderRepository(db *sqlx.DB) OrderRepository {
	c := orderRepository{}
	c.DB = db
	c.contextTimeout = time.Second * 2

	return c
}

func (o orderRepository) FindAll() (orders []model.Order, err error) {
	orders = []model.Order{}
	query := `SELECT c.id, concat(c.first_name,' ', c.last_name) as customer_name, o.template, o.status, o.created_at, o.expired_at
	FROM "order" o
	join customer c on c.id = o.customer_id
	`

	err = o.DB.Select(&orders, query)

	if err != nil {
		utils.LogError(err)
	}

	return
}
