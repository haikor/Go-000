package dao

import (
	"database/sql"

	"github.com/pkg/errors"
)

type User struct {
	Id   int64
	Name string
}

func FindUserByNameLIke(name string) (*User, error) {
	err := sql.ErrNoRows
	return nil, errors.Wrap(err, "data not found")
}
