package storage

import (
	"app/api/models"
	"context"
)

type StorageI interface {
	Close()
	Category() CategoryRepoI
	Product() ProductRepoI
	Market() MarketRepoI
	User() UserRepoI
}

type CategoryRepoI interface {
	Create(context.Context, *models.CreateCategory) (string, error)
	GetByID(context.Context, *models.CategoryPrimaryKey) (*models.Category, error)
	GetList(context.Context, *models.CategoryGetListRequest) (*models.CategoryGetListResponse, error)
	Update(context.Context, *models.UpdateCategory) (int64, error)
	Delete(context.Context, *models.CategoryPrimaryKey) error
}

type ProductRepoI interface {
	Create(context.Context, *models.CreateProduct) (string, error)
	GetByID(context.Context, *models.ProductPrimaryKey) (*models.Product, error)
	GetList(context.Context, *models.ProductGetListRequest) (*models.ProductGetListResponse, error)
	Update(context.Context, *models.UpdateProduct) (int64, error)
	Patch(context.Context, *models.PatchRequest) (int64, error)
	Delete(context.Context, *models.ProductPrimaryKey) error
}

type MarketRepoI interface {
	Create(context.Context, *models.CreateMarket) (string, error)
	GetByID(context.Context, *models.MarketPrimaryKey) (*models.Market, error)
	GetList(context.Context, *models.MarketGetListRequest) (*models.MarketGetListResponse, error)
	Update(context.Context, *models.UpdateMarket) (int64, error)
	Patch(context.Context, *models.PatchRequest) (int64, error)
	Delete(context.Context, *models.MarketPrimaryKey) error
}
type UserRepoI interface {
	Create(ctx context.Context, req *models.CreateUser) (string, error)
	GetByID(ctx context.Context, req *models.UserPrimaryKey) (*models.User, error)
	GetList(ctx context.Context, req *models.GetListUserRequest) (resp *models.GetListUserResponse, err error)
	Update(ctx context.Context, req *models.UpdateUser) (int64, error)
	Delete(ctx context.Context, req *models.UserPrimaryKey) (int64, error)
}
