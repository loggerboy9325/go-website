package types

import "time"

type Contact struct {
	CreatedAt time.Time `bun:"default:'now()'"`
	Email     string
	Name      string
	Message   string
	ID        int `bun:"id,pk,autoincrement"`
}
