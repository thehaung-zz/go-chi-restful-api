package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/joho/godotenv"
	"github.com/thehaung/go-chi-restful-api/pkg/service"
	"log"
	"net/http"
	"os"
)

func main() {
	port := "8080"

	err := godotenv.Load(".env")
	if err != nil {
		return
	}

	if fromEnv := os.Getenv("PORT"); fromEnv != "" {
		port = fromEnv
	}

	log.Printf("Starting up on http://localhost:%s", port)

	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	r.Mount("/posts", service.PostsResource{}.Routes())

	log.Fatal(http.ListenAndServe(":"+port, r))
}
