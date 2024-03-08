package api

import (
	"encoding/json"
	"fmt"
	"github.com/BananaPilot/sample-api/storage"
	"github.com/gorilla/mux"
	"net/http"
)

type ApiServer struct {
	ListenAddress string
	storage       *dbConn
}

func NewApiServer(listenAddress string, storage *dbConn) *ApiServer {
	return &ApiServer{
		ListenAddress: listenAddress,
		storage:       storage,
	}
}

type apiFunc func(w http.ResponseWriter, r *http.Request) error

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

func handlerFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, ApiError{Error: err})
		}
	}
}

type ApiError struct {
	Error error
}

func (server *ApiServer) Run() {

	router := mux.NewRouter()

	router.HandleFunc("/todo", handlerFunc(server.handleTodo))
	router.HandleFunc("/todo/{id}", handlerFunc(server.handleTodoWithID))

	err := http.ListenAndServe(server.ListenAddress, router)
	if err != nil {
		return
	}
}

func (server *ApiServer) handleTodo(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return server.handleGetTodo(w, r)
	}
	if r.Method == "POST" {
		return server.handleCreateTodo(w, r)
	}
	if r.Method == "DELETE" {
		return server.handleDeleteTodo(w, r)
	}
	return fmt.Errorf("method not allowed")
}

func (server *ApiServer) handleTodoWithID(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return server.handleGetIdTodo(w, r)
	}
	if r.Method == "DELETE" {
		return server.handleDeleteTodo(w, r)
	}
	return fmt.Errorf("method not allowed")
}

func (server *ApiServer) handleGetTodo(w http.ResponseWriter, r *http.Request) error {
	todos, err := server.storage.getTodos()
	if err != nil {
		return err
	}
	return WriteJSON(w, http.StatusOK, todos)
}

func (server *ApiServer) handleGetIdTodo(w http.ResponseWriter, r *http.Request) error {
	id := mux.Vars(r)["id"]
	todo, err := server.storage.getTodo(id)
	if err != nil {
		return err
	}
	return WriteJSON(w, http.StatusOK, todo)
}

func (server *ApiServer) handleDeleteTodo(w http.ResponseWriter, r *http.Request) error {
	id := mux.Vars(r)["id"]
	err := server.storage.deleteTodo(id)
	if err != nil {
		return err
	}
	return WriteJSON(w, http.StatusOK, "deleted")
}

func (server *ApiServer) handleCreateTodo(w http.ResponseWriter, r *http.Request) error {
	todoNewRequest := new(storage.TodoNewRequest)
	if err := json.NewDecoder(r.Body).Decode(todoNewRequest); err != nil {
		return err
	}

	todo := storage.NewTodo(todoNewRequest.Title, todoNewRequest.Description)
	if err := server.storage.postTodo(todo); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, todo)
}
