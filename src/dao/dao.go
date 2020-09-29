package dao

import (
	domain "github.com/gdgtoledo/linneo/src/plants"
)

// Query to Dao search
type Query string

// ID to Dao delete
type ID string

// Response to Dao action
type Response struct {
	message  string
	status   int
	response domain.Plants
}

// Interface to the data access object
type Interface interface {
	Search(query Query) (Response, error)
	Delete(id ID) (Response, error)
}
