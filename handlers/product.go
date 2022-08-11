package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/fxmb/go-server/data"
	"github.com/gorilla/mux"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	err := lp.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Unable to Marshal JSON", http.StatusInternalServerError)
	}
}

func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	fmt.Println(prod)
	if err != nil {
		http.Error(rw, "Unbale to Unmarshal JSON", http.StatusBadRequest)
	}
	data.AddProduct(prod)
}

func (p *Products) UpdateProducts(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unable to Unpack ID", http.StatusBadRequest)
	}

	prod := &data.Product{}
	err = prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unbale to Unmarshal JSON", http.StatusBadRequest)
	}
	err = data.UpdateProduct(id, prod)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(rw, "Product could not be added", http.StatusInternalServerError)
		return
	}
}
