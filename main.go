package main

import (
	"embed"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/loggerboy9325/website-go-htmx-temple/article"
	"github.com/loggerboy9325/website-go-htmx-temple/db"
	"github.com/loggerboy9325/website-go-htmx-temple/handler"
)

//go:embed public
var FS embed.FS

func main() {
	if err := initdb(); err != nil {
		log.Fatal(err)
	}

	parser := article.NewParser()

	router := chi.NewMux()

	router.Handle("/*", http.StripPrefix("/", http.FileServer(http.FS(FS))))
	router.Get("/", handler.Make(handler.HandlerHomeIndex))
	router.Get("/journal", handler.Make(handler.BlogIndexHandler))
	router.Get("/journal/{slug}", handler.Make(func(w http.ResponseWriter, r *http.Request) error {
		return handler.BlogArticleHandler(w, r, parser)
	}))
	router.Get("/contact-form", handler.Make(handler.ContactFormHandler))
	router.Get("/about-me", handler.Make(handler.AboutMeHandler))
	router.Get("/resume", handler.Make(handler.ResumeHandler))
	router.Get("/website-info", handler.Make(handler.HandlerWebsiteInfo))
	router.Post("/contact-form", handler.Make(handler.ContactFormSubmit))

	router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})
	slog.Info("listening on port 3000")

	port := os.Getenv("HTTP_LISTEN_ADDR")

	log.Fatal(http.ListenAndServe(port, router))
}

func initdb() error {
	if err := godotenv.Load(); err != nil {
		return err
	}
	if err := db.Init(); err != nil {
		return err
	}
	return nil
}
