package main

import (
	"log"
	"net/http"

	"github.com/glebarez/sqlite"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	"github.com/victor-bologna/pos-curso-go-expert-apis/configs"
	_ "github.com/victor-bologna/pos-curso-go-expert-apis/docs"
	"github.com/victor-bologna/pos-curso-go-expert-apis/internal/entity"
	"github.com/victor-bologna/pos-curso-go-expert-apis/internal/infra/database"
	"github.com/victor-bologna/pos-curso-go-expert-apis/internal/infra/webservers/handlers"
	"gorm.io/gorm"
)

// @title Pós Curso Go Expert API
// @version 1.0
// @description Product API with user autentication via JWT.
// @termsOfService http://swagger.io/terms

// @contact.name Victor Barros Bologna
// @contact.url https://linkedin.com/in/victor-bologna
// @contact.email victorbologna@hotmail.com

// @host localhost:8000
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	config := configs.LoadConfig("./")

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.User{}, &entity.Product{})

	productDB := database.NewProductDB(db)
	productHandler := handlers.NewProductHandler(productDB)

	userDB := database.NewUserDB(db)
	// userHandler := handlers.NewUserHandler(userDB, config.JWTAuth, config.ExpiresIn) // -> Passa coisas do JWT para  Handler
	userHandler := handlers.NewUserHandler(userDB)

	r := chi.NewRouter()
	r.Use(middleware.Logger)    // Chi Middleware
	r.Use(LogHandler)           // Custom Middleware
	r.Use(middleware.Recoverer) // Evita de cair o servidor caso algum erro intero aconteça.

	// Ou Passar coisas do JWT pelo middleware (Context) (Chave valor)
	r.Use(middleware.WithValue("Jwt", config.JWTAuth))
	r.Use(middleware.WithValue("JwtExpiresIn", config.ExpiresIn))

	r.Route("/products", func(r chi.Router) { // Mesma coisa que o RequestMapping do Controller do Spring.
		r.Use(jwtauth.Verifier(config.JWTAuth)) // (Middleware) Pega o token enviado e injeta o config.JWTAuth no contexto do chi.
		r.Use(jwtauth.Authenticator)            // Valida se o token é válido mesmo.

		r.Post("/", productHandler.CreateProduct)
		r.Get("/{id}", productHandler.FindProductByID)
		r.Get("/", productHandler.FindAllProducts)
		r.Put("/{id}", productHandler.UpdateProduct)
		r.Delete("/{id}", productHandler.DeleteProduct)
	})

	r.Post("/users", userHandler.CreateUser)
	r.Post("/users/generate_token", userHandler.GenerateToken)

	r.Get("/swagger/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8000/swagger/doc.json")))
	http.ListenAndServe(":8000", r)
}

func LogHandler(next http.Handler) http.Handler { // Custom Middleware
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
