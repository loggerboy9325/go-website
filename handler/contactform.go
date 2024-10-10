package handler

import (
	"net/http"

	"github.com/loggerboy9325/website-go-htmx-temple/db"
	validate "github.com/loggerboy9325/website-go-htmx-temple/pkg/kit"
	"github.com/loggerboy9325/website-go-htmx-temple/types"
	contactform "github.com/loggerboy9325/website-go-htmx-temple/view/contact_form"
)

func ContactFormHandler(w http.ResponseWriter, r *http.Request) error {
	return render(w, r, contactform.Index())
}

func ContactFormSubmit(w http.ResponseWriter, r *http.Request) error {
	params := contactform.ContactParams{
		Email:   r.FormValue("email"),
		Name:    r.FormValue("name"),
		Message: r.FormValue("message"),
	}

	var errors contactform.ContactErrors

	ok := validate.New(params, validate.Fields{
		"Message": validate.Rules(validate.Min(10), validate.Max(100)),
	}).Validate(&errors)
	if !ok {
		return render(w, r, contactform.Contactform(params, errors))
	}
	ok = validate.New(params, validate.Fields{
		"Name": validate.Rules(validate.Min(5), validate.Max(20)),
	}).Validate(&errors)
	if !ok {
		return render(w, r, contactform.Contactform(params, errors))
	}

	ok = validate.New(params, validate.Fields{
		"Email": validate.Rules(validate.Min(5), validate.Max(100)),
	}).Validate(&errors)
	if !ok {
		return render(w, r, contactform.Contactform(params, errors))
	}

	params1 := types.Contact{
		Email:   params.Email,
		Name:    params.Name,
		Message: params.Message,
	}

	if err := db.CreateContact(&params1); err != nil {
		return err
	}

	return render(w, r, contactform.ContactSuccess(params))
}
