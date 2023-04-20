package token

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"time"
)

type Payload struct {
	jwt.RegisteredClaims
	Username string `json:"username"`
}

func (payload *Payload) Valid() bool {
	return time.Now().Before(payload.ExpiresAt.Time)
}

func NewPayload(username string, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &Payload{Username: username}
	payload.ID = tokenID.String()
	payload.IssuedAt = jwt.NewNumericDate(time.Now())
	payload.ExpiresAt = jwt.NewNumericDate(time.Now().Add(duration))

	return payload, nil
}
