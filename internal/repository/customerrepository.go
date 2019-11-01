package repository

import (
	"scm-api/internal/domainobjects"

	"github.com/google/uuid"
)

// Repository interface
type Repository interface {
	Add(c *domainobjects.Customer) domainobjects.Customer
	Get(id uuid.UUID) domainobjects.Customer
	GetAll() []domainobjects.Customer
	DeleteAll() []domainobjects.Customer
}
