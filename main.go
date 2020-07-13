package main

import (
	"GoRestApi/database"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// "net/http"
// "github.com/go-chi/chi"
// "github.com/go-chi/chi/middleware"

func main() {
	databaseConn := database.InitDb()
	defer databaseConn.Close() //se va a cerrar cuando se dejede usar

	fmt.Println(databaseConn)
	// r := chi.NewRouter()
	// r.Use(middleware.Logger)
	// r.Get("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("welcome"))
	// })
	// http.ListenAndServe(":3000", r)
}
