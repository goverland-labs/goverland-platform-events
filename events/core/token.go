package core

import (
	"github.com/google/uuid"
	"github.com/goverland-labs/goverland-platform-events/events"
)

const (
	DaoTokenPriceUpdated = "token.price.updated"
)

type TokenPricePayload struct {
	DaoID uuid.UUID `json:"dao_id"`
	Price float64   `json:"price"`
}

type TokenPricesPayload []TokenPricePayload

type TokenPricesHandler = events.Handler[TokenPricesPayload]
