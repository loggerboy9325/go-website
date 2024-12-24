package db

import (
	"context"

	"github.com/loggerboy9325/website-go-htmx-temple/types"
)

func CreateContact(contact *types.Contact) error {
	_, err := Bun.NewInsert().
		Model(contact).
		Exec(context.Background())
	return err
}
