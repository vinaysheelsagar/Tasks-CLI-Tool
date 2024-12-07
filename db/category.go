package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/vinaysheelsagar/Tasks-CLI-Tool/utilities"
)

type Category struct {
	ID       int64
	Name     string
	Priority int
}

func CategoryTableCheckup() {
	db := getDB()

	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS "categories" (
		"id"	INTEGER,
		"name"	TEXT,
		"priority"	TEXT,

		PRIMARY KEY("id" AUTOINCREMENT)
	)`)
	if err != nil {
		log.Fatal(err)
	}

}

func CreateCategory(name string, priority int) (int64, error) {
	db := getDB()

	result, err := db.Exec(
		`INSERT INTO categories (name, priority) VALUES (?, ?)`,
		name,
		priority,
	)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return id, err
	}

	return id, nil
}

func ReadCategory(id int64) (Category, error) {
	db := getDB()

	row := db.QueryRow(`SELECT * FROM categories WHERE id=?`, id)

	defer db.Close()

	category := Category{}
	err := row.Scan(&category.ID, &category.Name, &category.Priority)
	if err != nil {
		return category, err
	}

	return category, nil
}

func UpdateCategory(category Category) (Category, error) {

	db := getDB()

	_, err := db.Exec(
		`UPDATE categories SET name = ?, priority = ? WHERE id=?`,
		category.Name,
		category.Priority,
		category.ID,
	)

	defer db.Close()
	if err != nil {
		return category, err
	}

	category, err = ReadCategory(category.ID)
	if err != nil {
		return category, err
	}

	return category, nil
}

func GetCategoryID(nameOrID string) (int64, error) {
	db := getDB()

	response := db.QueryRow(`SELECT id FROM categories WHERE name=? OR id=?`, nameOrID, nameOrID)

	defer db.Close()

	var id int64

	err := response.Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func DeleteCategory(id int64) error {
	db := getDB()

	_, err := db.Exec(`DELETE FROM categories WHERE id=?`, id)

	defer db.Close()

	return err
}

func ListCategory() ([]Category, error) {
	db := getDB()

	rows, err := db.Query(`SELECT * FROM categories`)

	defer db.Close()
	defer rows.Close()
	utilities.CheckNil(err, "", "")

	var categories []Category

	for rows.Next() {
		category := Category{}
		err := rows.Scan(&category.ID, &category.Name, &category.Priority)
		if err != nil {
			if err == sql.ErrNoRows {
				fmt.Println("ERROR: No such ID")
				os.Exit(0)
			} else {
				utilities.CheckNil(err, "", "")
			}
		}
		categories = append(categories, category)
	}
	if err = rows.Err(); err != nil {
		return categories, err
	}
	return categories, nil
}
