package models

import (
	"database/sql"
	"errors"

	"github.com/pborman/uuid"
)

type Subscription struct {
	ID uuid.UUID `json: "id"`
	UserID uuid.UUID `json: "userId"`
	Status SubscriptionStatus `json:"status"` 
	CreatedAt string `json:"createdAt"` 
	StartAt string `json:"startAt"` 
	ShopID uuid.UUID `json: "shopId"`
	OZInBag float64 `json: "ozInBag"`
	BeanName string `json:"beanName"`
	RoastName string `json: "roastName"`
	Price float64 `json: "price"`
}

func NewSubscription(userID uuid.UUID, createdAt string, startAt string, shopID uuid.UUID, ozInBag float64, beanName string, roastName string, price float64) *Subscription {
	return &Subscription{ 
		ID: uuid.NewUUID(), 
		UserID: userID, 
		Status: ACTIVE, 
		CreatedAt: createdAt, 
		StartAt: startAt, 
		ShopID: shopID, 
		OZInBag: ozInBag, 
		BeanName: beanName,
		RoastName: roastName,
		Price: price,
	}
}


func SubscriptionFromSql(rows *sql.Rows) ([]*Subscription, error) {
	subscription := make([]*Subscription,0)

	for rows.Next() {
		s := &Subscription{}
		var sStatus string
		rows.Scan(&s.ID, &s.UserID, &sStatus, &s.CreatedAt, &s.StartAt, &s.ShopID, &s.OZInBag, &s.BeanName, &s.RoastName, &s.Price)
		subscription = append(subscription, s)

		var ok bool
		s.Status, ok = toSubscriptionType(sStatus)
		if !ok {
			return nil, errors.New("invalid subscriptionStatus string")
		}
		subscription = append(subscription, s)
	}


	return subscription, nil
}

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
	ACTIVE = "ACTIVE"
	PENDING = "PENDING"
	CANCELLED = "CANCELLED"
	INACTIVE = "INACTIVE"
)
