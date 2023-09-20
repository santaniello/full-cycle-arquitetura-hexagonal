package handler

import (
	"encoding/json"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/santaniello/full-cycle-arquitetura-hexagonal/adapters/dto"
	"github.com/santaniello/full-cycle-arquitetura-hexagonal/application"
	"net/http"
)

// Esse arquivo seria como se fosse uma espécie de Controller do Spring
/*
Esse arquivo seria como se fosse uma espécie de Controller do Spring

A função negroni.Wrap é usada para converter um handler padrão do Go
(http.HandlerFunc ou http.Handler) em um middleware compatível com Negroni.
Isso permite que você insira um handler padrão do Go na pilha de middlewares
do Negroni.

n.With(...) é uma função do Negroni que permite criar uma nova instância
de Negroni com middlewares adicionais apenas para um endpoint específico.
No caso abaixo, o endpoint /product/{id} terá apenas o middleware fornecido
(o handler de produto).
*/

func MakeProductHandlers(router *mux.Router, n *negroni.Negroni, service application.ProductServiceInterface) {
	router.Handle("/product/{id}", n.With(
		negroni.Wrap(getProduct(service)),
	)).Methods("GET", "OPTIONS")
	router.Handle("/product", n.With(
		negroni.Wrap(createProduct(service)),
	)).Methods("POST", "OPTIONS")
	router.Handle("/product/{id}/enable", n.With(
		negroni.Wrap(enableProduct(service)),
	)).Methods("PUT", "OPTIONS")
	router.Handle("/product/{id}/disable", n.With(
		negroni.Wrap(disableProduct(service)),
	)).Methods("PUT", "OPTIONS")
}

func createProduct(service application.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var productDto dto.ProductDto
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
		err = json.NewEncoder(w).Encode(product)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonError(err.Error()))
			return
		}
	})
}

/*
 */
func getProduct(service application.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(r)
		id := vars["id"]
		product, err := service.Get(id)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		err = json.NewEncoder(w).Encode(product)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	})
}

func enableProduct(service application.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(r)
		id := vars["id"]
		product, err := service.Get(id)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		result, err := service.Enable(product)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonError(err.Error()))
			return
		}
		err = json.NewEncoder(w).Encode(result)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	})
}

func disableProduct(service application.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(r)
		id := vars["id"]
		product, err := service.Get(id)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		err = product.ChangePrice(0)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonError(err.Error()))
			return
		}

		result, err := service.Disable(product)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonError(err.Error()))
			return
		}
		err = json.NewEncoder(w).Encode(result)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	})
}
