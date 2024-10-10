package types

import "time"

type Article struct {
	CreatedAt time.Time `bun:"default:'now()'"`
	UpdatedAt time.Time `bun:"default:'now()'"`
	Title     string
	Slug      string
	Filename  string
	ID        int `bun:"id,pk,autoincrement"`
}
