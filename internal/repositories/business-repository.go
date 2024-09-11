package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/henrybravo/micro-report/pkg/db"
)

type BusinessRepository struct {
	Connection *db.Connection
}

func NewBusinessRepository(connection *db.Connection) *BusinessRepository {
	return &BusinessRepository{Connection: connection}
}

type Business struct {
	ID           uuid.UUID `json:"id"`
	BusinessName string    `json:"razonSocial"`
	RUC          string    `json:"ruc"`
	Address      string    `json:"direccion"`
}

func (r *BusinessRepository) GetBusinessByID(id string) (*Business, error) {
	getBusinessByIDQuery := `SELECT id, razon_social, ruc, direccion  FROM empresas WHERE id=$1`
	var company Business
	err := r.Connection.Pool.QueryRow(context.Background(), getBusinessByIDQuery, id).Scan(&company.ID, &company.BusinessName, &company.RUC, &company.Address)
	if err != nil {
		return nil, err
	}
	return &company, nil
}
