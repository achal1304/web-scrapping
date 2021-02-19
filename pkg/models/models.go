package models

import (
	"errors"
	"time"
)

var (
	ErrNoRecord = errors.New("models: no matching record found")
	// Add a new ErrInvalidCredentials error. We'll use this later if a user
	// tries to login with an incorrect email address or password.
	ErrInvalidCredentials = errors.New("models: invalid credentials")
	// Add a new ErrDuplicateEmail error. We'll use this later if a user
	// tries to signup with an email address that's already in use.
	ErrDuplicateEmail = errors.New("models: duplicate email")
)

type Scrap struct {
	ID      int
	Soc_id  string
	Email   string
	Guid    string
	Count   int
	Created time.Time
	Expires time.Time
}

type Otps struct {
	ID       int
	Otp      int64
	Verified bool
	Created  time.Time
	Expires  time.Time
}

/*
type User struct {
	ID             int
	Email          string
	HashedPassword []byte
	Created        time.Time
	Active         bool
}
*/
