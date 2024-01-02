package models

import "gitlab.com/alura-courses-code/golang/web-crud/db"

type Product struct {
	Id          int
	Name        string
	Description string
	Price       float64
	Quantity    int
}

func GetAllProducts() ([]Product, error) {
	database := db.ConnectDatabase()
	defer database.Close()

	query := "SELECT id, name, description, price, quantity FROM products"
	rows, err := database.Query(query)

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

func InsertProduct(p Product) error {
	database := db.ConnectDatabase()
	defer database.Close()

	query := "INSERT INTO products (name, description, price, quantity) VALUES ($1, $2, $3, $4)"

	stmt, err := database.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(p.Name, p.Description, p.Price, p.Quantity)
	if err != nil {
		return err
	}

	return nil
}
