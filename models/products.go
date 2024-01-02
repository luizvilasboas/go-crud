package models

import (
	"database/sql"
)

type Product struct {
	Id          int
	Name        string
	Description string
	Price       float64
	Quantity    int
}

func GetAllProducts(db *sql.DB) ([]Product, error) {
	query := "SELECT id, name, description, price, quantity FROM products"
	rows, err := db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	products := []Product{}

	for rows.Next() {
		var p Product

		err = rows.Scan(&p.Id, &p.Name, &p.Description, &p.Price, &p.Quantity)
		if err != nil {
			return nil, err
		}

		products = append(products, p)
	}

	return products, nil
}
