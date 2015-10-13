package db

import (
	"errors"
)

// define the errors not found
var ErrNotFound = errors.New("Not found")

// A DB interface 
type DB interface{
	Get(key string) (string, error)

	Set(key string, val string) error
}

// Our in-memory DB
type memoryDB struct{
	m map[string] string
}

// NewMemoryDB creates a new DB implementation that stores all data in memory
func NewMemoryDB() DB{
	return &memoryDB{m : make(map[string]string)}
}

func (db *memoryDB) Get(key string) (string, error){
	v, ok := db.m[key]

	if !ok{
		return "", ErrNotFound
	}

	return v, nil
}

func (db *memoryDB) Set(key string, val string) error{
	db.m[key] = val
	return nil
}