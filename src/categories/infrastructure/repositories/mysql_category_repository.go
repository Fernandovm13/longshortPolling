package repositories

import (
	"database/sql"
	"fmt"
	"holamundo/src/categories/domain/entities"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type MySQLCategoryRepository struct {
	db *sql.DB
}

func NewMySQLCategoryRepository() *MySQLCategoryRepository {
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
	return &MySQLCategoryRepository{db: db}
}

func (repo *MySQLCategoryRepository) Save(category *entities.Category) error {
	query := "INSERT INTO categories (name) VALUES (?)"
	_, err := repo.db.Exec(query, category.Name)
	return err
}

func (repo *MySQLCategoryRepository) GetAll() ([]entities.Category, error) {
	query := "SELECT id, name FROM categories"
	rows, err := repo.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []entities.Category
	for rows.Next() {
		var category entities.Category
		if err := rows.Scan(&category.ID, &category.Name); err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	return categories, nil
}

func (repo *MySQLCategoryRepository) Update(category *entities.Category) error {
	query := "UPDATE categories SET name=? WHERE id=?"
	_, err := repo.db.Exec(query, category.Name, category.ID)
	return err
}

func (repo *MySQLCategoryRepository) Delete(id int32) error {
	query := "DELETE FROM categories WHERE id=?"
	_, err := repo.db.Exec(query, id)
	return err
}
