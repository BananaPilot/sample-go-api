package main

import "database/sql"

type dbConn struct {
	db *sql.DB
}

func newDbConn(conn *sql.DB) *dbConn {
	return &dbConn{
		db: conn,
	}
}

func (conn *dbConn) getTodo() ([]*Todo, error) {
	var todos []*Todo
	res, err := conn.db.Query("SELECT * FROM todo")
	if err != nil {
		panic(err.Error())
	}
	for res.Next() {
		todo := new(Todo)
		err := res.Scan(&todo.ID, &todo.Title, &todo.Description)
		if err != nil {
			panic(err.Error())
		}
		todos = append(todos, todo)
	}
	return todos, nil
}

func (conn *dbConn) postConn(todo *Todo) error {
	query := `insert into Todo(description, title) values (?, ?)`

	_, err := conn.db.Query(query, todo.Title, todo.Description)
	if err != nil {
		return err
	}

	return nil
}
