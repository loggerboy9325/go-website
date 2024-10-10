package handler

import (
	"net/http"

	aboutme "github.com/loggerboy9325/website-go-htmx-temple/view/about_me"
)

func AboutMeHandler(w http.ResponseWriter, r *http.Request) error {
	return render(w, r, aboutme.Aboutme())
}
