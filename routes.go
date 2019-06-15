package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"GetTodos",
		"GET",
		"/v1/todos",
		GetTodos,
	},
	Route{
		"GetTodosByID",
		"GET",
		"/v1/todos/{id}",
		GetTodosByID,
	},
	Route{
		"PostTodosByID",
		"POST",
		"/v1/todos",
		PostTodos,
	},
	Route{
		"PutTodosByID",
		"PUT",
		"/v1/todos/{id}",
		PutTodosByID,
	},
	Route{
		"DeleteTodosByID",
		"DELETE",
		"/v1/todos/{id}",
		DeleteTodosByID,
	},
}
