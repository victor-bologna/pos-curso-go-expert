package dto

type ProductDTO struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type UserDTO struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type GenerateTokenDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type JWTString struct {
	AccessToken string `json:"access_token"`
}
