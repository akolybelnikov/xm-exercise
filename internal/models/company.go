package models

import (
	"encoding/hex"
	"errors"
	"net/http"

	"github.com/akolybelnikov/xm-exercise/db"
)

type CreateRequest struct {
	Name          string `json:"name" validate:"required,max=15"`
	Description   string `json:"description" validate:"max=3000"`
	EmployeeCount int32  `json:"employee_count" validate:"required"`
	Registered    string `json:"registered" validate:"required,oneof='true' 'false'"`
	CompanyType   string `json:"company_type" validate:"required,oneof='Corporations' 'NonProfit' 'Cooperative' 'Sole Proprietorship'"` //nolint:lll
}

type UpdateRequest struct {
	ID            string `json:"id" validate:"required"`
	Name          string `json:"name" validate:"max=15"`
	Description   string `json:"description" validate:"max=3000"`
	EmployeeCount int32  `json:"employee_count"`
	Registered    string `json:"registered" validate:"oneof='true' 'false'"`
	CompanyType   string `json:"company_type" validate:"oneof='Corporations' 'NonProfit' 'Cooperative' 'Sole Proprietorship'"` //nolint:lll
}

type Company struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	EmployeeCount int32  `json:"employee_count"`
	Registered    bool   `json:"registered"`
	CompanyType   string `json:"company_type"`
}

func (c *Company) Render(_ http.ResponseWriter, _ *http.Request) error {
	return nil
}

func FromDBCompany(dbCompany *db.Company) (*Company, error) {
	value, err := dbCompany.ID.UUIDValue()
	if err != nil {
		return nil, errors.New("Error getting the UUID value: " + err.Error())
	}
	hexString := hex.EncodeToString(value.Bytes[:])

	return &Company{
		ID:            hexString,
		Name:          dbCompany.Name,
		Description:   dbCompany.Description.String,
		EmployeeCount: dbCompany.EmployeeCount,
		Registered:    dbCompany.Registered,
		CompanyType:   dbCompany.Type,
	}, nil
}
