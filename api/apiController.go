package api

import (
	"database/sql"
	"github.com/BananaPilot/sample-api/storage"
)

type dbConn struct {
	db *sql.DB
}

func NewDbConn(conn *sql.DB) *dbConn {
	return &dbConn{
		db: conn,
	}
}

func (conn *dbConn) getTodos() ([]*storage.Todo, error) {
	var todos []*storage.Todo
	res, err := conn.db.Query("SELECT * FROM todo")
	if err != nil {
		panic(err.Error())
	}
	for res.Next() {
		todo := new(storage.Todo)
		err := res.Scan(&todo.ID, &todo.Title, &todo.Description)
		if err != nil {
			panic(err.Error())
		}
		todos = append(todos, todo)
	}
	return todos, nil
}

func (conn *dbConn) postTodo(todo *storage.Todo) error {
	query := `insert into Todo(description, title) values (?, ?)`

	_, err := conn.db.Query(query, todo.Title, todo.Description)
	if err != nil {
		return err
	}

	return nil

}

func (conn *dbConn) getTodo(ID string) (*storage.Todo, error) {
	query := `select * from Todo where id = ?`
	res, err := conn.db.Query(query, ID)
	todo := new(storage.Todo)
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
