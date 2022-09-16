package repository

import (
	"CRUD/entity"
	"context"
	"database/sql"
	"errors"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

type ProductsRepositoryImpl struct {
	DB *sql.DB
}

func NewProductsRepository(db *sql.DB) ProductsRepository {
	return &ProductsRepositoryImpl{DB: db}
}

func (repo *ProductsRepositoryImpl) Insert(ctx context.Context, product entity.Products) (entity.Products, error) {
	query := "INSERT INTO products(product_name,price,quantity) VALUES(?,?,?)"
	result, err := repo.DB.ExecContext(ctx, query, product.Product_name, product.Price, product.Quantity)
	if err != nil {
		return product, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return product, err
	}
	product.Id = int32(id)

	return product, nil
}

func (repo *ProductsRepositoryImpl) FindById(ctx context.Context, id int32) (entity.Products, error) {
	query := "SELECT id,product_name,price,quantity,created_at FROM products where id = ? limit 1"
	rows, err := repo.DB.QueryContext(ctx, query, id)
	var product entity.Products
	if err != nil {
		return product, err
	}
	defer rows.Close()
	product.Id = int32(id)
	if rows.Next() {
		rows.Scan(&product.Id, &product.Product_name, &product.Price, &product.Quantity, &product.Created_at)
		return product, nil
	} else {
		return product, errors.New("Id :" + strconv.Itoa(int(id)) + "Not Found")
	}
}

func (repo *ProductsRepositoryImpl) FindAll(ctx context.Context) ([]entity.Products, error) {
	query := "SELECT id,product_name,price,quantity,created_at FROM products"
	rows, err := repo.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var products []entity.Products
	for rows.Next() {
		var product entity.Products
		rows.Scan(&product.Id, &product.Product_name, &product.Price, &product.Quantity, &product.Created_at)
		products = append(products, product)
	}
	return products, nil
}

func (repo *ProductsRepositoryImpl) Delete(ctx context.Context, product entity.Products) (entity.Products, error) {
	query := "DELETE from products where product_name = ?"
	_, err := repo.DB.ExecContext(ctx, query, product.Product_name)
	if err != nil {
		return product, err
	}
	return product, nil
}

func (repo *ProductsRepositoryImpl) Update(ctx context.Context, product entity.Products) (entity.Products, error) {
	query := "UPDATE products set product_name = ? where id = ?"
	_, err := repo.DB.ExecContext(ctx, query, product.Product_name, product.Id)
	if err != nil {
		return product, err
	}
	return product, nil
}
