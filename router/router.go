package router

import (
	"github.com/gorilla/mux"
	"github.com/piyush/golangtodobackend/controllers"
	"github.com/piyush/golangtodobackend/middleware"
	"github.com/piyush/golangtodobackend/middlewarebetween"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/signup", controllers.Signup).Methods("POST", "OPTIONS")
	router.HandleFunc("/login", controllers.Login).Methods("POST", "OPTIONS")

	authenticatedRouter := router.PathPrefix("/api").Subrouter()
	authenticatedRouter.Use(middlewarebetween.Authenticate)
	router.HandleFunc("/api/task", middleware.GetAllTasks).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/tasks", middleware.CreateTask).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/tasks/{id}", middleware.TaskComplete).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/undoTask/{id}", middleware.UndoTask).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/deleteTask/{id}", middleware.DeleteTask).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/api/deleteAllTasks", middleware.DeleteAllTasks).Methods("DELETE", "OPTIONS")

	return router
}
