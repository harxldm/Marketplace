package Storage

import (
	model "backend_en_go/Model"
	"database/sql"
	"fmt"
)

const (
	psqlCreateProduct = `INSERT INTO product (name1, sku, amount, price, selerID, created_date_product) 
	                     VALUES ($1, $2, $3, $4, $5, $6) RETURNING productID`

	psqlGetAllProduct = `SELECT productID, name1, sku, amount, price, selerID, created_date_product FROM product`

	psqlGetProductsBySellerID = `SELECT productID, name1, sku, amount, price, selerID, created_date_product 
	                              FROM product WHERE selerID = $1`
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

func (p *psqlProduct) GetAll() ([]model.Product, error) {
	rows, err := p.db.Query(psqlGetAllProduct)
	if err != nil {
		return nil, fmt.Errorf("error al ejecutar la consulta SQL: %w", err)
	}
	defer rows.Close()

	var products []model.Product
	for rows.Next() {
		var p model.Product
		if err := rows.Scan(
			&p.ProductID,
			&p.Name,
			&p.SKU,
			&p.Amount,
			&p.Price,
			&p.SellerID,
			&p.CreatedAt,
		); err != nil {
			return nil, fmt.Errorf("error al leer la fila: %w", err)
		}
		products = append(products, p)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error al iterar sobre las filas: %w", err)
	}

	return products, nil
}

func (p *psqlProduct) GetProductsBySellerID(sellerID int) ([]model.Product, error) {
	rows, err := p.db.Query(psqlGetProductsBySellerID, sellerID)
	if err != nil {
		return nil, fmt.Errorf("error al ejecutar la consulta SQL: %w", err)
	}
	defer rows.Close()

	var products []model.Product
	for rows.Next() {
		var p model.Product
		if err := rows.Scan(
			&p.ProductID,
			&p.Name,
			&p.SKU,
			&p.Amount,
			&p.Price,
			&p.SellerID,
			&p.CreatedAt,
		); err != nil {
			return nil, fmt.Errorf("error al leer la fila: %w", err)
		}
		products = append(products, p)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error al iterar sobre las filas: %w", err)
	}

	return products, nil
}
