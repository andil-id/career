package repository

import (
	"career/model/domain"
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/segmentio/ksuid"
)

type JobRespositoryImpl struct{}

func NewJobRespository() JobRepository {
	return &JobRespositoryImpl{}
}

func (r *JobRespositoryImpl) CreateJob(ctx context.Context, tx *sql.Tx, job domain.Job) (string, error) {
	SQL := "INSERT INTO job (id, category_id, company_logo, company_name, location, title, type, banner, description, email, website_url, created_at, updated_at) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?)"
	id := ksuid.New().String()
	_, err := tx.ExecContext(ctx, SQL, id, job.CategoryId, job.CompanyLogo, job.CompanyName, job.Location, job.Title, job.Type, job.Banner, job.Description, job.Email, job.WebsiteUrl, job.CreatedAt, job.UpdatedAt)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (r *JobRespositoryImpl) GetAllJob(ctx context.Context, db *sql.DB, companyName string, categoryId string, limit int, offset int) ([]domain.Job, error) {
	var err error
	var rows *sql.Rows

	if companyName != "" && categoryId != "" {
		SQL := "SELECT * FROM job WHERE category_id = ? AND company_name LIKE ? LIMIT ? OFFSET ?"
		rows, err = db.QueryContext(ctx, SQL, categoryId, companyName, limit, offset)
		if err != nil {
			panic(err)
		}
	} else if companyName != "" {
		SQL := "SELECT * FROM job WHERE company_name LIKE ? LIMIT ? OFFSET ?"
		rows, err = db.QueryContext(ctx, SQL, companyName, limit, offset)
		if err != nil {
			panic(err)
		}
	} else if categoryId != "" {
		SQL := "SELECT * FROM job WHERE category_id = ? LIMIT ? OFFSET ?"
		rows, err = db.QueryContext(ctx, SQL, categoryId, limit, offset)
		if err != nil {
			panic(err)
		}
	} else {
		SQL := "SELECT * FROM job LIMIT ? OFFSET ?"
		rows, err = db.QueryContext(ctx, SQL, limit, offset)
		if err != nil {
			panic(err)
		}
	}
	defer rows.Close()
	jobs := []domain.Job{}
	for rows.Next() {
		job := domain.Job{}
		err := rows.Scan(&job.Id, &job.CategoryId, &job.CompanyLogo, &job.CompanyName, &job.Location, &job.Title, &job.Type, &job.Banner, &job.Description, &job.Email, &job.WebsiteUrl, &job.CreatedAt, &job.UpdatedAt)
		if err != nil {
			panic(err)
		}
		jobs = append(jobs, job)
	}
	return jobs, nil
}

func (r *JobRespositoryImpl) GetJobTotal(ctx context.Context, db *sql.DB, companyName string, categoryId string) (int, error) {
	var err error
	var total int

	if companyName != "" && categoryId != "" {
		SQL := "SELECT COUNT(*) FROM job WHERE category_id = ? AND company_name LIKE ?"
		err = db.QueryRowContext(ctx, SQL, categoryId, companyName).Scan(&total)
		if err != nil {
			panic(err)
		}
	} else if companyName != "" {
		SQL := "SELECT COUNT(*) FROM job WHERE company_name LIKE ?"
		err = db.QueryRowContext(ctx, SQL, companyName).Scan(&total)
		if err != nil {
			panic(err)
		}
	} else if categoryId != "" {
		SQL := "SELECT COUNT(*) FROM job WHERE category_id = ?"
		err = db.QueryRowContext(ctx, SQL, categoryId).Scan(&total)
		if err != nil {
			panic(err)
		}
	} else {
		SQL := "SELECT COUNT(*) FROM job"
		err = db.QueryRowContext(ctx, SQL).Scan(&total)
		if err != nil {
			panic(err)
		}
	}
	return total, nil
}

func (r *JobRespositoryImpl) GetJobById(ctx context.Context, db *sql.DB, jobId string) (domain.Job, error) {
	SQL := "SELECT * FROM job WHERE id = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, SQL, jobId)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	job := domain.Job{}
	if rows.Next() {
		err := rows.Scan(&job.Id, &job.CategoryId, &job.CompanyLogo, &job.CompanyName, &job.Location, &job.Title, &job.Type, &job.Banner, &job.Description, &job.Email, &job.WebsiteUrl, &job.CreatedAt, &job.UpdatedAt)
		if err != nil {
			panic(err)
		}
		return job, nil
	} else {
		return job, errors.New("job not found")
	}
}

func (r *JobRespositoryImpl) UpdateJob(ctx context.Context, tx *sql.Tx, job domain.Job) error {
	SQL := "UPDATE job SET category_id=?, location=?, title=?, type=?, banner=?, description=?, email=?, website_url=?, created_at=?, updated_at=? WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, job.CategoryId, job.Location, job.Title, job.Type, job.Banner, job.Description, job.Email, job.WebsiteUrl, job.CreatedAt, time.Now(), job.Id)
	if err != nil {
		return err
	}
	return nil
}

func (r *JobRespositoryImpl) DeleteJob(ctx context.Context, tx *sql.Tx, jobId string) error {
	SQL := "DELETE FROM job WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, jobId)
	if err != nil {
		return err
	}
	return nil
}
