package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/victor-bologna/pos-curso-go-expert-apis/internal/dto"
	"github.com/victor-bologna/pos-curso-go-expert-apis/internal/entity"
	"github.com/victor-bologna/pos-curso-go-expert-apis/internal/httputil"
	"github.com/victor-bologna/pos-curso-go-expert-apis/internal/infra/database"
	"github.com/victor-bologna/pos-curso-go-expert-apis/pkg/property"
)

type ProductHandler struct {
	ProductInterface database.ProductInterface
}

func NewProductHandler(pi database.ProductInterface) *ProductHandler {
	return &ProductHandler{ProductInterface: pi}
}

// CreateProduct godoc
//
//	@Summary		Create Product
//	@Description	Generate an Product.
//	@Tags			Products
//	@Accept			json
//	@Produce		json
//	@Param			request	body	dto.ProductDTO	true	"Product information"
//	@Success		201
//	@Failure		400	{object}	httputil.Error
//	@Failure		500	{object}	httputil.Error
//	@Router			/products/ [post]
//	@Security 		ApiKeyAuth
func (p *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var productDTO dto.ProductDTO
	err := json.NewDecoder(r.Body).Decode(&productDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err := httputil.Error{Message: err.Error()}
		json.NewEncoder(w).Encode(err)
		return
	}

	product, err := entity.NewProduct(productDTO.Name, productDTO.Price)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err := httputil.Error{Message: err.Error()}
		json.NewEncoder(w).Encode(err)
		return
	}

	err = p.ProductInterface.Create(product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		err := httputil.Error{Message: err.Error()}
		json.NewEncoder(w).Encode(err)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// FindProductByID godoc
//
//	@Summary		Find Product
//	@Description	Find Product by ID.
//	@Tags			Products
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string		true	"Product ID"
//	@Success		200	{object} 	entity.Product
//	@Failure		400	{object}	httputil.Error
//	@Failure		404	{object}	httputil.Error
//	@Failure		500	{object}	httputil.Error
//	@Router			/products/{id} [get]
//	@Security 		ApiKeyAuth
func (p *ProductHandler) FindProductByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		err := httputil.Error{Message: "ID param must not be empty."}
		json.NewEncoder(w).Encode(err)
		return
	}

	product, err := p.ProductInterface.FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		err := httputil.Error{Message: err.Error()}
		json.NewEncoder(w).Encode(err)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&product)
}

// FindProducts godoc
//
//	@Summary		Find Products
//	@Description	Find all products.
//	@Tags			Products
//	@Accept			json
//	@Produce		json
//	@Param			page	query	string	false	"Page number"
//	@Param			limit	query	string	false	"Limit per page"
//	@Success		200	{array} 	entity.Product
//	@Failure		500	{object}	httputil.Error
//	@Router			/products/ [get]
//	@Security 		ApiKeyAuth
func (p *ProductHandler) FindAllProducts(w http.ResponseWriter, r *http.Request) {
	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")
	sort := r.URL.Query().Get("sort")

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		page = 0
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 0
	}

	products, err := p.ProductInterface.FindAll(page, limit, sort)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		err := httputil.Error{Message: err.Error()}
		json.NewEncoder(w).Encode(err)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}

// UpdateProduct godoc
//
//	@Summary		Update Product
//	@Description	Update product by ID.
//	@Tags			Products
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string		true	"Product ID"
//	@Param			request	body	dto.ProductDTO	true	"Product information to be updated."
//	@Success		200	{object} 	entity.Product
//	@Failure		400	{object}	httputil.Error
//	@Failure		404	{object}	httputil.Error
//	@Failure		500	{object}	httputil.Error
//	@Router			/products/{id} [put]
//	@Security 		ApiKeyAuth
func (p *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		err := httputil.Error{Message: "ID param must not be empty."}
		json.NewEncoder(w).Encode(err)
		return
	}

	var product entity.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err := httputil.Error{Message: err.Error()}
		json.NewEncoder(w).Encode(err)
		return
	}

	product.ID, err = property.ParseID(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err := httputil.Error{Message: err.Error()}
		json.NewEncoder(w).Encode(err)
		return
	}

	err = p.ProductInterface.Update(&product)
	if err != nil {
		if err.Error() == "record not found" {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
		err := httputil.Error{Message: err.Error()}
		json.NewEncoder(w).Encode(err)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&product)
}

// DeleteProduct godoc
//
//	@Summary		Delete Products
//	@Description	Delete product by ID.
//	@Tags			Products
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string		true	"Product ID"
//	@Success		204
//	@Failure		400	{object}	httputil.Error
//	@Failure		404	{object}	httputil.Error
//	@Failure		500	{object}	httputil.Error
//	@Router			/products/{id} [delete]
//	@Security 		ApiKeyAuth
func (p *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		err := httputil.Error{Message: "ID param must not be empty."}
		json.NewEncoder(w).Encode(err)
		return
	}

	_, err := property.ParseID(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err := httputil.Error{Message: err.Error()}
		json.NewEncoder(w).Encode(err)
		return
	}

	err = p.ProductInterface.DeleteByID(id)
	if err != nil {
		if err.Error() == "record not found" {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
		err := httputil.Error{Message: err.Error()}
		json.NewEncoder(w).Encode(err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
