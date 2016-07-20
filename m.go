package main

import (
	"fmt"

	"github.com/tidwall/buntdb"
)

func main() {
	db, err := buntdb.Open("data.db")

	if err != nil {
		panic(err)
	}

	db.Update(func(tx *buntdb.Tx) error {
		tx.Set("a", "1", nil)
		tx.Set("b", "2", nil)
		tx.Set("c", "3", nil)

		return nil
	})

	db.View(func(tx *buntdb.Tx) error {
		a, _ := tx.Get("a")
		b, _ := tx.Get("b")
		c, _ := tx.Get("c")

		fmt.Println(a, b, c)

		return nil
	})
}
