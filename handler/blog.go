package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/loggerboy9325/website-go-htmx-temple/article"
	"github.com/loggerboy9325/website-go-htmx-temple/view/blog"
)

func BlogIndexHandler(w http.ResponseWriter, r *http.Request) error {
	articles, err := article.GetAssets()
	if err != nil {
		return err
	}
	return render(w, r, blog.Blog(articles))
}

func BlogArticleHandler(w http.ResponseWriter, r *http.Request, parser article.Parser) error {
	slug := chi.URLParam(r, "slug")

	article, err := parser.Parse(slug)
	if err != nil {
		return err
	}

	return render(w, r, blog.Article(article))
}
