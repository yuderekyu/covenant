package models

import (
	"database/sql"
	"errors"
	"time"

	"github.com/pborman/uuid"
)

/*Subscription is the representation of subscription entries in covenant*/
type Subscription struct {
	ID        uuid.UUID          `json:"id"`
	UserID    uuid.UUID          `json:"userId"`
	Status    SubscriptionStatus `json:"status"`
	CreatedAt time.Time          `json:"createdAt"`
	NextOrder time.Time          `json:"time"`
	Frequency string             `json:"frequency"`
	RoasterID uuid.UUID          `json:"roasterId"`
	ItemID    uuid.UUID          `json:"itemId"`
	Quantity  uint64             `json:"quantity"`
}

/*RequestSubscription represents the data needed to create a new subscription entry*/
type RequestSubscription struct {
	UserID    uuid.UUID `json:"userId" binding:"required"`
	Frequency string    `json:"frequency" binding:"required"`
	RoasterID uuid.UUID `json:"roasterId" binding:"required"`
	ItemID    uuid.UUID `json:"itemId" binding:"required"`
	Quantity  uint64    `json:"quantity" binding:"required"`
}

/*RequestOrder represents the data needed to create a new order entry wihin Warehouse*/
type RequestOrder struct {
	UserID    uuid.UUID `json:"userId" binding:"required"`
	ItemID    uuid.UUID `json:"itemId" binding:"required"`
	NextOrder time.Time `json:"nextOrder" bindng:"required"`
	Quantity  uint64    `json"quantity" binding:"required"`
}

/*NewSubscription creates a new subscription with a new uuid*/
func NewSubscription(userID uuid.UUID, frequency string, roasterID uuid.UUID, itemID uuid.UUID, quantity uint64) *Subscription {
	return &Subscription{
		ID:        uuid.NewUUID(),
		UserID:    userID,
		Status:    PENDING,
		CreatedAt: time.Now(),
		Frequency: frequency,
		RoasterID: roasterID,
		ItemID:    itemID,
		Quantity:  quantity,
	}
}

/*SubscriptionFromSql returns a new subscription slice from a group of sql rows*/
func SubscriptionFromSql(rows *sql.Rows) ([]*Subscription, error) {
	subscription := make([]*Subscription, 0)

	for rows.Next() {
		s := &Subscription{}
		var sStatus string

		rows.Scan(&s.ID, &s.UserID, &sStatus, &s.CreatedAt, &s.Frequency, &s.RoasterID, &s.ItemID, &s.Quantity, &s.NextOrder)

		var ok bool
		s.Status, ok = toSubscriptionType(sStatus)
		if !ok {
			return nil, errors.New("invalid subscriptionStatus string")
		}
		subscription = append(subscription, s)
	}

	return subscription, nil
}

/*toSubscriptionType checks whether a given status is valid*/
func toSubscriptionType(s string) (SubscriptionStatus, bool) {
	switch s {
	case ACTIVE:
		return ACTIVE, true
	case PENDING:
		return PENDING, true
	case CANCELLED:
		return CANCELLED, true
	case INACTIVE:
		return INACTIVE, true
	default:
		return "", false
	}
}

/*SubscriptionStatus is an enum wrapper for valid subscription type*/
type SubscriptionStatus string

/*Valid Subscription Statuses*/
const (
	ACTIVE    = "ACTIVE"
	PENDING   = "PENDING"
	CANCELLED = "CANCELLED"
	INACTIVE  = "INACTIVE"
)
