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

func (conn *dbConn) getTodos() ([]*Todo, error) {
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

func (conn *dbConn) postTodo(todo *Todo) error {
	query := `insert into Todo(description, title) values (?, ?)`

	_, err := conn.db.Query(query, todo.Title, todo.Description)
	if err != nil {
		return err
	}

	return nil

}

func (conn *dbConn) getTodo(ID string) (*Todo, error) {
	query := `select * from Todo where id = ?`
	res, err := conn.db.Query(query, ID)
	todo := new(Todo)
	if err != nil {
		return nil, err
	}
	for res.Next() {
		err := res.Scan(&todo.ID, &todo.Title, &todo.Description)
		if err != nil {
			panic(err.Error())
		}
	}
	return todo, nil
}

func (conn *dbConn) deleteTodo(ID string) error {
	query := `delete from Todo where id = ?`

	_, err := conn.db.Query(query, ID)
	if err != nil {
		return err
	}
	return nil
}
