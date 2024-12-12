package database

import (
	"database/sql"
	"log"

	"github.com/vinaysheelsagar/Tasks-CLI-Tool/utilities"
)

type Task struct {
	ID         int64
	Name       string
	Priority   int
	CategoryID *int64
	Status     bool
}

func TaskTableCheckup(db *sql.DB) {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS "tasks" (
		"id"			INTEGER,
		"name"			TEXT NOT NULL,
		"priority"		INTEGER NOT NULL DEFAULT 5,
		"category_id"	INTEGER,
		"status"		BOOL NOT NULL DEFAULT FALSE,

		PRIMARY KEY("id" AUTOINCREMENT),
		FOREIGN KEY("category_id") REFERENCES category("id")
	)`)
	if err != nil {
		log.Fatal(err)
	}

}

func CreateTask(db *sql.DB, name string, priority int, categoryID *int64) (int64, error) {

	result, err := db.Exec(
		`INSERT INTO tasks (name, priority, category_id) VALUES (?, ?, ?)`,
		name,
		priority,
		categoryID,
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

func ReadTask(db *sql.DB, id int64) (Task, error) {

	row := db.QueryRow(`
	SELECT id, name, priority, category_id, status 
	FROM tasks 
	WHERE id=?`, id)

	task := Task{}

	err := row.Scan(
		&task.ID,
		&task.Name,
		&task.Priority,
		&task.CategoryID,
		&task.Status,
	)
	if err != nil {
		return task, err
	}

	return task, nil
}

func CompleteTask(db *sql.DB, id int64) error {

	_, err := db.Exec(
		`
		UPDATE tasks 
		SET status=?
		WHERE id=?
		`,
		true,
		id,
	)

	return err
}

func IncompleteTask(db *sql.DB, id int64) error {

	_, err := db.Exec(
		`
		UPDATE tasks 
		SET status=?
		WHERE id=?
		`,
		false,
		id,
	)

	return err
}

func UpdateTask(db *sql.DB, task Task) (Task, error) {

	_, err := db.Exec(
		`
		UPDATE tasks 
		SET 
			name=?,
			priority=?,
			category_id=?
		WHERE id=?`,
		task.Name, task.Priority, task.CategoryID, task.ID,
	)

	if err != nil {
		return task, err
	}

	task, err = ReadTask(db, task.ID)
	if err != nil {
		return task, err
	}

	return task, nil
}

func GetTaskID(db *sql.DB, name string) (int64, error) {

	response := db.QueryRow(`SELECT id FROM tasks WHERE name=? OR id=?`, name, name)

	var id int64

	err := response.Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func DeleteTask(db *sql.DB, id int64) error {
	_, err := db.Exec(`DELETE FROM tasks WHERE id=?`, id)
	return err
}

func ListTask(db *sql.DB) ([]Task, error) {

	rows, err := db.Query(`SELECT id, name, priority, category_id, status FROM tasks`)

	defer rows.Close()
	utilities.CheckNil(err, "", "")

	var tasks []Task

	for rows.Next() {
		task := Task{}

		err := rows.Scan(&task.ID, &task.Name, &task.Priority, &task.CategoryID, &task.Status)
		utilities.CheckNil(err, "", "")

		tasks = append(tasks, task)
	}
	if err = rows.Err(); err != nil {
		return tasks, err
	}
	return tasks, nil
}

func ShowUncheckedTasks(db *sql.DB) ([]Task, error) {

	rows, err := db.Query(`
		SELECT 
			id,
			name,
			category_id,
			normalized_task_priority + normalized_category_priority AS total_priority
		FROM (
			SELECT 
				tasks.id, 
				tasks.name, 
				tasks.category_id,
				CAST(tasks.priority AS FLOAT) / AVG(tasks.priority) OVER () AS normalized_task_priority, 
				CASE WHEN categories.priority IS NULL THEN 0
				ELSE CAST(categories.priority AS FLOAT) / AVG(categories.priority) OVER () END AS normalized_category_priority 
			FROM tasks
			FULL OUTER JOIN categories
			ON tasks.category_id = categories.id
			WHERE tasks.status=FALSE
		)
		ORDER BY total_priority DESC
	`)

	defer rows.Close()
	utilities.CheckNil(err, "", "")

	var tasks []Task
	var weightage *float64

	for rows.Next() {
		task := Task{}

		err := rows.Scan(&task.ID, &task.Name, &task.CategoryID, &weightage)
		utilities.CheckNil(err, "", "")

		tasks = append(tasks, task)
	}
	if err = rows.Err(); err != nil {
		return tasks, err
	}
	return tasks, nil
}
