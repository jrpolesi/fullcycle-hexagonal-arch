package handler

import (
	"encoding/json"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/jrpolesi/fullcycle-hexagonal-arch/adapters/dto"
	"github.com/jrpolesi/fullcycle-hexagonal-arch/application"
)

func MakeProductHandlers(r *mux.Router, n *negroni.Negroni, service application.ProductServiceInterface) {
	r.Handle("/products", n.With(
		negroni.Wrap(createProduct(service)),
	)).Methods("POST", "OPTIONS")
	r.Handle("/products/{id}", n.With(
		negroni.Wrap(getProduct(service)),
	)).Methods("GET", "OPTIONS")
	r.Handle("/products/{id}/enable", n.With(
		negroni.Wrap(enableProduct(service)),
	)).Methods("PATCH", "OPTIONS")
	r.Handle("/products/{id}/disable", n.With(
		negroni.Wrap(disableProduct(service)),
	)).Methods("PATCH", "OPTIONS")
}

func getProduct(service application.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")

		vars := mux.Vars(r)
		id := vars["id"]

		product, err := service.Get(id)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		productDto := dto.NewProduct().ToDTO(product)
		err = json.NewEncoder(w).Encode(productDto)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	})
}

func createProduct(service application.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")

		var productDto dto.Product
		err := json.NewDecoder(r.Body).Decode(&productDto)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonError(err.Error()))
			return
		}

		product, err := service.Create(productDto.Name, productDto.Price)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonError(err.Error()))
			return
		}

		productDto.ToDTO(product)
		err = json.NewEncoder(w).Encode(productDto)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonError(err.Error()))
			return
		}
	})
}

func enableProduct(service application.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")

		vars := mux.Vars(r)
		id := vars["id"]

		product, err := service.Get(id)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		product, err = service.Enable(product)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonError(err.Error()))
			return
		}

		productDto := dto.NewProduct().ToDTO(product)
		err = json.NewEncoder(w).Encode(productDto)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	})
}

func disableProduct(service application.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")

		vars := mux.Vars(r)
		id := vars["id"]

		product, err := service.Get(id)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		product, err = service.Disable(product)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonError(err.Error()))
			return
		}

		productDto := dto.NewProduct().ToDTO(product)
		err = json.NewEncoder(w).Encode(productDto)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	})
}
