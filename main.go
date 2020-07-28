package main

import (
	"GoRestApi/database"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

//global
var databaseConn *sql.DB

type Product struct {
	ID           int    `json:"id"`
	Product_Code string `json:"product_code"`
	Description  string `json:"description"`
}

func main() {
	databaseConn = database.InitDb()
	defer databaseConn.Close() //se va a cerrar cuando se dejede usar

	fmt.Println(databaseConn)

	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	r.Get("/products", AllProductos)
	http.ListenAndServe(":3000", r)
}

func respWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _:= json.Marshal(payload)

	fmt.Println(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
func catch(err error) {
	if err != nil {
		panic(err)
	}
}

func AllProductos(w http.ResponseWriter, r *http.Request) {
	const sql = `Select id, product_code,COALESCE(description,'') 
 					from products`
	result, err := databaseConn.Query(sql)
	catch(err)
	var products []*Product
	for result.Next() {
		product := &Product{}
		err = result.Scan(&product.ID, &product.Product_Code, &product.Description)
		catch(err)
		products = append(products, product)
	}
	respWithJSON(w, http.StatusOK, products)
}
