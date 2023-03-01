package repository

import (
	"career/model/domain"
	"context"
	"database/sql"
	"errors"
)

type AdminRepositoryImpl struct{}

func NewAdminRepository() AdminRepository {
	return &AdminRepositoryImpl{}
}

func (r *AdminRepositoryImpl) GetAdminByUsername(ctx context.Context, db *sql.DB, username string) (domain.Admin, error) {
	SQL := "SELECT * FROM admin WHERE username = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, SQL, username)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	admin := domain.Admin{}
	if rows.Next() {
		err := rows.Scan(&admin.Id, &admin.Username, &admin.Password, &admin.CreatedAt, &admin.UpdatedAt)
		if err != nil {
			panic(err)
		}
		return admin, nil
	} else {
		return admin, errors.New("data not found")
	}
}
