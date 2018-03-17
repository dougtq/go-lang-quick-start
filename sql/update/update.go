package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@/cursogo")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	smt, _ := db.Prepare("update usaurios set name = ? where id = ?")

	smt.Exec("Teste1", 1)
	smt.Exec("Teste2", 2)

	smt2, _ := db.Prepare("delete from usaurios where id = ?")
	smt2.Exec(1)
}
