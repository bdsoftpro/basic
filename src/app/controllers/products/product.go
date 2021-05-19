package products

import (
	"fmt"
	"strconv"
	"net/http"
	"html/template"
	"app/views"
	"app/config"
	"app/models/products"
)
func getField(r *http.Request, index int) string{
	field := r.Context().Value("fields").([]string)
	return field[index]
}

// Product Methods Defination
func Product(w http.ResponseWriter, r *http.Request) {
	page := getField(r, 0)
	p := struct {
		Name string
	}{
		Name: page,
	}
	pageTpl, _ := resources.Asset("templates/pages/page.html")
	t := template.Must(template.New("page.html").Parse(string(pageTpl)))
	if err := t.Execute(w, p); err != nil {
		panic(err)
	}
}

// ProductItem Methods Defination
func ProductItem(w http.ResponseWriter, r *http.Request) {
	pid, _ := strconv.ParseInt(getField(r, 0), 10, 64)
	db, err := config.GetMySQLDB()
	if err != nil {
		fmt.Println(err)
	} else {
		pdb := pdmodels.ProductDB{
			Db: db,
		}
		product, err := pdb.Find(pid)
		if err != nil {
			pageTpl, _ := resources.Asset("templates/products/notfound.html")
			t := template.Must(template.New("product.html").Parse(string(pageTpl)))
			item := struct {
				Name int64
			}{
				Name: pid,
			}
			if err := t.Execute(w, item); err != nil {
				panic(err)
			}
		} else {
			pageTpl, _ := resources.Asset("templates/products/product.html")
			t := template.Must(template.New("product.html").Parse(string(pageTpl)))
			if err := t.Execute(w, product); err != nil {
				panic(err)
			}
		}
	}
}
// ProductAll Methods Defination
func ProductAll(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Product All\n")
}

// ProductCategory Methods Defination
func ProductCategory(w http.ResponseWriter, r *http.Request) {
	page := getField(r, 0)
	cat := getField(r, 1)
	fmt.Fprintf(w, "Product is %s with %s Category\n", page, cat)
}