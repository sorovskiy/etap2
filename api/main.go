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
	// query := `TRUNCATE TABLE orders`
	// pool.Exec(context.Background(), query)

	r := chi.NewRouter()
	r.Mount("/", handler.Routes())
	log.Fatal(http.ListenAndServe(":5000", r))

}

func main2() {
	//
	//resp := domain.APIResp{Application: "first-client", Param1: 10, Param2: "abc", Version: 1}
	//resp2 := domain.APIResp{Application: "sec-client", Param1: 10, Param2: "abc", Version: 1}
	//
	//
	//_, err = db.Exec("insert into States (application, param1, param2, version) values ($1, $2, $3, $4)",
	//	resp.Application, resp.Param1, resp.Param2, resp.Version)
	//_, err = db.Exec("insert into States (application, param1, param2, version) values ($1, $2, $3, $4)",
	//	resp2.Application, resp2.Param1, resp2.Param2, resp2.Version)
	//
	//
	//
	//resp3 := domain.APIResp{Application: "third-client", Param1: 10, Param2: "abc", Version: 1}
	//row := db.QueryRow("select * from States where application = $1", resp3.Application)
	//
	//prod := domain.APIResp{}
	//err = row.Scan(&prod.Id, &prod.Application, &prod.Param1, &prod.Param2, &prod.Version)
	//if err != nil {
	//	_, err = db.Exec("insert into States (application, param1, param2, version) values ($1, $2, $3, $4)",
	//		resp.Application, resp.Param1, resp.Param2, 1)
	//} else {
	//	if resp.Param1 == prod.Param1 && resp.Param1 == prod.Param1 {
	//		return
	//	}
	//	_, err = db.Exec("update States (application, param1, param2, version) values ($1, $2, $3, $4)",
	//		resp.Application, resp.Param1, resp.Param2, 1)
	//
	//	db.Exec("update Products set version = $1 where application = $2", prod.Version + 1, resp.Application)
	//}

	//fmt.Println(prod)

}
