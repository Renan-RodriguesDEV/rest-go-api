package handlers

import (
	"encoding/json"
	"net/http"
	"rest-go-api/internal/service"
	"strconv"
)

// cria uma estrutura UserHandler que contém um serviço de usuário
// o * é usado para indicar que service é um ponteiro para UserService, permitindo que o handler acesse os métodos do serviço
type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) Get(w http.ResponseWriter, r *http.Request) {
	id,err:=strconv.Atoi(r.URL.Query().Get("id"))
	if err!=nil{
		http.Error(w, "Invalid id parameter", http.StatusBadRequest)
		return
	}
	user,err:=h.service.Get(r.Context(),id)
	if err!=nil{
		http.Error(w, "Error occurred while fetching user", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)

}
