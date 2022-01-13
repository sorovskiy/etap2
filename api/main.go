package main

import (
	"database/sql"
	"etap2/handlers"
	"log"
	"net/http"

	"etap2/repository"
	"github.com/go-chi/chi/v5"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "store.db")
	if err != nil {
		log.Println(err)
	}
	_, err = db.Exec("delete from States")

	rep := repository.NewRepository(db)
	handler := handlers.NewParamsSetter(rep)

	r := chi.NewRouter()
	r.Mount("/", handler.Routes())
	log.Fatal(http.ListenAndServe(":5000", r))

}
