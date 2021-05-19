package pgmodels

import (
	"database/sql"
)

// Page Format
type Page struct {
	Id		int64	`json:"id"`
	Name	string	`json:"name"`
	Price	float32	`json:"price"`
	Quantity int	`json:"quantity"`
	Status	bool	`json:"status"`
}

// PageDB Format
type PageDB struct {
	Db *sql.DB
}

// Create Methos Defination
func (pgdb PageDB) Create(page *Page) error {
	result, err := pgdb.Db.Exec("insert into page(name, price, quantity, status) values(?, ?, ?, ?)", page.Name, page.Price, page.Quantity, page.Status)
	if err != nil {
		return err
	}
	page.Id, _ = result.LastInsertId()
	return nil
}

// Update Methos Defination
func (pgdb PageDB) Update(page Page) (int64, error) {
	result, err := pgdb.Db.Exec("update page set name = ?, price = ?, quantity = ?, status = ? where id = ?", page.Name, page.Price, page.Quantity, page.Status, page.Id)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

// Find Methos Defination
func (pgdb PageDB) Find(id int64) (Page, error) {
	rows, err := pgdb.Db.Query("select * from page where id = ?", id)
	if err != nil {
		return Page{}, err
	}

	page := Page{}
	for rows.Next() {
		var id int64
		var name string
		var price float32
		var quantity int
		var status bool
		err2 := rows.Scan(&id, &name, &price, &quantity, &status)
		if err2 != nil {
			return Page{}, err2
		}
		//page = Page{id, name, price, quantity, status}
		page = Page{
			Id:       id,
			Name:     name,
			Price:    price,
			Quantity: quantity,
			Status:   status,
		}
	}
	return page, nil
}

// FindAll Methos Defination
func (pgdb PageDB) FindAll() ([]Page, error) {
	rows, err := pgdb.Db.Query("select * from page")
	if err != nil {
		return nil, err
	}
	pages := []Page{}
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
		//page := Page{id, name, price, quantity, status}
		page := Page{
			Id:       id,
			Name:     name,
			Price:    price,
			Quantity: quantity,
			Status:   status,
		}
		pages = append(pages, page)
	}
	return pages, nil
}
