package postgres

import (
	"app/api/models"
	"context"
	"database/sql"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

type saleRepo struct {
	db *pgxpool.Pool
}

func NewSalesRepo(db *pgxpool.Pool) *saleRepo {
	return &saleRepo{
		db: db,
	}
}

func (r *saleRepo) GetByID(ctx context.Context, req *models.SalePrimaryKey) (*models.Sale, error) {

	var (
		query     string
		sales     models.Sale
		id        sql.NullString
		user_id   sql.NullString
		total     sql.NullInt64
		count     sql.NullInt64
		createdAt sql.NullString
		updatedAt sql.NullString
	)
	query = `
		SELECT
			id,
			user_id,
			total,
			count,
			CAST(created_at::timestamp AS VARCHAR),
			CAST(updated_at::timestamp AS VARCHAR)
		FROM sales
		WHERE id = $1
	`
	err := r.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&user_id,
		&total,
		&count,
		&createdAt,
		&updatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &sales, nil
}

func (r *saleRepo) GetList(ctx context.Context, req *models.SaleRequest) (*models.SaleResponse, error) {
	var (
		resp   = &models.SaleResponse{}
		query  string
		where  = " WHERE TRUE"
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	query = `
		SELECT
			COUNT(*) OVER(),
			id,
			user_id,
			total,
			count,
			created_at,
			updated_at
		FROM sales
	`

	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	if req.Search != "" {
		where += ` AND user_id ILIKE '%' || '` + req.Search + `' || '%'`
	}

	query += where + offset + limit

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			id        sql.NullString
			user_id   sql.NullString
			total     sql.NullInt64
			count     sql.NullInt64
			createdAt sql.NullString
			updatedAt sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&id,
			&user_id,
			&total,
			&count,
			&createdAt,
			&updatedAt,
		)

		if err != nil {
			return nil, err
		}

		resp.Sales = append(resp.Sales, &models.Sale{
			Id:        id.String,
			UserId:    user_id.String,
			Total:     int(total.Int64),
			Count:     int(count.Int64),
			CreatedAt: createdAt.String,
			UpdatedAt: updatedAt.String,
		})
	}

	return resp, nil
}
