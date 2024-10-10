package handler

import (
	"net/http"

	"github.com/loggerboy9325/website-go-htmx-temple/view/websiteinfo"
)

func HandlerWebsiteInfo(w http.ResponseWriter, r *http.Request) error {
	return render(w, r, websiteinfo.Websiteinfo())
}
