// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Mutation struct {
}

type OrderRequest struct {
	ID    string  `json:"ID"`
	Price float64 `json:"Price"`
	Tax   float64 `json:"Tax"`
}

type OrderResponse struct {
	ID         string  `json:"ID"`
	Price      float64 `json:"Price"`
	Tax        float64 `json:"Tax"`
	FinalPrice float64 `json:"FinalPrice"`
}

type OrderResponseList struct {
	Orders []*OrderResponse `json:"Orders"`
}

type Query struct {
}
