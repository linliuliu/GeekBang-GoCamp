package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"log"
)


func QueryNameById(id int) (name string,err error){
	db,err :=sql.Open("mysql","admin:admin@tcp4(localhost:3306)/test")
	if err != nil{
		panic(err)
	}

	err = db.QueryRow("SELECT name FROM test1 WHERE id=?", id).Scan(&name)
	switch {
	case err == sql.ErrNoRows:
		//log.Printf("No user with that ID.")
		err = errors.Wrap(err, "No user with that ID")
	case err != nil:
		//log.Fatal(err)
		err = errors.Wrap(err, "Query Failed")
	default:
		log.Printf("GET name success ,name is %s\n", name)
	}
	return name,err
}

func main(){
     _,err := QueryNameById(0)
     if err !=nil {
     	fmt.Printf("original error: %T %V\n ",errors.Cause(err),errors.Cause(err))
     }
}
