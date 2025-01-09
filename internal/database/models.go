// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package database

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Language struct {
	ID   uuid.UUID
	Name string
}

type RefreshToken struct {
	Token     string
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    uuid.UUID
	ExpiresAt time.Time
	RevokedAt sql.NullTime
}

type Snippet struct {
	ID          uuid.UUID
	CreatedAt   time.Time
	UpdatedAt   time.Time
	LanguageID  uuid.UUID
	UserID      uuid.UUID
	SnippetText string
}

type User struct {
	ID             uuid.UUID
	CreatedAt      time.Time
	UpdatedAt      time.Time
	Email          string
	HashedPassword string
}
