package crud

import (
	"CRUD/database"
	"CRUD/entity"
	"CRUD/repository"
	"context"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestInsert(t *testing.T) {

	command := repository.NewProductsRepository(database.GetConnection())
	ctx := context.Background()

	products := entity.Products{
		Product_name: "Kartu Perdana Telkomsel",
		Price:        30000,
		Quantity:     15,
	}
	product, err := command.Insert(ctx, products)
	if err != nil {
		panic(err)
	}
	fmt.Println(product)
	fmt.Println("Success Insert Data To Products")
}

func TestFindAll(t *testing.T) {
	command := repository.NewProductsRepository(database.GetConnection())
	ctx := context.Background()
	products, err := command.FindAll(ctx)
	if err != nil {
		panic(err)
	}

	for _, product := range products {
		fmt.Println(product)
	}

}

func TestFindById(t *testing.T) {
	command := repository.NewProductsRepository(database.GetConnection())
	ctx := context.Background()

	products, err := command.FindById(ctx, 7)
	if err != nil {
		panic(err)
	}

	fmt.Println(products)
}

func TestUpdate(t *testing.T) {
	command := repository.NewProductsRepository(database.GetConnection())
	ctx := context.Background()

	product := entity.Products{
		Product_name: "Kartu Perdana Three 13 GB",
		Id:           7,
	}
	_, err := command.Update(ctx, product)
	if err != nil {
		panic(err)
	}
	fmt.Println("Success Update Name Product")

	after, err := command.FindById(ctx, 7)
	if err != nil {
		panic(err)
	}
	fmt.Println(after)
}

func TestDelete(t *testing.T) {
	command := repository.NewProductsRepository(database.GetConnection())
	ctx := context.Background()

	product := entity.Products{
		Product_name: "Kartu Perdana Axis",
	}

	_, err := command.Delete(ctx, product)
	if err != nil {
		panic(err)
	}
	fmt.Println("Success Delete Products By Name")

	results, err := command.FindAll(ctx)
	if err != nil {
		panic(err)
	}

	for _, result := range results {
		fmt.Println(result)
	}

}
