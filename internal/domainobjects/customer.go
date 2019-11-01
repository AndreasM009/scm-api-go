package domainobjects

import "github.com/google/uuid"

type Customer struct {
	Id uuid.UUID
	Firstname string
	Lastname string
}