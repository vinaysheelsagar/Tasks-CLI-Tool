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

func CategoryTableCheckup(db *sql.DB) {

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

func CreateCategory(db *sql.DB, name string, priority int) (int64, error) {

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

func ReadCategory(db *sql.DB, id int64) (Category, error) {
	category := Category{}

	row := db.QueryRow(`SELECT * FROM categories WHERE id=?`, id)

	err := row.Scan(&category.ID, &category.Name, &category.Priority)
	if err != nil {
		return category, err
	}

	return category, nil
}

func UpdateCategory(db *sql.DB, category Category) (Category, error) {

	_, err := db.Exec(
		`UPDATE categories SET name = ?, priority = ? WHERE id=?`,
		category.Name,
		category.Priority,
		category.ID,
	)

	if err != nil {
		return category, err
	}

	category, err = ReadCategory(db, category.ID)
	if err != nil {
		return category, err
	}

	return category, nil
}

func GetCategoryID(db *sql.DB, nameOrID string) (int64, error) {
	var id int64

	response := db.QueryRow(`SELECT id FROM categories WHERE name=? OR id=?`, nameOrID, nameOrID)

	err := response.Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func DeleteCategory(db *sql.DB, id int64) error {

	// TODO: remove id from tasks

	_, err := db.Exec(`DELETE FROM categories WHERE id=?`, id)

	return err
}

func ListCategory(db *sql.DB) ([]Category, error) {
	var categories []Category

	rows, err := db.Query(`SELECT * FROM categories`)
	defer rows.Close()

	utilities.CheckNil(err, "", "")

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
