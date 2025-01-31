package repositories

import (
	"database/sql"
	"fmt"
	"holamundo/src/products/domain/entities"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type MySQLProductRepository struct {
	db *sql.DB
}

func NewMySQLProductRepository() *MySQLProductRepository {
	godotenv.Load()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	return &MySQLProductRepository{db: db}
}

func (repo *MySQLProductRepository) Save(product *entities.Product) error {
	query := "INSERT INTO products (name, price) VALUES (?, ?)"
	_, err := repo.db.Exec(query, product.Name, product.Price)
	return err
}

func (repo *MySQLProductRepository) GetAll() ([]entities.Product, error) {
	query := "SELECT id, name, price FROM products"
	rows, err := repo.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []entities.Product
	for rows.Next() {
		var product entities.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Price); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (repo *MySQLProductRepository) Update(product *entities.Product) error {
	query := "UPDATE products SET name=?, price=? WHERE id=?"
	_, err := repo.db.Exec(query, product.Name, product.Price, product.ID)
	return err
}

func (repo *MySQLProductRepository) Delete(id int32) error {
	query := "DELETE FROM products WHERE id=?"
	_, err := repo.db.Exec(query, id)
	return err
}
