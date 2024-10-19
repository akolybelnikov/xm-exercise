package models

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
