package service

import (
	"database/sql"
	"week02/dao"

	"github.com/pkg/errors"
)

func FindNameLIke(name string) (string, error) {
	user, err := dao.FindUserByNameLIke(name)
	if errors.Is(err, sql.ErrNoRows) {
		return "", nil
	}
	if err != nil {
		return "", err
	}
	return user.Name, nil
}
