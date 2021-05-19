package pdmodels

import (
	"database/sql"
)

// Product Format
type Product struct {
	Id		int64
	Name	string
	Price	float32
	Quantity int
	Status	bool
}

// ProductDB Format
type ProductDB struct {
	Db *sql.DB
}

// Create Methos Defination
func (pdb ProductDB) Create(product *Product) error {
	result, err := pdb.Db.Exec("insert into product(name, price, quantity, status) values(?, ?, ?, ?)", product.Name, product.Price, product.Quantity, product.Status)
	if err != nil {
		return err
	}
	product.Id, _ = result.LastInsertId()
	return nil
}

// Update Methos Defination
func (pdb ProductDB) Update(product Product) (int64, error) {
	result, err := pdb.Db.Exec("update product set name = ?, price = ?, quantity = ?, status = ? where id = ?", product.Name, product.Price, product.Quantity, product.Status, product.Id)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

// Find Methos Defination
func (pdb ProductDB) Find(id int64) (Product, error) {
	rows, err := pdb.Db.Query("select * from product where id = ?", id)
	if err != nil {
		return Product{}, err
	}

	product := Product{}
	for rows.Next() {
		var id int64
		var name string
		var price float32
		var quantity int
		var status bool
		err2 := rows.Scan(&id, &name, &price, &quantity, &status)
		if err2 != nil {
			return Product{}, err2
		}
		//product = Product{id, name, price, quantity, status}
		product = Product{
			Id:       id,
			Name:     name,
			Price:    price,
			Quantity: quantity,
			Status:   status,
		}
	}
	return product, nil
}

// FindAll Methos Defination
func (pdb ProductDB) FindAll() ([]Product, error) {
	rows, err := pdb.Db.Query("select * from product")
	if err != nil {
		return nil, err
	}
	products := []Product{}
	for rows.Next() {
		var id int64
		var name string
		var price float32
		var quantity int
		var status bool
		err2 := rows.Scan(&id, &name, &price, &quantity, &status)
		if err2 != nil {
			return nil, err2
		}
		//product := Product{id, name, price, quantity, status}
		product := Product{
			Id:       id,
			Name:     name,
			Price:    price,
			Quantity: quantity,
			Status:   status,
		}
		products = append(products, product)
	}
	return products, nil
}
