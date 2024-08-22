package handler

import "net/http"

type person struct {
	Storage storage
}

func newUserHandler(s storage) person {
	return person{s}

}

func RouteUser(mux *http.ServeMux, storage storage) {
	h := newUserHandler(storage)

	mux.HandleFunc("/v1/users/register", h.Storage.Register)
	mux.HandleFunc("/v1/users/login", h.Storage.Login) // Nueva ruta para login
	// mux.HandleFunc("/v1/users/get-all", h.getAllUsers)
	// mux.HandleFunc("/v1/users/update", h.updateUser)
	// mux.HandleFunc("/v1/users/delete", h.deleteUser)
	// mux.HandleFunc("/v1/users/get-by-id", h.getUserByID)
}
