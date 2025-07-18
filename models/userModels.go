package models

import "github.com/lib/pq"

type PublicEmployeeRequest struct {
	Email string `json:"email" validate:"required,email"`
}

type EmployeeRegisterByStaffRequest struct {
	Username string `json:"userName" validate:"required,username"`
	Email    string `json:"email" validate:"required,email"`
	Contact  string `json:"contact" validate:"required,contact"`
	Type     string `json:"type" validate:"required,type"`
}

type UpdateUserRole struct {
	UserID string `json:"userId" validate:"required,user_id"`
	Type   string `json:"type" validate:"required,type"`
}

type GetEmployeeResponse struct {
	ID             string         `json:"id" db:"id"`
	Username       string         `json:"userName" db:"username"`
	Email          string         `json:"email" db:"email"`
	Contact        string         `json:"contactNo" db:"contact"`
	EmployeeType   string         `json:"employeeType" db:"employee_type"`
	AssignedAssets pq.StringArray `json:"assignedAssets" db:"assigned_assets"`
}
