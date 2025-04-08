package main

import (
	"database/sql"
	"time"
)

type product struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	Stock     int       `json:"stock"`
	CreatedAt time.Time `json:"createdAt"`
}

func (p *product) getProduct(db *sql.DB) error {
	return db.QueryRow("SELECT name, price, stock, created_at FROM products WHERE id=$1",
		p.ID).Scan(&p.Name, &p.Price, &p.Stock, &p.CreatedAt)
}

func (p *product) updateProduct(db *sql.DB) error {
	_, err := db.Exec("UPDATE products SET name=$1, price=$2, stock=$3 WHERE id=$4",
		p.Name, p.Price, p.Stock, p.ID)

	return err
}

func (p *product) deleteProduct(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM products WHERE id=$1", p.ID)

	return err
}

func (p *product) createProduct(db *sql.DB) error {
	err := db.QueryRow(
		"INSERT INTO products(name, price, stock) VALUES($1, $2, $3) RETURNING id",
		p.Name, p.Price, p.Stock).Scan(&p.ID)
	if err != nil {
		return err
	}

	return nil
}

func getProducts(db *sql.DB, start, count int, search string) ([]product, error) {
	rows, err := db.Query(
		"SELECT id, name, price, stock, created_at FROM products WHERE name LIKE $1 ORDER BY created_at DESC LIMIT $2 OFFSET $3",
		"%"+search+"%", count, start)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	products := []product{}

	for rows.Next() {
		var p product
		if err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.Stock, &p.CreatedAt); err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}
