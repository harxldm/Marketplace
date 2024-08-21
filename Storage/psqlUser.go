package Storage

import (
	model "backend_en_go/Model"
	"database/sql"
	"fmt"
)

// type scanner interface {
// 	Scan(dest ...interface{}) error
// }

const (
	psqlCreateUser = `INSERT INTO User1 (email, password1, rol, created_date) VALUES ($1, $2, $3, $4) RETURNING userID`
)

type PsqlUser struct {
	db *sql.DB
}

func NewPsqlUser(db *sql.DB) *PsqlUser {
	return &PsqlUser{db}
}

func (p *PsqlUser) CreateUser(m *model.User) error {
	// Preparar la declaración SQL para insertar un nuevo usuario
	stmt, err := p.db.Prepare(psqlCreateUser)
	if err != nil {
		return fmt.Errorf("error preparando la consulta SQL: %w", err)
	}
	defer stmt.Close()

	// Ejecutar la declaración con los parámetros adecuados
	err = stmt.QueryRow(
		m.Email,
		m.Password,
		m.Rol,
		m.Created_date,
	).Scan(&m.UserID) // Solo hacer Scan si UserID es generado por la DB
	if err != nil {
		return fmt.Errorf("error ejecutando la consulta SQL: %w", err)
	}

	fmt.Println("El registro fue un éxito")
	return nil
}
