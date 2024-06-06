package api

import (
	"net/http"
	"user/api/handler"
	"user/storage"
)

func Router(st *storage.User) http.Handler {
	mux := http.NewServeMux()
	handler := handler.NewHandeler(st)
	mux.HandleFunc("POST /user", handler.CreateUser)
	mux.HandleFunc("PUT /user/id", handler.UpdateUserById)
	mux.HandleFunc("DELETE /user/id", handler.DeleteUserByID)

	return mux
}
