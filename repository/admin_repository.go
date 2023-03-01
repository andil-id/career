package repository

import (
	"career/model/domain"
	"context"
	"database/sql"
)

type AdminRepository interface {
	GetAdminByUsername(ctx context.Context, db *sql.DB, username string) (domain.Admin, error)
}
