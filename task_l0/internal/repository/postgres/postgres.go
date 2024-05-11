package postgres

import (
	"fmt"
	"module/internal/repository"

	"github.com/jmoiron/sqlx"
)

const (
	OrdersTable     = "orders"
	ItemsTable      = "items"
	DeliveriesTable = "deliveries"
	PaymentsTable   = "payments"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPgSQLConnection(conn Config) (*sqlx.DB, error) {
	connString := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		conn.Username,
		conn.Password,
		conn.Host,
		conn.Port,
		conn.DBName,
		conn.SSLMode,
	)

	db, err := sqlx.Open("postgres", connString)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

func NewPGRepository(db *sqlx.DB) *repository.Repository {
	return &repository.Repository{
		OrderRep: NewOrderPG(db),
	}

}
