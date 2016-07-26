package main

import "database/sql"
import _ "github.com/Boostport/avatica"

import (
	"fmt"
	"encoding/json"
)

func main() {
	db, err := sql.Open("avatica", "http://52.208.47.78:8765")
	if err != nil {
		fmt.Print("error:%s", err)
	}
	rows, err := db.Query("SELECT * FROM hello")
	if err != nil {
		fmt.Print("error:%s", err)
	}
	fmt.Print("rows:%s", rows)

	b, err := json.Marshal(rows)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("con chim :%s", string(b))

}
