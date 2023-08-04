package postgres

import (
	"app/api/models"
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type salesProductRepo struct {
	db *pgxpool.Pool
}

func NewSalesProductRepo(db *pgxpool.Pool) *salesProductRepo {
	return &salesProductRepo{
		db: db,
	}
}

func (r *salesProductRepo) Create(ctx context.Context, req *models.CreateSaleProduct) (string, error) {
	var (
		product             []models.Product
		query               string
		id                  string
		price_with_discount float64
		discount_price      float64
		total_price         float64
		product_name        string
		produc_price        float64
	)
	id = uuid.NewString()
	for _, val := range product {
		if req.ProductId == val.Id {
			product_name = val.Name
			produc_price = val.Price

		}
	}
	if req.DiscountType == "Fixed" {
		price_with_discount = produc_price - req.Discount
		discount_price = req.Discount
	} else if req.DiscountType == "Precent" {
		price_with_discount = produc_price * (100 - discount_price) / 100
	}
	total_price = price_with_discount * req.Count

	query = `
		INSERT INTO sale_product(
			id,
			product_id,
			product_name,
    		produc_price,
    		discount,
    		discount_type,
    		price_with_discount,
    		discount_price,
    		count,
    		total_price,
    		updated_at
		)
		VALUES ( $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, now())
	`

	_, err := r.db.Exec(ctx, query,
		id,
		req.ProductId,
		product_name,
		produc_price,
		req.Discount,
		req.DiscountType,
		price_with_discount,
		discount_price,
		req.Count,
		total_price,
	)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (r *salesProductRepo) GetByID(ctx context.Context, req *models.SaleProductPrimaryKey) (*models.SaleProduct, error) {

	var (
		query               string
		saleProduct         models.SaleProduct
		id                  sql.NullString
		product_id          sql.NullString
		product_name        sql.NullString
		produc_price        sql.NullFloat64
		discount            sql.NullFloat64
		discount_type       sql.NullString
		price_with_discount sql.NullFloat64
		discount_price      sql.NullFloat64
		count               sql.NullFloat64
		total_price         sql.NullFloat64
		createdAt           sql.NullString
		updatedAt           sql.NullString
	)
	query = `
		SELECT
			id,
			product_id,
			product_name,
			produc_price,
			discount,
			discount_type,
			price_with_discount,
			discount_price,
			count,
			total_price,
			CAST(created_at::timestamp AS VARCHAR),
			CAST(updated_at::timestamp AS VARCHAR)
		FROM sale_product
		WHERE id = $1
	`
	err := r.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&product_id,
		&product_name,
		&produc_price,
		&discount,
		&discount_type,
		&price_with_discount,
		&discount_price,
		&count,
		&total_price,
		&createdAt,
		&updatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &saleProduct, nil
}

// func (r *userRepo) GetList(ctx context.Context, req *models.GetListUserRequest) (resp *models.GetListUserResponse, err error) {

// 	resp = &models.GetListUserResponse{}

// 	var (
// 		query  string
// 		filter = " WHERE TRUE "
// 		offset = " OFFSET 0"
// 		limit  = " LIMIT 10"
// 	)

// 	query = `
// 		SELECT
// 			COUNT(*) OVER(),
// 			id,
// 			first_name,
// 			last_name,
// 			login,
// 			password,
// 			phone_number,
// 			CAST(created_at::timestamp AS VARCHAR),
// 			CAST(updated_at::timestamp AS VARCHAR)
// 		FROM users
// 	`

// 	if len(req.Search) > 0 {
// 		filter += " AND name ILIKE '%' || '" + req.Search + "' || '%' "
// 	}

// 	if req.Offset > 0 {
// 		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
// 	}

// 	if req.Limit > 0 {
// 		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
// 	}

// 	query += filter + offset + limit

// 	rows, err := r.db.Query(ctx, query)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	for rows.Next() {
// 		var user models.User
// 		err = rows.Scan(
// 			&resp.Count,
// 			&user.Id,
// 			&user.FirstName,
// 			&user.LastName,
// 			&user.Login,
// 			&user.Password,
// 			&user.PhoneNumber,
// 			&user.CreatedAt,
// 			&user.UpdatedAt,
// 		)
// 		if err != nil {
// 			return nil, err
// 		}

// 		resp.Users = append(resp.Users, &user)
// 	}

// 	return resp, nil
// }

// func (r *userRepo) Update(ctx context.Context, req *models.UpdateUser) (int64, error) {
// 	var (
// 		query  string
// 		params map[string]interface{}
// 	)

// 	query = `
// 		UPDATE
// 		users
// 		SET
// 			id = :id,
// 			first_name = :first_name,
// 			last_name = :last_name,
// 			login = :login,
// 			password = :password,
// 			phone_number = :phone_number,
// 			updated_at = now()
// 		WHERE id = :id
// 	`

// 	params = map[string]interface{}{
// 		"id":           req.Id,
// 		"first_name":   req.FirstName,
// 		"last_name":    req.LastName,
// 		"login":        req.Login,
// 		"password":     req.Password,
// 		"phone_number": req.PhoneNumber,
// 	}

// 	query, args := helper.ReplaceQueryParams(query, params)

// 	result, err := r.db.Exec(ctx, query, args...)
// 	if err != nil {
// 		return 0, err
// 	}

// 	return result.RowsAffected(), nil
// }

// func (r *userRepo) Delete(ctx context.Context, req *models.UserPrimaryKey) (int64, error) {
// 	query := `
// 		DELETE
// 		FROM users
// 		WHERE id = $1
// 	`

// 	result, err := r.db.Exec(ctx, query, req.Id)
// 	if err != nil {
// 		return 0, err
// 	}

// 	return result.RowsAffected(), nil
// }
