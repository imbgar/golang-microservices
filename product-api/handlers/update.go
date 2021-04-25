package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/imbgar/golang-microservices/product-api/data"
)

func (p *Products) UpdateProduct(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	_, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unable to convert id", http.StatusBadRequest)
		return
	}

	p.l.Println("Handle PUT Product")

	prod := r.Context().Value(KeyProduct{}).(data.Product)

	err = data.UpdateProduct(prod)

	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}
}
