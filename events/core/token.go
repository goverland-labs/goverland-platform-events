package core

import (
	"github.com/google/uuid"
	"github.com/goverland-labs/goverland-platform-events/events"
	"time"
)

const (
	DaoTokenPriceUpdated = "token.price.updated"
)

type TokenPricePayload struct {
	DaoID uuid.UUID `json:"dao_id"`
	Time  time.Time `json:"time"`
	Price float64   `json:"price"`
}

type TokenPricesPayload []TokenPricePayload

type TokenPricesHandler = events.Handler[TokenPricesPayload]
