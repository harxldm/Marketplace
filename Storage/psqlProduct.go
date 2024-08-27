package Storage

import (
	model "backend_en_go/Model"
	"database/sql"
	"fmt"
)

const (
	psqlCreateProduct = `INSERT INTO product (name1, sku, amount, price, selerID, created_date_product) 
	                     VALUES ($1, $2, $3, $4, $5, $6) RETURNING productID`
)

type psqlProduct struct {
	db *sql.DB
}

func NewPsqlProduct(db *sql.DB) *psqlProduct {
	return &psqlProduct{db: db}
}

// CreateProduct inserta un nuevo producto en la base de datos
func (p *psqlProduct) CreateProduct(m *model.Product) error {
	stmt, err := p.db.Prepare(psqlCreateProduct)
	if err != nil {
		return fmt.Errorf("error preparando la consulta SQL: %w", err)
	}
	defer stmt.Close()

	// Ejecutar la declaración con los parámetros adecuados
	err = stmt.QueryRow(
		m.Name,
		m.SKU,
		m.Amount,
		m.Price,
		m.SellerID,
		m.CreatedAt,
	).Scan(&m.ProductID)
	if err != nil {
		return fmt.Errorf("error ejecutando la consulta SQL: %w", err)
	}

	return nil
}
