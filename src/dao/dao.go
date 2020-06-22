package dao

import (
	"github.com/gdgtoledo/linneo/src/domain"
)

// Dao interface for the data access object
type Dao interface {
	Search() (Response, error)
	Delete() (Response, error)
}

// Response to a Dao action
type Response struct {
	message  string
	status   int
	response domain.Plants
}
