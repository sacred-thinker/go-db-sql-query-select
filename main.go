package main

import (
	"database/sql"
	"fmt"
	"log"
	_ "modernc.org/sqlite"
)

type Sale struct {
	Product int
	Volume  int
	Date    string
}

// String реализует метод интерфейса fmt.Stringer для Sale, возвращает строковое представление объекта Sale.
// Теперь, если передать объект Sale в fmt.Println(), то выведется строка, которую вернёт эта функция.
func (s Sale) String() string {
	return fmt.Sprintf("Product: %d Volume: %d Date: %s", s.Product, s.Volume, s.Date)
}

func selectSales(client int) ([]Sale, error) {
	var sales []Sale
	//String(sales)
	//sales := Sale{}
	// напишите код здесь
	db, err := sql.Open("sqlite", "demo.db")
	if err != nil {
		log.Println(err)
		return sales, nil
	}
	defer db.Close()

	rows, err := db.Query("SELECT Product, Volume, Date FROM sales")
	if err != nil {
		log.Println(err)
		return sales, nil
	}
	defer rows.Close()

	for rows.Next() {
		saleses := Sale{}

		err := rows.Scan(&saleses.Product, &saleses.Volume, &saleses.Date)
		if err != nil {
			log.Println(err)
			return sales, nil
		}

		// fmt.Println(sales)
	}

	if err := rows.Err(); err != nil {
		log.Println(err)
		return sales, nil
	}
	return sales, nil
}

// return sales, nil

func main() {
	client := 208

	sales, err := selectSales(client)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, sale := range sales {
		fmt.Println(sale)
	}
}
