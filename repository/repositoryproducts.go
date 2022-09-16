package repository

import (
	"CRUD/entity"
	"context"
)

type ProductsRepository interface {
	Insert(ctx context.Context, product entity.Products) (entity.Products, error)
	FindById(ctx context.Context, id int32) (entity.Products, error)
	FindAll(ctx context.Context) ([]entity.Products, error)
	Update(ctx context.Context, product entity.Products) (entity.Products, error)
	Delete(ctx context.Context, product entity.Products) (entity.Products, error)
}
